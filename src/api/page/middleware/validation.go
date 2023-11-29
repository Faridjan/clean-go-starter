package middleware

import (
	"context"
	"tiny-template/src/api/page"
	"tiny-template/src/api/page/transport"
)

type validationMiddleware struct {
	next page.ServiceInterface
}

func NewValidationMiddleware() Middleware {
	return func(next page.ServiceInterface) page.ServiceInterface {
		return &validationMiddleware{
			next: next,
		}
	}
}

func (v validationMiddleware) GetPageByID(ctx context.Context, req *transport.GetPageByIDRequest) (*transport.GetPageByIDResponse, error) {
	//TODO implement me
	panic("implement me")
}
