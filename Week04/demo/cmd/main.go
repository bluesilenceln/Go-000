package main

import (
	pb "demo/api/user/v1"
	"demo/internal/di"
	"demo/pkg"
	"demo/pkg/transport/grpc"
	"flag"
	"github.com/go-kratos/kratos/pkg/log"
)

func main() {
	flag.Parse()
	log.Init(nil)
	defer log.Close()
	log.Info("demo start")

	app := pkg.NewApp()
	service, _, err := di.InitService()
	if err != nil {
		panic(err)
	}

	grpcSrv := grpc.NewServer("tcp", ":9000")

	pb.RegisterUserServer(grpcSrv, service)

	app.Append(pkg.Hook{OnStart: grpcSrv.Start, OnStop: grpcSrv.Stop})

	if err := app.Run(); err != nil {
		log.Error("%+v", err)
	}
}
