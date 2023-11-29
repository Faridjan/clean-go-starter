package pkg

import (
	"context"
	"errors"
	"fmt"
	"os"
	"os/signal"
	"strings"
	"sync"
	"syscall"
	"time"

	"golang.org/x/sync/errgroup"
)

const shutdownTimeout = 5 * time.Minute

type Func func(ctx context.Context) error

type GracefullyShutdown struct {
	mu             sync.Mutex
	ctx            context.Context
	errGroup       *errgroup.Group
	closeFunctions []Func
}

func NewGracefullyShutdown(parentCtx context.Context) (gracefully *GracefullyShutdown) {
	g, ctx := errgroup.WithContext(parentCtx)

	gfl := &GracefullyShutdown{
		ctx:      ctx,
		errGroup: g,
	}

	gfl.Go(gfl.listenerOS)
	gfl.Go(gfl.killer)

	return gfl
}

func (gfl *GracefullyShutdown) Go(foo func() error) {
	gfl.errGroup.Go(func() (err error) {
		defer func() {
			if errPanic := recover(); errPanic != nil { //catch panic
				err = errors.New(fmt.Sprint("Recovered in gracefully thread: ", errPanic))
			}
		}()

		if errF := foo(); errF != nil {
			return errF
		}

		return err
	})
}

func (gfl *GracefullyShutdown) Wait() {
	_ = gfl.errGroup.Wait()
}

func (gfl *GracefullyShutdown) MustClose(f Func) {
	gfl.mu.Lock()
	defer gfl.mu.Unlock()

	gfl.closeFunctions = append(gfl.closeFunctions, f)
}

func (gfl *GracefullyShutdown) listenerOS() error {
	ch := make(chan os.Signal, 1)
	signal.Notify(ch, syscall.SIGTERM, syscall.SIGINT, syscall.SIGKILL, syscall.SIGQUIT)

	select {
	case <-gfl.ctx.Done():
		return nil
	case <-ch:
		return errors.New("OS Signal")
	}
}

func (gfl *GracefullyShutdown) killer() error {
	<-gfl.ctx.Done()

	ctx, closeT := context.WithTimeout(context.Background(), shutdownTimeout)
	defer closeT()

	if err := gfl.close(ctx); err != nil {
		panic(err)
	}

	return nil
}

func (gfl *GracefullyShutdown) close(ctx context.Context) error {
	gfl.mu.Lock()
	defer gfl.mu.Unlock()

	closeErrMessages := make([]string, 0, len(gfl.closeFunctions))
	complete := make(chan struct{}, 1)

	go func() {
		for _, f := range gfl.closeFunctions {
			if err := f(ctx); err != nil {
				closeErrMessages = append(closeErrMessages, fmt.Sprintf("[!] %v", err))
			}
		}

		complete <- struct{}{}
	}()

	select {
	case <-complete:
		break
	case <-ctx.Done():
		return fmt.Errorf("shutdown cancelled: %v", ctx.Err())
	}

	if len(closeErrMessages) > 0 {
		return fmt.Errorf(
			"shutdown finished with error(s): \n%s",
			strings.Join(closeErrMessages, "\n"),
		)
	}

	return nil
}
