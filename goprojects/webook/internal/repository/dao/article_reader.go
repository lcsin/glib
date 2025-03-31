package dao

import (
	"context"
	"fmt"
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
	res := a.db.WithContext(ctx).Model(&model.ArticleWriter{}).
		Where("id = ? and author_id = ?", article.ID, article.AuthorID).
		UpdateColumns(map[string]interface{}{
			"status":       article.Status,
			"deleted":      true,
			"deleted_time": time.Now().UnixMilli(),
		})
	if res.RowsAffected == 0 {
		return fmt.Errorf("更新结果为0. id=%d author_id=%d", article.ID, article.AuthorID)
	}
	return res.Error
}
