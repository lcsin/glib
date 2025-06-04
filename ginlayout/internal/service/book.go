package service

import (
	"context"

	"github.com/lcsin/glib/ginlayout/internal/domain"
	"github.com/lcsin/glib/ginlayout/internal/repository"
)

type IBookService interface {
	Add(ctx context.Context, book *domain.Book) error
	EditByID(ctx context.Context, book *domain.Book) error
	FindByID(ctx context.Context, id int64) (*domain.Book, error)
	FindByPage(ctx context.Context, page, pageSize int64) ([]*domain.Book, int64, error)
}

type BookService struct {
	repo repository.IBookRepository
}

func NewBookService(repo repository.IBookRepository) IBookService {
	return &BookService{repo: repo}
}

func (b *BookService) Add(ctx context.Context, book *domain.Book) error {
	return b.repo.Create(ctx, book)
}

func (b *BookService) EditByID(ctx context.Context, book *domain.Book) error {
	return b.repo.ModifyByID(ctx, book)
}

func (b *BookService) FindByID(ctx context.Context, id int64) (*domain.Book, error) {
	return b.repo.GetByID(ctx, id)
}

func (b *BookService) FindByPage(ctx context.Context, page, pageSize int64) ([]*domain.Book, int64, error) {
	return b.repo.GetByPage(ctx, page, pageSize)
}
