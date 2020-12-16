package main

import (
	"demo/internal/di"
	"flag"
	"os"
	"os/signal"
	"syscall"
	"time"
)
import "github.com/go-kratos/kratos/pkg/log"

func main() {
	flag.Parse()
	log.Init(nil)
	defer log.Close()
	log.Info("demo start")

	_, closeFunc, err := di.InitApp()
	if err != nil {
		panic(err)
	}

	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGHUP, syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT)
	for {
		s := <-c
		log.Info("get a signal %s", s.String())
		switch s {
		case syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT:
			closeFunc()
			log.Info("demo exit")
			time.Sleep(time.Second)
			return
		case syscall.SIGHUP:
		default:
			return
		}
	}
}
