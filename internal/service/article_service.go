package service

import (
	"context"
	"time"

	"github.com/RayhanAsauqi/blog-app/internal/dto"
	"github.com/RayhanAsauqi/blog-app/internal/entity"
	"github.com/RayhanAsauqi/blog-app/internal/repository"
)

type ArticleService interface {
	FindAll(ctx context.Context) ([]dto.Article, error)
	FindByID(ctx context.Context, id int64) (dto.Article, error)
	Create(ctx context.Context, request dto.CreateArticleRequest) error
	Update(ctx context.Context, request dto.UpdateArticleRequest) error
	Delete(ctx context.Context, id int64) error
}

type articleService struct {
	repository repository.ArticleRepository
}

func NewArticleService(repository repository.ArticleRepository) ArticleService {
	return &articleService{repository}
}

func (a *articleService) FindAll(ctx context.Context) ([]dto.Article, error) {
	articles, err := a.repository.FindAll(ctx, 0, 0) // Added default pagination parameters
	if err != nil {
		return nil, err
	}

	articlesDTO := make([]dto.Article, len(articles))
	for i, article := range articles {
		articlesDTO[i] = dto.Article{
			ID:      article.ID,
			Title:   article.Title,
			Author:  article.Author,
			Content: article.Content,
		}
	}
	return articlesDTO, nil
}

func (a *articleService) FindByID(ctx context.Context, id int64) (dto.Article, error) {
	article, err := a.repository.FindByID(ctx, id)
	if err != nil {
		return dto.Article{}, err
	}

	return dto.Article{ID: article.ID, Title: article.Title, Content: article.Content, Author: article.Author, CreatedAt: article.CreatedAt}, nil
}

func (a *articleService) Create(ctx context.Context, request dto.CreateArticleRequest) error {
	article := entity.Article{
		Title:     request.Title,
		Content:   request.Content,
		Author:    request.Author,
		CreatedAt: time.Now(),
	}
	// Set CreatedAt to the current time
	article.CreatedAt = time.Now()

	return a.repository.Create(ctx, &article)
}

func (a *articleService) Update(ctx context.Context, request dto.UpdateArticleRequest) error {
	article, err := a.repository.FindByID(ctx, request.ID)
	if err != nil {
		return err
	}

	if request.Title != "" {
		article.Title = request.Title
	}

	if request.Content != "" {
		article.Content = request.Content
	}

	if request.Author != "" {
		article.Author = request.Author
	}

	return a.repository.Update(ctx, &article)
}

func (a *articleService) Delete(ctx context.Context, id int64) error {
	return a.repository.Delete(ctx, id)
}