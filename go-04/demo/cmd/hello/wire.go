// +build wireinject

package main

import (
	"github.com/google/wire"
	"hello/internal/server"
)

func initApp() (func(), error) {
	panic(wire.Build(server.ProviderSet, newApp))
}
