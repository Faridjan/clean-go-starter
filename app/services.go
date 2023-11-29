package app

import (
	"tiny-template/src/api/category"
	serviceCategory "tiny-template/src/api/category/service"
	"tiny-template/src/api/page"
	servicePage "tiny-template/src/api/page/service"
)

type Services struct {
	CategoryService category.ServiceInterface
	PageService     page.ServiceInterface
}

func InitServices(repositories *Repositories) *Services {
	return &Services{
		CategoryService: serviceCategory.NewService(repositories.MainPageRepository, repositories.SearchCategoryRepository),
		PageService:     servicePage.NewService(repositories.MainPageRepository, repositories.SearchPageRepository),
	}
}
