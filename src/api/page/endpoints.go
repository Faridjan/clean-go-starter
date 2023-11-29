package page

import (
	"context"
	"tiny-template/src/api/page/transport"

	"git.centerhome.kz/bcc/backend/toolchain/common-libs/errors"
	"github.com/go-kit/kit/endpoint"
)

type Endpoints struct {
	GetPageByID endpoint.Endpoint
}

func NewEndpoints(s ServiceInterface) *Endpoints {
	return &Endpoints{
		GetPageByID: makeStoreAdvertEndpoint(s),
	}
}

func makeStoreAdvertEndpoint(s ServiceInterface) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		ctx = errors.SetEndpointEnumToCtx(ctx, "01")
		req := request.(transport.GetPageByIDRequest)
		return s.GetPageByID(ctx, &req)
	}
}
