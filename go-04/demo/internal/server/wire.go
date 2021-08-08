package server

import (
	"github.com/google/wire"
	"hello/internal/server/repo"
	"hello/internal/server/service"
	"hello/internal/server/usecase"
)

var ProviderSet = wire.NewSet(
	service.NewGreeterService,
	usecase.NewGreeterUsecase,
	repo.NewGreeterRepo,
)
