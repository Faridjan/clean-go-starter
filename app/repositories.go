package app

import (
	"tiny-template/src/repository"
	"tiny-template/src/repository/elastic"
	"tiny-template/src/repository/postgres"
)

type Repositories struct {
	// Main Store:
	MainCategoryRepository repository.MainCategoryRepositoryInterface
	MainPageRepository     repository.MainPageRepositoryInterface
	// Search Store:
	SearchCategoryRepository repository.SearchCategoryRepositoryInterface
	SearchPageRepository     repository.SearchPageRepositoryInterface
}

func InitRepositories() *Repositories {
	return &Repositories{
		MainCategoryRepository:   &postgres.CategoryRepository{},
		MainPageRepository:       &postgres.PageRepository{},
		SearchCategoryRepository: &elastic.CategoryRepository{},
		SearchPageRepository:     &elastic.PageRepository{},
	}
}
