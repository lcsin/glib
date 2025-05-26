package service

import (
	"context"

	"ginlayout/internal/domain"
	"ginlayout/internal/repository"
)

type IHelloService interface {
	SayHello(ctx context.Context, id int64) (*domain.Hello, error)
}

type HelloService struct {
	repo repository.IHelloRepository
}

func NewHelloService(repo repository.IHelloRepository) IHelloService {
	return &HelloService{repo: repo}
}

func (h HelloService) SayHello(ctx context.Context, id int64) (*domain.Hello, error) {
	return &domain.Hello{}, nil
}
