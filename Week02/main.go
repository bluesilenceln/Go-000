package main

import (
	"Go-000/Week02/dao"
	"Go-000/Week02/service"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	stopChan := make(chan os.Signal)
	signal.Notify(stopChan, syscall.SIGINT, syscall.SIGTERM, syscall.SIGKILL, syscall.SIGHUP)

	d := dao.NewDao("root:lxhroot@tcp(localhost:3306)/demo")
	_ = service.NewService(":80", d)

	<- stopChan
}
