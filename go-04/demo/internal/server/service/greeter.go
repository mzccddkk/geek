package service

import (
	"context"
	v1 "hello/api/hello/v1"
	"hello/internal/server/usecase"
)

type GreeterService struct {
	uc *usecase.GreeterUsecase
}

func NewGreeterService(uc *usecase.GreeterUsecase) *GreeterService {
	return &GreeterService{uc: uc}
}

func (s *GreeterService) SayHello(ctx context.Context, in *v1.HelloRequest) (*v1.HelloReply, error) {
	if in.GetName() == "error" {
		return nil, v1.ErrorUserNotFound("user not found: %s", in.GetName())
	}
	return &v1.HelloReply{Message: "Hello " + in.GetName()}, nil
}
