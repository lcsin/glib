package dao

import (
	"context"
	"fmt"
	"time"

	"github.com/lcsin/webook/internal/repository/model"
	"gorm.io/gorm"
)

type IArticleWriterDAO interface {
	Insert(ctx context.Context, article model.ArticleWriter) (int64, error)
	SelectByID(ctx context.Context, id int64) (*model.ArticleWriter, error)
	SelectByUID(ctx context.Context, uid int64) ([]*model.ArticleWriter, error)
	UpdateByID(ctx context.Context, article model.ArticleWriter) error
	DeleteByID(ctx context.Context, article model.ArticleWriter) error
	UpdateStatusByID(ctx context.Context, article model.ArticleWriter) error
}

type ArticleWriterDAO struct {
	db *gorm.DB
}

func NewArticleWriterDAO(db *gorm.DB) IArticleWriterDAO {
	return &ArticleWriterDAO{db: db}
}

func (a *ArticleWriterDAO) Insert(ctx context.Context, article model.ArticleWriter) (int64, error) {
	art := model.ArticleWriter{
		AuthorID:    article.AuthorID,
		Title:       article.Title,
		Content:     article.Content,
		Status:      article.Status,
		CreatedTime: time.Now().UnixMilli(),
		PublishTime: article.PublishTime,
	}
	err := a.db.WithContext(ctx).Create(&art).Error
	return art.ID, err
}

func (a *ArticleWriterDAO) SelectByID(ctx context.Context, id int64) (*model.ArticleWriter, error) {
	var article model.ArticleWriter
	if err := a.db.WithContext(ctx).Where("id = ?", id).First(&article).Error; err != nil {
		return nil, err
	}
	return &article, nil
}

func (a *ArticleWriterDAO) SelectByUID(ctx context.Context, uid int64) ([]*model.ArticleWriter, error) {
	var articles []*model.ArticleWriter
	if err := a.db.WithContext(ctx).Where("uid = ?", uid).Find(&articles).Error; err != nil {
		return nil, err
	}
	return articles, nil
}

func (a *ArticleWriterDAO) UpdateByID(ctx context.Context, article model.ArticleWriter) error {
	res := a.db.WithContext(ctx).Model(&model.ArticleWriter{}).
		Where("id = ? and author_id = ?", article.ID, article.AuthorID).
		UpdateColumns(map[string]interface{}{
			"title":        article.Title,
			"content":      article.Content,
			"status":       article.Status,
			"updated_time": time.Now().UnixMilli(),
			"publish_time": article.PublishTime,
		})
	if res.RowsAffected == 0 {
		return fmt.Errorf("更新结果为0. id=%d author_id=%d", article.ID, article.AuthorID)
	}
	return res.Error
}

func (a *ArticleWriterDAO) DeleteByID(ctx context.Context, article model.ArticleWriter) error {
	return a.db.WithContext(ctx).
		Where("id = ? and author_id = ?", article.ID, article.AuthorID).
		Delete(&model.ArticleWriter{}).Error
}

func (a *ArticleWriterDAO) UpdateStatusByID(ctx context.Context, article model.ArticleWriter) error {
	res := a.db.WithContext(ctx).Model(&model.ArticleWriter{}).
		Where("id = ? and author_id = ?", article.ID, article.AuthorID).
		UpdateColumns(map[string]interface{}{
			"status":       article.Status,
			"updated_time": time.Now().UnixMilli(),
		})
	if res.RowsAffected == 0 {
		return fmt.Errorf("更新结果为0. id=%d author_id=%d", article.ID, article.AuthorID)
	}
	return res.Error
}
