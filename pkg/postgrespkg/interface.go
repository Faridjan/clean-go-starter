package postgrespkg

import (
	"context"

	"github.com/jackc/pgx/v4"
)

type ProviderPostgresInterface interface {
	StartTransaction(ctx context.Context, options pgx.TxOptions) (TxInterface, error)
	Migrate(migrationPath string) error
	CloseDB(ctx context.Context) error
}

type TxInterface interface {
	Commit(ctx context.Context) error
	Rollback(ctx context.Context) error
	Get(ctx context.Context, query string, args ...interface{}) pgx.Row
	Select(ctx context.Context, query string, args ...interface{}) (pgx.Rows, error)
	Exec(ctx context.Context, query string, args ...interface{}) error
}
