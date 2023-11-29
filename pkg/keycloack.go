package pkg

import (
	"net/http"

	"git.centerhome.kz/bcc/backend/toolchain/common-libs/middleware/keycloak"
)

type SSOMiddlewareOptions struct {
	Actions         []string
	OptionalActions []string
	IsOptional      bool
	IsShared        bool
}

func SSOMiddlewareWrapper(options SSOMiddlewareOptions) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return keycloak.SSOMiddleware(next, options.Actions, options.OptionalActions, options.IsOptional, options.IsShared)
	}
}
