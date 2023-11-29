package postgrespkg

import (
	"context"

	"github.com/jackc/pgx/v4"

	"github.com/jackc/pgx/v4/pgxpool"
)

type Transaction struct {
	dbCon *pgxpool.Pool
	tx    pgx.Tx
}

func NewTransaction(ctx context.Context, dbCon *pgxpool.Pool, options pgx.TxOptions) (*Transaction, error) {
	tx, err := dbCon.BeginTx(ctx, options)
	if err != nil {
		return nil, err
	}

	return &Transaction{
		dbCon: dbCon,
		tx:    tx,
	}, nil
}

func (tr *Transaction) Get(ctx context.Context, query string, args ...interface{}) pgx.Row {
	return nil
}
func (tr *Transaction) Select(ctx context.Context, query string, args ...interface{}) (pgx.Rows, error) {
	return nil, nil
}
func (tr *Transaction) Exec(ctx context.Context, query string, args ...interface{}) error {
	return nil
}

func (tr *Transaction) Rollback(ctx context.Context) error {
	return tr.tx.Rollback(ctx)
}

func (tr *Transaction) Commit(ctx context.Context) error {
	return tr.tx.Commit(ctx)
}
