package repository

import (
	"context"

	"github.com/lcsin/webook/internal/domain"
	"github.com/lcsin/webook/internal/repository/dao"
	"github.com/lcsin/webook/internal/repository/model"
)

type IArticleRepository interface {
	Create(ctx context.Context, article domain.Article) (int64, error)
	GetByID(ctx context.Context, id int64) (*domain.Article, error)
	GetByUID(ctx context.Context, uid int64) ([]*domain.Article, error)
	Update(ctx context.Context, article domain.Article) error
	DeleteByID(ctx context.Context, article domain.Article) error
	PublishByID(ctx context.Context, article domain.Article) error
}

type ArticleRepository struct {
	article dao.IArticleDAO
	user    dao.IUserDAO
}

func NewArticleRepository(article dao.IArticleDAO, user dao.IUserDAO) IArticleRepository {
	return &ArticleRepository{article: article, user: user}
}

func (a *ArticleRepository) Create(ctx context.Context, article domain.Article) (int64, error) {
	return a.article.Insert(ctx, model.Article{
		AuthorID: article.Author.ID,
		Title:    article.Title,
		Content:  article.Content,
	})
}

func (a *ArticleRepository) GetByID(ctx context.Context, id int64) (*domain.Article, error) {
	article, err := a.article.SelectByID(ctx, id)
	if err != nil {
		return nil, err
	}
	user, err := a.user.SelectByID(ctx, article.AuthorID)
	if err != nil {
		return nil, err
	}

	return &domain.Article{
		ID:          article.ID,
		Title:       article.Title,
		Content:     article.Content,
		Status:      article.Status,
		CreatedTime: article.CreatedTime,
		UpdatedTime: article.UpdatedTime,
		PublishTime: article.PublishTime,
		Author: domain.Author{
			ID:   article.ID,
			Name: user.Username,
		},
	}, nil
}

func (a *ArticleRepository) GetByUID(ctx context.Context, uid int64) ([]*domain.Article, error) {
	articles, err := a.article.SelectByUID(ctx, uid)
	if err != nil {
		return nil, err
	}
	user, err := a.user.SelectByID(ctx, uid)
	if err != nil {
		return nil, err
	}

	list := make([]*domain.Article, 0, len(articles))
	for _, v := range articles {
		list = append(list, &domain.Article{
			ID:          v.ID,
			Title:       v.Title,
			Content:     v.Content,
			Status:      v.Status,
			CreatedTime: v.CreatedTime,
			UpdatedTime: v.UpdatedTime,
			PublishTime: v.PublishTime,
			Author: domain.Author{
				ID:   v.AuthorID,
				Name: user.Username,
			},
		})
	}
	return list, nil
}

func (a *ArticleRepository) Update(ctx context.Context, article domain.Article) error {
	return a.article.UpdateByID(ctx, model.Article{
		ID:       article.ID,
		AuthorID: article.Author.ID,
		Title:    article.Title,
		Content:  article.Content,
	})
}

func (a *ArticleRepository) DeleteByID(ctx context.Context, article domain.Article) error {
	return a.article.DeleteByID(ctx, model.Article{
		ID:       article.ID,
		AuthorID: article.Author.ID,
		Status:   article.Status,
	})
}

func (a *ArticleRepository) PublishByID(ctx context.Context, article domain.Article) error {
	return a.article.UpdateStatusByID(ctx, model.Article{
		ID:       article.ID,
		AuthorID: article.Author.ID,
		Status:   article.Status,
	})
}
