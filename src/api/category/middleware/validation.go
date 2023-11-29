package middleware

import (
	"context"
	"tiny-template/src/api/category"
	"tiny-template/src/api/category/transport"
)

type validationMiddleware struct {
	next category.ServiceInterface
}

func NewValidationMiddleware() Middleware {
	return func(next category.ServiceInterface) category.ServiceInterface {
		return &validationMiddleware{
			next: next,
		}
	}
}

func (v validationMiddleware) GetCategoryByID(ctx context.Context, req *transport.GetCategoryByIDRequest) (*transport.GetCategoryByIDResponse, error) {
	//TODO implement me
	panic("implement me")
}
