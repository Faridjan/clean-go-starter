package category

import (
	"context"
	"tiny-template/src/api/category/transport"
)

type ServiceInterface interface {
	GetCategoryByID(ctx context.Context, req *transport.GetCategoryByIDRequest) (*transport.GetCategoryByIDResponse, error)
}
