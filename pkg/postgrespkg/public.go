package postgrespkg

import (
	"context"
	"fmt"
	"strings"
	"tiny-template/pkg"
	"tiny-template/pkg/logger"

	"github.com/jackc/pgx/v4"

	"github.com/jackc/pgx/v4/pgxpool"
)

type ProviderPostgres struct {
	ctx      context.Context
	connPool *pgxpool.Pool
}

func NewProvider(ctx context.Context, connPool *pgxpool.Pool) ProviderPostgresInterface {
	return &ProviderPostgres{ctx: ctx, connPool: connPool}
}

func (provider *ProviderPostgres) StartTransaction(ctx context.Context, options pgx.TxOptions) (TxInterface, error) {
	return NewTransaction(ctx, provider.connPool, options)
}

func (provider *ProviderPostgres) Migrate(migrationPathDir string) error {
	ctx := context.Background()

	migrationPathDir = strings.TrimSuffix(migrationPathDir, "/") + "/"

	files, err := pkg.GetFiles(migrationPathDir, true)
	if err != nil {
		return err
	}

	tr, err := provider.connPool.BeginTx(ctx, pgx.TxOptions{IsoLevel: pgx.Serializable})
	if err != nil {
		return err
	}

	for key, fileName := range files {
		queryString, err := pkg.ReadFile(migrationPathDir + fileName)
		if err != nil {
			if errRollback := tr.Rollback(ctx); errRollback != nil {
				logger.Error(errRollback)
			}
			return err
		}

		if _, err := tr.Exec(provider.ctx, queryString); err != nil {
			if errRollback := tr.Rollback(ctx); errRollback != nil {
				logger.Error(errRollback)
			}
			return err
		}

		logger.Info(fmt.Sprintf("Migrated (%d/%d) %v%v", key+1, len(files), migrationPathDir, fileName))
	}

	return tr.Commit(ctx)
}

func (provider *ProviderPostgres) CloseDB(_ context.Context) error {
	provider.connPool.Close()

	return nil
}
