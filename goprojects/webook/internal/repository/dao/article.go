package dao

import (
	"context"
	"fmt"
	"time"

	"github.com/lcsin/webook/internal/repository/model"
	"gorm.io/gorm"
)

type IArticleDAO interface {
	Insert(ctx context.Context, article model.Article) (int64, error)
	SelectByID(ctx context.Context, id int64) (*model.Article, error)
	SelectByUID(ctx context.Context, uid int64) ([]*model.Article, error)
	UpdateByID(ctx context.Context, article model.Article) error
	DeleteByID(ctx context.Context, article model.Article) error
	UpdateStatusByID(ctx context.Context, article model.Article) error
}

type ArticleDAO struct {
	db *gorm.DB
}

func NewArticleDAO(db *gorm.DB) IArticleDAO {
	return &ArticleDAO{db: db}
}

func (a *ArticleDAO) Insert(ctx context.Context, article model.Article) (int64, error) {
	art := model.Article{
		AuthorID:    article.AuthorID,
		Title:       article.Title,
		Content:     article.Content,
		CreatedTime: time.Now().UnixMilli(),
	}
	err := a.db.WithContext(ctx).Create(&art).Error
	return art.ID, err
}

func (a *ArticleDAO) SelectByID(ctx context.Context, id int64) (*model.Article, error) {
	var article model.Article
	if err := a.db.WithContext(ctx).Where("id = ?", id).First(&article).Error; err != nil {
		return nil, err
	}
	return &article, nil
}

func (a *ArticleDAO) SelectByUID(ctx context.Context, uid int64) ([]*model.Article, error) {
	var articles []*model.Article
	if err := a.db.WithContext(ctx).Where("uid = ?", uid).Find(&articles).Error; err != nil {
		return nil, err
	}
	return articles, nil
}

func (a *ArticleDAO) UpdateByID(ctx context.Context, article model.Article) error {
	res := a.db.WithContext(ctx).Model(&model.Article{}).
		Where("id = ? and author_id = ?", article.ID, article.AuthorID).
		UpdateColumns(map[string]interface{}{
			"title":        article.Title,
			"content":      article.Content,
			"updated_time": time.Now().UnixMilli(),
		})
	if res.RowsAffected == 0 {
		return fmt.Errorf("更新结果为0. id=%d author_id=%d", article.ID, article.AuthorID)
	}
	return res.Error
}

func (a *ArticleDAO) DeleteByID(ctx context.Context, article model.Article) error {
	res := a.db.WithContext(ctx).Model(&model.Article{}).
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

func (a *ArticleDAO) UpdateStatusByID(ctx context.Context, article model.Article) error {
	res := a.db.WithContext(ctx).Model(&model.Article{}).
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
