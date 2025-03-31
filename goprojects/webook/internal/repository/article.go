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

	Publish(ctx context.Context, article domain.Article) (int64, error)
}

type ArticleRepository struct {
	writer dao.IArticleWriterDAO
	reader dao.IArticleReaderDAO
	user   dao.IUserDAO
}

func NewArticleRepository(writer dao.IArticleWriterDAO, reader dao.IArticleReaderDAO, user dao.IUserDAO) IArticleRepository {
	return &ArticleRepository{writer: writer, reader: reader, user: user}
}

func (a *ArticleRepository) Create(ctx context.Context, article domain.Article) (int64, error) {
	return a.writer.Insert(ctx, model.ArticleWriter{
		AuthorID: article.Author.ID,
		Title:    article.Title,
		Content:  article.Content,
	})
}

func (a *ArticleRepository) GetByID(ctx context.Context, id int64) (*domain.Article, error) {
	article, err := a.writer.SelectByID(ctx, id)
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
		Status:      domain.ArticleStatus(article.Status),
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
	articles, err := a.writer.SelectByUID(ctx, uid)
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
			Status:      domain.ArticleStatus(v.Status),
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
	return a.writer.UpdateByID(ctx, model.ArticleWriter{
		ID:       article.ID,
		AuthorID: article.Author.ID,
		Title:    article.Title,
		Content:  article.Content,
	})
}

func (a *ArticleRepository) DeleteByID(ctx context.Context, article domain.Article) error {
	err := a.writer.DeleteByID(ctx, model.ArticleWriter{
		ID:       article.ID,
		AuthorID: article.Author.ID,
	})
	err = a.reader.DeleteByID(ctx, model.ArticleReader{
		ID:       article.ID,
		AuthorID: article.Author.ID,
	})

	return err
}

func (a *ArticleRepository) Publish(ctx context.Context, article domain.Article) (int64, error) {
	var (
		id  int64
		err error
	)
	if article.ID == 0 {
		id, err = a.writer.Insert(ctx, model.ArticleWriter{
			AuthorID: article.Author.ID,
			Title:    article.Title,
			Content:  article.Content,
			Status:   article.Status.ToInt8(),
		})
	} else {
		err = a.writer.UpdateByID(ctx, model.ArticleWriter{
			ID:       article.ID,
			AuthorID: article.Author.ID,
			Title:    article.Title,
			Content:  article.Content,
			Status:   article.Status.ToInt8(),
		})
	}

	id, err = a.reader.Upset(ctx, model.ArticleReader{
		ID:       article.ID,
		AuthorID: article.Author.ID,
		Title:    article.Title,
		Content:  article.Content,
		Status:   article.Status.ToInt8(),
	})
	return id, err
}
