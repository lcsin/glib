package dao

import (
	"context"

	"github.com/lcsin/glib/ginlayout/internal/repository/gen/model"
	"github.com/lcsin/glib/ginlayout/internal/repository/gen/query"
	"gorm.io/gorm"
)

type IBookDAO interface {
	Insert(ctx context.Context, book *model.Book) error
	UpdateByID(ctx context.Context, book *model.Book) error
	SelectByID(ctx context.Context, id int64) (*model.Book, error)
	SelectByPage(ctx context.Context, page, pageSize int) ([]*model.Book, int64, error)
}

type BookDAO struct {
	db *gorm.DB
}

func NewBookDAO(db *gorm.DB) IBookDAO {
	return &BookDAO{db: db}
}

func (b *BookDAO) Insert(ctx context.Context, book *model.Book) error {
	return query.Book.WithContext(ctx).Create(book)
}

func (b *BookDAO) UpdateByID(ctx context.Context, book *model.Book) error {
	_, err := query.Book.WithContext(ctx).
		Where(query.Book.ID.Eq(book.ID)).
		Updates(&book)
	if err != nil {
		return err
	}
	return nil
}

func (b *BookDAO) SelectByID(ctx context.Context, id int64) (*model.Book, error) {
	return query.Book.WithContext(ctx).
		Where(query.Book.ID.Eq(id)).
		First()
}

func (b *BookDAO) SelectByPage(ctx context.Context, page, pageSize int) ([]*model.Book, int64, error) {
	return query.Book.WithContext(ctx).
		FindByPage((page-1)*pageSize, pageSize)
}
