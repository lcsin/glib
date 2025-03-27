package service

import (
	"context"

	"github.com/lcsin/webook/internal/domain"
	"github.com/lcsin/webook/internal/repository"
)

type IArticleService interface {
	Save(ctx context.Context, article domain.Article) (int64, error)
	Detail(ctx context.Context, id int64) (*domain.Article, error)
}

type ArticleService struct {
	repo repository.IArticleRepository
}

func NewArticleService(repo repository.IArticleRepository) IArticleService {
	return &ArticleService{repo: repo}
}

func (a *ArticleService) Save(ctx context.Context, article domain.Article) (int64, error) {
	if article.ID > 0 {
		err := a.repo.Update(ctx, article)
		return article.ID, err
	}
	return a.repo.Create(ctx, article)
}

func (a *ArticleService) Detail(ctx context.Context, id int64) (*domain.Article, error) {
	return a.repo.GetByID(ctx, id)
}
