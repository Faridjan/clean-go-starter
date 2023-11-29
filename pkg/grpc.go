package pkg

import (
	"context"
	"net"

	"google.golang.org/grpc"
)

type GRPCServerInterface interface {
	Serve() error
	Stop(ctx context.Context) error
}

type GRPCServer struct {
	server   *grpc.Server
	listener net.Listener
}

func NewGRPCServer(server *grpc.Server, listener net.Listener) GRPCServerInterface {
	return &GRPCServer{
		server:   server,
		listener: listener,
	}
}

func (s GRPCServer) Serve() error {
	return s.server.Serve(s.listener)
}

func (s GRPCServer) Stop(_ context.Context) error {
	s.server.GracefulStop()
	return nil
}
