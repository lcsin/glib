package domain

import (
	"time"

	"github.com/lcsin/glib/ginlayout/internal/repository/gen/model"
)

type Book struct {
	ID          int64
	Title       string
	Author      string
	Price       int64
	PublishDate int64
}

func (b *Book) ToModel() *model.Book {
	return &model.Book{
		ID:          b.ID,
		Title:       b.Title,
		Author:      b.Author,
		Price:       int32(b.Price),
		PublishDate: time.UnixMilli(b.PublishDate),
	}
}
