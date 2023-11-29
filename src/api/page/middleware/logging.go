package middleware

import (
	"context"
	"tiny-template/pkg/logger"
	"tiny-template/src/api/page"
	"tiny-template/src/api/page/transport"
)

type loggingMiddleware struct {
	next page.ServiceInterface
}

func (lm *loggingMiddleware) logging(method string, err error) {
	if err != nil {
		logger.Error("method", method, "err", err)
	} else {
		logger.Info("method", method)
	}
}

func NewLoggingMiddleware() Middleware {
	return func(next page.ServiceInterface) page.ServiceInterface {
		return &loggingMiddleware{
			next: next,
		}
	}
}
func (lm *loggingMiddleware) GetPageByID(ctx context.Context, req *transport.GetPageByIDRequest) (_ *transport.GetPageByIDResponse, err error) {
	defer lm.logging("GetPageByID", err)
	return lm.next.GetPageByID(ctx, req)
}
