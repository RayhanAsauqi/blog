package repository

import (
	"context"

	"github.com/RayhanAsauqi/blog-app/internal/entity"
	"gorm.io/gorm"
)

type ArticleRepository interface {
	FindAll(ctx context.Context, page, limit int) ([]entity.Article, error)
	FindByID(ctx context.Context, id int64) (entity.Article, error)
	Create(ctx context.Context, article *entity.Article) error
	Update(ctx context.Context, article *entity.Article) error
	Delete(ctx context.Context, id int64) error
}

// articleRepository is a concrete implementation of ArticleRepository.
type articleRepository struct {
	db *gorm.DB
}

// NewArticleRepository creates a new instance of ArticleRepository.
func NewArticleRepository(db *gorm.DB) ArticleRepository {
	return &articleRepository{db}
}

// FindAll retrieves all articles from the data store.
func (r *articleRepository) FindAll(ctx context.Context, page, limit int) ([]entity.Article, error) {
	var articles []entity.Article
	err := r.db.WithContext(ctx).Find(&articles).Error
	return articles, err
}

// FindByID retrieves an article by its ID from the data store.
func (r *articleRepository) FindByID(ctx context.Context, id int64) (entity.Article, error) {
	var article entity.Article
	err := r.db.WithContext(ctx).First(&article, id).Error
	return article, err
}

// Create adds a new article to the data store.
func (r *articleRepository) Create(ctx context.Context, article *entity.Article) error {
	return r.db.WithContext(ctx).Create(article).Error
}

// Update modifies an existing article in the data store.
func (r *articleRepository) Update(ctx context.Context, article *entity.Article) error {
	return r.db.WithContext(ctx).Save(article).Error
}

// Delete removes an article by its ID from the data store.
func (r *articleRepository) Delete(ctx context.Context, id int64) error {
	return r.db.WithContext(ctx).Delete(&entity.Article{}, id).Error
}