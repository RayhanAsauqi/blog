package service

import (
	"context"
	"errors"
	"time"

	"github.com/RayhanAsauqi/blog-app/config"
	"github.com/RayhanAsauqi/blog-app/internal/dto"
	"github.com/RayhanAsauqi/blog-app/internal/entity"
	"github.com/RayhanAsauqi/blog-app/internal/repository"
	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
)

type AuthService interface {
	Register(ctx context.Context, request dto.RegisterRequest) error
	Login(ctx context.Context, request dto.LoginRequest) (dto.LoginResponse, error) 
}

type authService struct {
	repository repository.UserRepository
	cfg        *config.Config
}

func NewAuthService(repository repository.UserRepository, cfg *config.Config) AuthService {
	return &authService{repository, cfg}
}

func (s *authService) Register(ctx context.Context, request dto.RegisterRequest) error {
	if request.Password != request.ConfirmPassword {
		return errors.New("passwords do not match") // **[UPDATED]** Changed error message
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(request.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	user := entity.User{
		Username: request.Username,
		Email:    request.Email,
		Password: string(hash),
		Role:     request.Role, 
	}

	return s.repository.Create(ctx, &user)
}

func (s *authService) Login(ctx context.Context, request dto.LoginRequest) (dto.LoginResponse, error) {
	user, err := s.repository.FindByUsername(ctx, request.Username)
	if err != nil {
		return dto.LoginResponse{}, errors.New("invalid credentials") // **[UPDATED]** Return an empty LoginResponse on error
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(request.Password))
	if err != nil {
		return dto.LoginResponse{}, errors.New("invalid credentials") // **[UPDATED]** Return an empty LoginResponse on error
	}

	if user.Role == "" { // Check if role is empty
		return dto.LoginResponse{}, errors.New("user role is not assigned")
	}

	token, err := s.generateToken(user.Username, user.Role)
	if err != nil {
		return dto.LoginResponse{}, err
	}

	return dto.LoginResponse{Token: token, Role: user.Role}, nil
}

func (s *authService) generateToken(username, role string) (string, error) {
	claims := jwt.MapClaims{
		"username": username,
		"role":     role, // **[ADDED]** Include role in token claims
		"exp":      time.Now().Add(time.Hour * 24).Unix(),
		"iat":      time.Now().Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims) // **[UPDATED]** Changed to HS256 for signing method
	return token.SignedString([]byte(s.cfg.JWTSecret))
}