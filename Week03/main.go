package main

import (
	"context"
	"fmt"
	"golang.org/x/sync/errgroup"
	"net/http"
	_ "net/http/pprof"
	"os"
	"os/signal"
	"syscall"
)

type Serve struct {
	addr string
	s *http.Server
}

func New(addr string, handle http.Handler) *Serve {
	serve := new(Serve)
	serve.addr = addr
	serve.s = &http.Server{
		Addr: addr,
		Handler: handle,
	}

	return serve
}

func (s *Serve) Start(ctx context.Context, stop chan os.Signal) error {
	go func() {
		select {
		case sig := <- stop:
			if sig != nil {
				fmt.Printf("%s signal: %+v\n", s.addr, sig)
			}
		case <- ctx.Done():
			fmt.Printf("%s done\n", s.addr)
		}
		_ = s.s.Shutdown(context.Background())
	}()

	return s.s.ListenAndServe()
}

func main() {
	stop := make(chan os.Signal)
	signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM, syscall.SIGHUP)

	g, ctx := errgroup.WithContext(context.Background())
	serv := New("0.0.0.0:8081", http.DefaultServeMux)
	servDebug := New("0.0.0.0:8081", http.DefaultServeMux)

	g.Go(func() error {
		return serv.Start(ctx, stop)
	})

	g.Go(func() error {
		return servDebug.Start(ctx, stop)
	})

	if err := g.Wait(); err != nil {
		fmt.Printf("%+v\n", err)
	}
}