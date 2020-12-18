package http

import (
	"context"
	"net/http"
)

type Server struct {
	*http.Server
}

func NewServer(network, addr string) *Server {
	s := new(Server)
	return s
}

func (s *Server) Start(ctx context.Context) error {
	return nil
}

func (s *Server) Stop(ctx context.Context) error {
	return nil
}
