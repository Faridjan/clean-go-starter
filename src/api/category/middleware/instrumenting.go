package middleware

import (
	"context"
	"time"
	"tiny-template/src/api/category"
	"tiny-template/src/api/category/transport"

	"github.com/go-kit/kit/metrics"
)

type instrumentingMiddleware struct {
	next           category.ServiceInterface
	requestCount   metrics.Counter
	requestError   metrics.Counter
	requestLatency metrics.Histogram
}

func (im *instrumentingMiddleware) instrumenting(begin time.Time, method string, err error) {
	im.requestCount.With("method", method).Add(1)
	if err != nil {
		im.requestError.With("method", method).Add(1)
	}
	im.requestLatency.With("method", method).Observe(time.Since(begin).Seconds())
}

func NewInstrumentingMiddleware(counter, counterErr metrics.Counter, latency metrics.Histogram) Middleware {
	return func(next category.ServiceInterface) category.ServiceInterface {
		return &instrumentingMiddleware{
			next:           next,
			requestCount:   counter,
			requestError:   counterErr,
			requestLatency: latency,
		}
	}
}
func (im *instrumentingMiddleware) GetCategoryByID(ctx context.Context, req *transport.GetCategoryByIDRequest) (*transport.GetCategoryByIDResponse, error) {
	//TODO implement me
	panic("implement me")
}
