package postgrespkg

import (
	"context"
	"strconv"

	"github.com/jackc/pgx/v4/pgxpool"
)

func initConnect(ctx context.Context, cfg *ConnectionConfig) (*pgxpool.Pool, error) {
	dsn := "postgres://" +
		cfg.DBUser + ":" +
		cfg.DBPass + "@" +
		cfg.DBHost + ":" +
		strconv.Itoa(cfg.DBPort) + "/" +
		cfg.DBName +
		"?sslmode=disable" +
		"&pool_max_conns=" + strconv.Itoa(cfg.MaxConnections)

	pgxpoolCfg, err := pgxpool.ParseConfig(dsn)
	if err != nil {
		return nil, err
	}

	connPool, err := pgxpool.ConnectConfig(ctx, pgxpoolCfg)
	if err != nil {
		return nil, err
	}

	return connPool, nil
}
