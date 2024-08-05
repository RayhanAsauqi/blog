package repository

import (
	"context"

	"github.com/RayhanAsauqi/blog-app/internal/entity"
	"gorm.io/gorm"
)

type UserRepository interface {
	FindByUsername(ctx context.Context, username string) (entity.User, error)
	Create(ctx context.Context, user *entity.User) error
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{db}
}

func (r *userRepository) FindByUsername(ctx context.Context, username string) (entity.User, error) {
	var user entity.User
	err := r.db.WithContext(ctx).Where("username = ?", username).First(&user).Error
	return user, err
}

func (r *userRepository) Create(ctx context.Context, user *entity.User) error {
	return r.db.WithContext(ctx).Create(user).Error
}
