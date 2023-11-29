package postgres

import (
	"context"
	"tiny-template/pkg/postgrespkg"
	"tiny-template/src/domain"

	"git.centerhome.kz/bcc/backend/toolchain/common-libs/errors"
)

const (
	NotFoundCategory = "no rows in result set"
)

type CategoryRepository struct {
}

func (r *CategoryRepository) StoreCategory(ctx context.Context, tr postgrespkg.TxInterface, category *domain.Category) error {
	query := `INSERT INTO category(id, created_at, updated_at, slag) VALUES ($1, $2, $3, $4);`

	if err := tr.Exec(ctx, query,
		category.ID,
		category.CreatedAt,
		category.UpdatedAt,
		category.Slag,
	); err != nil {
		return errors.InternalServerError(ctx).SetDevMessage("StoreCategory.Exec.Insert: " + err.Error())
	}

	return nil
}

func (r *CategoryRepository) UpdateCategory(ctx context.Context, tr postgrespkg.TxInterface, category *domain.Category) error {
	return nil
}

func (r *CategoryRepository) GetCategoryByID(ctx context.Context, tr postgrespkg.TxInterface) (uint32, error) {
	return 0, nil
}

func (r *CategoryRepository) FilterCategory(ctx context.Context, tr postgrespkg.TxInterface) ([]*domain.Category, error) {
	return nil, nil
}
