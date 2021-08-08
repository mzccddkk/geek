package repo

import (
	"context"
	"gorm.io/gorm"
	"hello/internal/server/usecase"
)

type greeterRepo struct {
	db *gorm.DB
}

func NewGreeterRepo(db *gorm.DB) usecase.GreeterRepo {
	return &greeterRepo{db: db}
}

func (r *greeterRepo) CreateGreeter(ctx context.Context, g *usecase.Greeter) error {
	return nil
}

func (r *greeterRepo) UpdateGreeter(ctx context.Context, g *usecase.Greeter) error {
	return nil
}
