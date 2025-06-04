package repository

import (
	"context"

	"github.com/lcsin/glib/ginlayout/internal/domain"
	"github.com/lcsin/glib/ginlayout/internal/repository/dao"
)

type IBookRepository interface {
	Create(ctx context.Context, book *domain.Book) error
	ModifyByID(ctx context.Context, book *domain.Book) error
	GetByID(ctx context.Context, id int64) (*domain.Book, error)
	GetByPage(ctx context.Context, page, pageSize int64) ([]*domain.Book, int64, error)
}

type BookRepository struct {
	dao dao.IBookDAO
}

func NewBookRepository(dao dao.IBookDAO) IBookRepository {
	return &BookRepository{dao: dao}
}

func (b *BookRepository) Create(ctx context.Context, book *domain.Book) error {
	return b.dao.Insert(ctx, book.ToModel())
}

func (b *BookRepository) ModifyByID(ctx context.Context, book *domain.Book) error {
	return b.dao.UpdateByID(ctx, book.ToModel())
}

func (b *BookRepository) GetByID(ctx context.Context, id int64) (*domain.Book, error) {
	book, err := b.dao.SelectByID(ctx, id)
	if err != nil {
		return nil, err
	}

	return &domain.Book{
		ID:          book.ID,
		Title:       book.Title,
		Author:      book.Author,
		Price:       int64(book.Price),
		PublishDate: book.PublishDate.UnixMilli(),
	}, err
}

func (b *BookRepository) GetByPage(ctx context.Context, page, pageSize int64) ([]*domain.Book, int64, error) {
	books, total, err := b.dao.SelectByPage(ctx, int(page), int(pageSize))
	if err != nil {
		return nil, 0, err
	}

	list := make([]*domain.Book, 0, len(books))
	for _, book := range books {
		list = append(list, &domain.Book{
			ID:          book.ID,
			Title:       book.Title,
			Author:      book.Author,
			Price:       int64(book.Price),
			PublishDate: book.PublishDate.UnixMilli(),
		})
	}
	return list, total, nil
}
