package app

import (
	"context"
	"tiny-template/pkg/logger"

	"tiny-template/pkg"
)

func Run(ctx context.Context, httpPort, grpcPort string) {
	graceful := pkg.NewGracefullyShutdown(ctx)

	// Init Providers...
	postgresProvider := InitPostgresProvider(ctx)
	graceful.MustClose(postgresProvider.CloseDB)

	elasticProvider := InitElasticProvider(ctx)
	graceful.MustClose(elasticProvider.Close)

	// Init App
	appRepositories := InitRepositories()
	appServices := InitServices(appRepositories)
	appEndpoints := InitEndpoints(appServices)

	// Run ...
	httpServer := InitHTTP(appEndpoints, httpPort)
	graceful.Go(httpServer.ListenAndServe)
	graceful.MustClose(httpServer.Shutdown)
	logger.Info("HTTP server listening on " + httpPort + "...")

	grpcServer := InitGRPC(appEndpoints, grpcPort)
	graceful.Go(grpcServer.Serve)
	graceful.MustClose(grpcServer.Stop)
	logger.Info("gRPC server listening on " + grpcPort + "...")

	graceful.Wait()
}
