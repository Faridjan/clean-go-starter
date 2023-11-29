package elastic

import (
	"context"
	"time"
	"tiny-template/src/domain"

	elasticLib "git.centerhome.kz/bcc/backend/toolchain/common-libs/databases/elastic"
)

const categoryIndex = "index"

type CategoryRepository struct {
	provider elasticLib.ProviderElasticInterface
}

func (r *CategoryRepository) StoreCategory(_ context.Context, category *domain.Category) error {
	category.UpdatedAt = time.Now().UTC()
	return r.provider.StoreByID(categoryIndex, category.ID, category)
}

func (r *CategoryRepository) GetCategoryByID(ctx context.Context) (*domain.Category, error) {
	return nil, nil
}

func (r *CategoryRepository) DeleteCategoryByID(ctx context.Context, categoryID string) error {
	return nil
}
