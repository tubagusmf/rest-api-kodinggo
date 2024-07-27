package usecase

import (
	"context"
	"golang-rest-api-articles/internal/model"

	"github.com/go-playground/validator"
	"github.com/labstack/gommon/log"
	"github.com/sirupsen/logrus"
)

type ArticleUsecase struct {
	articleRepo model.IArticleRepository
}

var v = validator.New()

func NewArticleUsecase(
	articleRepo model.IArticleRepository,
) model.IArticleUsecase {
	return &ArticleUsecase{
		articleRepo: articleRepo,
	}
}

func (s *ArticleUsecase) FindAll(ctx context.Context, filter model.ArticleFilter) ([]*model.Article, error) {
	log := logrus.WithFields(logrus.Fields{
		"ctx":    ctx,
		"limit":  filter.Limit,
		"offset": filter.Offset,
	})

	articles, err := s.articleRepo.FindAll(ctx, filter)
	if err != nil {
		log.Error(err)
		return nil, err
	}

	return articles, nil
}

func (s *ArticleUsecase) FindById(ctx context.Context, id int64) (*model.Article, error) {
	log := logrus.WithFields(logrus.Fields{
		"ctx": ctx,
		"id":  id,
	})

	article, err := s.articleRepo.FindById(ctx, id)
	if err != nil {
		log.Error(err)
		return nil, err
	}
	return article, nil
}

func (s *ArticleUsecase) Create(ctx context.Context, in model.CreateArticleInput) error {
	log := logrus.WithFields(logrus.Fields{
		"ctx":     ctx,
		"title":   in.Title,
		"content": in.Content,
	})

	err := s.validateCreateArticleInput(ctx, in)
	if err != nil {
		log.Error(err)
		return err
	}

	article := model.Article{
		Title:   in.Title,
		Content: in.Content,
	}

	err = s.articleRepo.Create(ctx, article)
	if err != nil {
		log.Error(err)
		return err
	}
	return nil
}

func (s *ArticleUsecase) Update(ctx context.Context, in model.UpdateArticleInput) error {
	log := logrus.WithFields(logrus.Fields{
		"ctx":     ctx,
		"id":      in.Id,
		"title":   in.Title,
		"content": in.Content,
	})

	err := s.validateUpdateArticleInput(ctx, in)
	if err != nil {
		log.Error(err)
		return err
	}

	newArticle := model.Article{
		Id:          in.Id,
		Title:       in.Title,
		Content:     in.Content,
		PublishedAt: in.PublishedAt,
	}

	err = s.articleRepo.Update(ctx, newArticle)
	if err != nil {
		log.Error(err)
		return err
	}
	return nil
}

func (s *ArticleUsecase) Delete(ctx context.Context, id int64) error {
	log := logrus.WithFields(logrus.Fields{
		"ctx": ctx,
		"id":  id,
	})

	err := s.articleRepo.Delete(ctx, id)
	if err != nil {
		log.Error(err)
		return err
	}
	return nil
}

func (s *ArticleUsecase) validateCreateArticleInput(ctx context.Context, in model.CreateArticleInput) error {
	err := v.StructCtx(ctx, in)
	if err != nil {
		log.Error(err)
		return model.ErrInvalidInput
	}

	return nil
}

func (s *ArticleUsecase) validateUpdateArticleInput(ctx context.Context, in model.UpdateArticleInput) error {
	// err := v.Struct(in)
	// if err != nil {
	// 	log.Error(err)
	// 	return model.ErrInvalidInput
	// }

	err := v.StructCtx(ctx, in)
	if err != nil {
		log.Error(err)
		return model.ErrInvalidInput
	}

	// article, err := s.articleRepo.FindById(ctx, in.Id)
	// if err != nil {
	// 	log.Error(err)
	// 	return err
	// }
	// if in.PublishedAt.Before(article.CreatedAt) {
	// 	return model.ErrPublishedAtLessThanCreatedAt
	// }

	return nil
}
