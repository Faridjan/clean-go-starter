package postgrespkg

import (
	"context"
	"time"
)

type ConnectionConfig struct {
	MaxConnections  int
	MaxIdleTimeConn time.Duration

	DBHost string
	DBPort int
	DBName string
	DBUser string
	DBPass string
}

func NewPostgresProvider(ctx context.Context, cfg *ConnectionConfig) (ProviderPostgresInterface, error) {
	connPool, err := initConnect(ctx, cfg)
	if err != nil {
		return nil, err
	}

	return NewProvider(ctx, connPool), nil
}
