package di

import (
	"demo/internal/biz"
	"demo/internal/data"
	"demo/internal/service"
	"github.com/google/wire"
)

func InitApp() (*App, func(), error) {
	panic(wire.Build(data.Provider, biz.Provider, service.Provider, NewApp))
}