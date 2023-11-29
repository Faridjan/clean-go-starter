package repository

import (
	"context"
	"tiny-template/pkg/postgrespkg"
	"tiny-template/src/domain"
)

type MainCategoryRepositoryInterface interface {
	StoreCategory(ctx context.Context, tr postgrespkg.TxInterface, category *domain.Category) error
	UpdateCategory(ctx context.Context, tr postgrespkg.TxInterface, category *domain.Category) error
	GetCategoryByID(ctx context.Context, tr postgrespkg.TxInterface) (uint32, error)
	FilterCategory(ctx context.Context, tr postgrespkg.TxInterface) ([]*domain.Category, error)
}

type MainPageRepositoryInterface interface {
	StorePage(ctx context.Context, tr postgrespkg.TxInterface, page *domain.Page) error
	UpdatePage(ctx context.Context, tr postgrespkg.TxInterface, page *domain.Page) error
	GetPageByID(ctx context.Context, tr postgrespkg.TxInterface) (uint32, error)
	FilterPage(ctx context.Context, tr postgrespkg.TxInterface) ([]*domain.Page, error)
}

type SearchCategoryRepositoryInterface interface {
	StoreCategory(ctx context.Context, category *domain.Category) error
	GetCategoryByID(ctx context.Context) (*domain.Category, error)
	DeleteCategoryByID(ctx context.Context, categoryID string) error
}

type SearchPageRepositoryInterface interface {
	StorePage(ctx context.Context, page *domain.Page) error
	GetPageByID(ctx context.Context) (*domain.Page, error)
	DeletePageByID(ctx context.Context, pageID string) error
}
