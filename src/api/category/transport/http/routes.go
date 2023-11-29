package http

import (
	"tiny-template/pkg"
	"tiny-template/src/api/category"

	"github.com/go-chi/chi"
	kithttp "github.com/go-kit/kit/transport/http"
)

func InitializeRoutes(mux *chi.Mux, endpoints *category.Endpoints, options []kithttp.ServerOption) {
	mux.Route("/tiny-template/v1/category", func(r chi.Router) {
		// swagger:route GET /tiny-template/v1/category Category CategoryStoreRequest
		// Category create/update
		// responses:
		//   200: CategoryStoreResponse
		r.Get("/get-category-json", pkg.HandlerJSON(endpoints.GetCategoryByID, []pkg.Middleware{
			pkg.CORS,
			pkg.SSOMiddlewareWrapper(pkg.SSOMiddlewareOptions{
				Actions:    []string{"get"},
				IsOptional: true,
			}),
		}))

		// swagger:route GET /tiny-template/v1/category Category CategoryStoreRequest
		// Category create/update
		// responses:
		//   200: CategoryStoreResponse
		r.Get("/get-category-xml", pkg.HandlerXML(endpoints.GetCategoryByID, []pkg.Middleware{
			pkg.CORS,
			pkg.SSOMiddlewareWrapper(pkg.SSOMiddlewareOptions{
				IsShared: true,
			}),
		}))
	})
}
