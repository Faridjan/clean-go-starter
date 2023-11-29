package postgres

import (
	"context"
	"tiny-template/pkg/postgrespkg"
	"tiny-template/src/domain"

	"git.centerhome.kz/bcc/backend/toolchain/common-libs/errors"
)

const (
	NotFoundPage = "no rows in result set"
)

type PageRepository struct {
}

func (r *PageRepository) StorePage(ctx context.Context, tr postgrespkg.TxInterface, page *domain.Page) error {
	query := `INSERT INTO page(id, created_at, updated_at, category_id, title, text) VALUES ($1, $2, $3, $4, $5, $6);`

	if err := tr.Exec(ctx, query,
		page.ID,
		page.CreatedAt,
		page.UpdatedAt,
		page.CategoryID,
		page.Title,
		page.Text,
	); err != nil {
		return errors.InternalServerError(ctx).SetDevMessage("StorePage.Exec.Insert: " + err.Error())
	}

	return nil
}

func (r *PageRepository) UpdatePage(ctx context.Context, tr postgrespkg.TxInterface, page *domain.Page) error {
	return nil
}

func (r *PageRepository) GetPageByID(ctx context.Context, tr postgrespkg.TxInterface) (uint32, error) {
	return 0, nil
}

func (r *PageRepository) FilterPage(ctx context.Context, tr postgrespkg.TxInterface) ([]*domain.Page, error) {
	return nil, nil
}
