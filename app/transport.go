package app

import (
	"net"
	"net/http"
	"time"
	"tiny-template/pkg"

	"google.golang.org/grpc"

	"git.centerhome.kz/bcc/backend/toolchain/common-libs/utils/healthcheck"
	"github.com/prometheus/client_golang/prometheus/promhttp"

	httpencoders "git.centerhome.kz/bcc/backend/toolchain/common-libs/transport/http"
	"github.com/go-chi/chi"
	kittransport "github.com/go-kit/kit/transport"
	kithttp "github.com/go-kit/kit/transport/http"

	"tiny-template/pkg/logger"
	categoryHttp "tiny-template/src/api/category/transport/http"
	pageHttp "tiny-template/src/api/page/transport/http"
)

func InitHTTP(endpoints *Endpoints, port string) *http.Server {
	serverOptions := []kithttp.ServerOption{
		kithttp.ServerErrorEncoder(httpencoders.EncodeErrorResponse),
		kithttp.ServerErrorHandler(kittransport.NewLogErrorHandler(logger.GetLogger())),
	}

	mux := chi.NewRouter()

	categoryHttp.InitializeRoutes(mux, endpoints.Category, serverOptions)
	pageHttp.InitializeRoutes(mux, endpoints.Page, serverOptions)

	mux.Get("/metrics", promhttp.Handler().ServeHTTP)
	mux.Get("/check", healthcheck.HealthCheck)

	return &http.Server{
		Addr:         ":" + port,
		Handler:      mux,
		ReadTimeout:  30 * time.Second,
		WriteTimeout: 30 * time.Second,
	}
}

func InitGRPC(endpoints *Endpoints, port string) pkg.GRPCServerInterface {
	grpcListener, err := net.Listen("tcp", ":"+port)
	if err != nil {
		logger.Fatal("gRPC listen err", err.Error())
	}

	serverGRPC := grpc.NewServer()
	// ... realize endpoints for gRPC
	// pb.RegisterGreeterServer(grpc, &gSrv)

	return pkg.NewGRPCServer(serverGRPC, grpcListener)
}
