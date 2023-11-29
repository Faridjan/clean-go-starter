package main

import (
	"context"
	"flag"
	"tiny-template/app"

	"tiny-template/config"
	"tiny-template/pkg/logger"

	"git.centerhome.kz/bcc/backend/toolchain/common-libs/errors"
)

const (
	projectCode = "01"
	serviceCode = "01"
	serviceName = "tiny-template"
)

func main() {
	// Ballast of 10 GiB (TODO/READ https://medium.com/clean-code-channel/go-memory-ballast-dec0c04830b1)
	_ = make([]byte, 10<<30)

	ctx := context.Background()
	errors.SetErrCodePrefix(projectCode, serviceCode)

	// Parse flags
	httpPort := flag.String("http.port", "8080", "HTTP listen address")
	grpcPort := flag.String("grpc.port", "50051", "gRPC listen address")
	flag.Parse()

	// Init logger
	logger.InitLogger(serviceName)
	logger.Success("Service started")

	// Init config
	config.InitConfigs()
	logger.Success("Initialized configs")

	// App Running...
	app.Run(ctx, *httpPort, *grpcPort)
}
