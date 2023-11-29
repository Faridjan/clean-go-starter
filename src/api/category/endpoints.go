package category

import (
	"context"
	"tiny-template/src/api/category/transport"

	"git.centerhome.kz/bcc/backend/toolchain/common-libs/errors"
	"github.com/go-kit/kit/endpoint"
)

type Endpoints struct {
	GetCategoryByID endpoint.Endpoint
}

func NewEndpoints(s ServiceInterface) *Endpoints {
	return &Endpoints{
		GetCategoryByID: makeStoreAdvertEndpoint(s),
	}
}

func makeStoreAdvertEndpoint(s ServiceInterface) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		ctx = errors.SetEndpointEnumToCtx(ctx, "01")
		req := request.(transport.GetCategoryByIDRequest)
		return s.GetCategoryByID(ctx, &req)
	}
}
