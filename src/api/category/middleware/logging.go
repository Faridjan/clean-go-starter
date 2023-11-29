package middleware

import (
	"context"
	"tiny-template/pkg/logger"
	"tiny-template/src/api/category"
	"tiny-template/src/api/category/transport"
)

type loggingMiddleware struct {
	next category.ServiceInterface
}

func (lm *loggingMiddleware) logging(method string, err error) {
	if err != nil {
		logger.Error("method", method, "err", err)
	} else {
		logger.Info("method", method)
	}
}

func NewLoggingMiddleware() Middleware {
	return func(next category.ServiceInterface) category.ServiceInterface {
		return &loggingMiddleware{
			next: next,
		}
	}
}
func (lm *loggingMiddleware) GetCategoryByID(ctx context.Context, req *transport.GetCategoryByIDRequest) (_ *transport.GetCategoryByIDResponse, err error) {
	defer lm.logging("GetCategoryByID", err)
	return lm.next.GetCategoryByID(ctx, req)
}
