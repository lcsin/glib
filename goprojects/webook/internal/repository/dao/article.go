package dao

import (
	"context"
	"time"

	"github.com/lcsin/webook/internal/repository/model"
	"gorm.io/gorm"
)

type IArticleDAO interface {
	Insert(ctx context.Context, article model.Article) (int64, error)
	SelectByID(ctx context.Context, id int64) (*model.Article, error)
	SelectByUID(ctx context.Context, uid int64) ([]*model.Article, error)
	Update(ctx context.Context, article model.Article) error
	DeleteByID(ctx context.Context, id int64) error
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

func (a *ArticleDAO) Update(ctx context.Context, article model.Article) error {
	return a.db.WithContext(ctx).Where("id = ?", article.ID).UpdateColumns(map[string]interface{}{
		"title":        article.Title,
		"content":      article.Content,
		"updated_time": time.Now(),
	}).Error
}

func (a *ArticleDAO) DeleteByID(ctx context.Context, id int64) error {
	return a.db.WithContext(ctx).Where("id = ?", id).Delete(&model.Article{}).Error
}
