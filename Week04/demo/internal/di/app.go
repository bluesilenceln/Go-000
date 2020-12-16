package di

import "demo/internal/service"

type App struct {
	svc *service.Service
}

func NewApp(svc *service.Service) *App {
	app := new(App)
	app.svc = svc
	return app
}