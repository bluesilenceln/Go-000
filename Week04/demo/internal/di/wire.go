// +build wireinject
// The build tag makes sure the stub is not built in the final build.

package di

import (
	"demo/internal/biz"
	"demo/internal/data"
	"demo/internal/service"
	"github.com/google/wire"
)

//go:generate demo t wire
func InitService() (*service.Service, func(), error) {
	panic(wire.Build(data.Provider, biz.Provider, service.Provider))
}