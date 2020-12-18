package grpc

import (
	"context"
	"google.golang.org/grpc"
	"net"
)

type Server struct {
	*grpc.Server
	network string
	addr string
}

func NewServer(network, addr string) *Server {
	s := new(Server)
	s.network = network
	s.addr = addr
	s.Server = grpc.NewServer()
	return s
}

func (s *Server) Start(ctx context.Context) error {
	lis, err := net.Listen(s.network, s.addr)
	if err != nil {
		return err
	}
	return s.Serve(lis)
}

// Stop stop the gRPC server.
func (s *Server) Stop(ctx context.Context) error {
	s.GracefulStop()
	return nil
}