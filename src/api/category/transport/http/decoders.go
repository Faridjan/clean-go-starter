package http

import (
	"context"
	"net/http"
)

func specificDecoder(ctx context.Context, r *http.Request) (interface{}, error) {
	return nil, nil
}
