package model

import (
	"context"
	"errors"
	"time"
)

var (
	ErrInvalidInput                 = errors.New("invalid input")
	ErrPublishedAtLessThanCreatedAt = errors.New("published_at less than created_at")
)

type IArticleRepository interface {
	FindAll(ctx context.Context, filter ArticleFilter) ([]*Article, error)
	FindById(ctx context.Context, id int64) (*Article, error)
	Create(ctx context.Context, article Article) error
	Update(ctx context.Context, article Article) error
	Delete(ctx context.Context, id int64) error
}

type IArticleUsecase interface {
	FindAll(ctx context.Context, filter ArticleFilter) ([]*Article, error)
	FindById(ctx context.Context, id int64) (*Article, error)
	Create(ctx context.Context, in CreateArticleInput) error
	Update(ctx context.Context, in UpdateArticleInput) error
	Delete(ctx context.Context, id int64) error
}

type Article struct {
	Id          int64      `json:"id"`
	Title       string     `json:"title"`
	Content     string     `json:"content"`
	PublishedAt *time.Time `json:"published_at"`
	CreatedAt   time.Time  `json:"created_at,omitempty"`
}

type ArticleFilter struct {
	Offset int32
	Limit  int32
}

type CreateArticleInput struct {
	Title   string `json:"title" validate:"required,min=4,max=50"`
	Content string `json:"content" validate:"required"`
}

type UpdateArticleInput struct {
	Id          int64      `json:"id"`
	Title       string     `json:"title" validate:"required"`
	Content     string     `json:"content"`
	PublishedAt *time.Time `json:"published_at"`
}
