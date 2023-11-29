package elastic

import (
	"context"
	"time"
	"tiny-template/src/domain"

	elasticLib "git.centerhome.kz/bcc/backend/toolchain/common-libs/databases/elastic"
)

const pageIndex = "index"

type PageRepository struct {
	provider elasticLib.ProviderElasticInterface
}

func (r *PageRepository) StorePage(_ context.Context, page *domain.Page) error {
	page.UpdatedAt = time.Now().UTC()
	return r.provider.StoreByID(pageIndex, page.ID, page)
}

func (r *PageRepository) GetPageByID(ctx context.Context) (*domain.Page, error) {
	return nil, nil
}

func (r *PageRepository) DeletePageByID(ctx context.Context, pageID string) error {
	return nil
}
