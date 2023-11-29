package app

import (
	middlewareCategory "tiny-template/src/api/category"
	middlewarePage "tiny-template/src/api/page"
)

type Endpoints struct {
	Category *middlewareCategory.Endpoints
	Page     *middlewarePage.Endpoints
}

func InitEndpoints(services *Services) *Endpoints {
	return &Endpoints{
		Category: middlewareCategory.NewEndpoints(services.CategoryService),
		Page:     middlewarePage.NewEndpoints(services.PageService),
	}
}
