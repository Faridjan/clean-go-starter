package page

import (
	"context"
	"tiny-template/src/api/page/transport"
)

type ServiceInterface interface {
	GetPageByID(ctx context.Context, req *transport.GetPageByIDRequest) (*transport.GetPageByIDResponse, error)
}
