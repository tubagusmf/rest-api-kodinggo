package repository

import (
	"context"
	"database/sql"
	"golang-rest-api-articles/internal/model"
)

type ArticleRepository struct {
	db *sql.DB
}

func NewArticleRepository(db *sql.DB) model.IArticleRepository {
	return &ArticleRepository{db: db}
}

func (s *ArticleRepository) FindAll(ctx context.Context, filter model.ArticleFilter) ([]*model.Article, error) {
	res, err := s.db.QueryContext(ctx, "SELECT id, title, content, published_at, created_at FROM articles LIMIT ? OFFSET ?", filter.Limit, filter.Offset)

	if err != nil {
		return nil, err
	}

	var articles []*model.Article
	for res.Next() {
		var article model.Article
		if err := res.Scan(&article.Id, &article.Title, &article.Content, &article.PublishedAt, &article.CreatedAt); err != nil {
			return nil, err
		}
		articles = append(articles, &article)
	}

	return articles, nil
}

func (s *ArticleRepository) FindById(ctx context.Context, id int64) (*model.Article, error) {
	res, err := s.db.QueryContext(ctx, "SELECT id, title, content, published_at, created_at FROM articles WHERE id=?", id)

	if err != nil {
		return nil, err
	}

	var article model.Article
	for res.Next() {
		if err := res.Scan(&article.Id, &article.Title, &article.Content, &article.PublishedAt, &article.CreatedAt); err != nil {
			return nil, err
		}
	}

	return &article, nil
}

func (s *ArticleRepository) Create(ctx context.Context, article model.Article) error {
	_, err := s.db.ExecContext(ctx, "INSERT INTO articles (title, content) VALUES (?, ?)", article.Title, article.Content)
	if err != nil {
		return err
	}
	return nil
}

func (s *ArticleRepository) Update(ctx context.Context, article model.Article) error {
	_, err := s.db.ExecContext(ctx, "UPDATE articles SET title=?, content=? WHERE id=?", article.Title, article.Content, article.Id)
	if err != nil {
		return err
	}
	return nil
}

func (s *ArticleRepository) Delete(ctx context.Context, id int64) error {
	_, err := s.db.ExecContext(ctx, "DELETE FROM articles WHERE id=?", id)
	if err != nil {
		return err
	}
	return nil
}
