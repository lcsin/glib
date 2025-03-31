package dao

import (
	"context"
	"time"

	"github.com/lcsin/webook/internal/repository/model"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type IArticleReaderDAO interface {
	Upset(ctx context.Context, article model.ArticleReader) (int64, error)
	DeleteByID(ctx context.Context, article model.ArticleReader) error
}

type ArticleReaderDAO struct {
	db *gorm.DB
}

func NewArticleReaderDAO(db *gorm.DB) IArticleReaderDAO {
	return &ArticleReaderDAO{db: db}
}

func (a *ArticleReaderDAO) Upset(ctx context.Context, article model.ArticleReader) (int64, error) {
	ts := time.Now().UnixMilli()
	article.CreatedTime = ts
	article.UpdatedTime = ts
	article.PublishTime = ts

	err := a.db.Clauses(clause.OnConflict{
		DoUpdates: clause.Assignments(map[string]interface{}{
			"status":       article.Status,
			"updated_time": ts,
			"publish_time": ts,
		}),
	}).Create(&article).Error
	return article.ID, err
}

func (a *ArticleReaderDAO) DeleteByID(ctx context.Context, article model.ArticleReader) error {
	return a.db.WithContext(ctx).
		Where("id = ? and author_id = ?", article.ID, article.AuthorID).
		Delete(&model.ArticleReader{}).Error
}
