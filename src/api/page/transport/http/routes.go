package http

import (
	"tiny-template/pkg"
	"tiny-template/src/api/page"

	"github.com/go-chi/chi"
	kithttp "github.com/go-kit/kit/transport/http"
)

func InitializeRoutes(mux *chi.Mux, endpoints *page.Endpoints, options []kithttp.ServerOption) {
	mux.Route("/tiny-template/v1/page", func(r chi.Router) {
		// swagger:route GET /tiny-template/v1/page Page PageStoreRequest
		// Page create/update
		// responses:
		//   200: PageStoreResponse
		r.Get("/get-page-json", pkg.HandlerJSON(endpoints.GetPageByID, []pkg.Middleware{
			pkg.CORS,
			pkg.SSOMiddlewareWrapper(pkg.SSOMiddlewareOptions{
				Actions:    []string{"get"},
				IsOptional: true,
			}),
		}))

		// swagger:route GET /tiny-template/v1/page Page PageStoreRequest
		// Page create/update
		// responses:
		//   200: PageStoreResponse
		r.Get("/get-page-xml", pkg.HandlerXML(endpoints.GetPageByID, []pkg.Middleware{
			pkg.CORS,
			pkg.SSOMiddlewareWrapper(pkg.SSOMiddlewareOptions{
				IsShared: true,
			}),
		}))
	})
}
