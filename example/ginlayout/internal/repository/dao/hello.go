package dao

import (
	"context"

	"ginlayout/internal/repository/dao/model"

	"gorm.io/gorm"
)

type IHelloDAO interface {
	SelectByID(ctx context.Context, id int64) (*model.Hello, error)
}

type HelloDAO struct {
	db *gorm.DB
}

func NewHelloDAO(db *gorm.DB) IHelloDAO {
	return &HelloDAO{db: db}
}

func (h *HelloDAO) SelectByID(ctx context.Context, id int64) (*model.Hello, error) {
	return &model.Hello{}, nil
}
