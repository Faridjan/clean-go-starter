package app

import (
	"context"
	"time"
	"tiny-template/config"
	"tiny-template/pkg/elasticpkg"
	"tiny-template/pkg/logger"
	"tiny-template/pkg/postgrespkg"
)

const (
	PostgresIdleTimeConnection = 30 * time.Minute
	MigrationPath              = "./migrations"
)

func InitPostgresProvider(ctx context.Context) postgrespkg.ProviderPostgresInterface {
	cfg := config.GetConfig()

	// Postgres Provider:
	postgresProvider, err := postgrespkg.NewPostgresProvider(ctx, &postgrespkg.ConnectionConfig{
		MaxConnections:  cfg.DB.MaxConnections,
		MaxIdleTimeConn: PostgresIdleTimeConnection,
		DBHost:          cfg.DB.Host,
		DBPort:          cfg.DB.Port,
		DBName:          cfg.DB.Name,
		DBUser:          cfg.DB.User,
		DBPass:          cfg.DB.Pass,
	})
	if err != nil {
		logger.Fatal(err.Error())
	}

	if err := postgresProvider.Migrate(MigrationPath); err != nil {
		logger.Fatal(err.Error())
	}

	return postgresProvider
}

func InitElasticProvider(ctx context.Context) elasticpkg.ProviderElasticInterface {
	cfg := config.GetConfig()

	// Elastic Provider:
	elasticProvider, err := elasticpkg.NewElasticProvider(
		ctx,
		cfg.ElasticConfig.ConnectionURL,
		cfg.ElasticConfig.SetSniff,
		cfg.ElasticConfig.Username,
		cfg.ElasticConfig.Password,
	)
	if err != nil {
		logger.Fatal(err.Error())
	}

	return elasticProvider
}
