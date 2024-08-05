package server

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"time"

	"github.com/RayhanAsauqi/blog-app/config"
	"github.com/RayhanAsauqi/blog-app/internal/http/handler"
	"github.com/RayhanAsauqi/blog-app/internal/http/router"
	"github.com/RayhanAsauqi/blog-app/internal/repository"
	"github.com/RayhanAsauqi/blog-app/internal/service"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type Server struct {
	*echo.Echo
	cfg *config.Config
}

func NewServer(cfg *config.Config, db *gorm.DB) *Server {
	e := echo.New()

	// Initialize repositories
	userRepository := repository.NewUserRepository(db)
	articleRepository := repository.NewArticleRepository(db)

	// Initialize services
	authService := service.NewAuthService(userRepository, cfg)
	articleService := service.NewArticleService(articleRepository)

	// Initialize handlers
	authHandler := handler.NewAuthHandler(authService)
	articleHandler := handler.NewArticleHandler(articleService)

	publicRoutes := router.PublicRoutes(authHandler, articleHandler)
	privateRoutes := router.PrivateRoutes(articleHandler)

	for _, r := range publicRoutes {
		e.Add(r.Method, r.Path, r.Handler)
	}
	for _, r := range privateRoutes {
		e.Add(r.Method, r.Path, r.Handler)
	}

	return &Server{e, cfg}
}

func (s *Server) Run() {
	go func() {
		err := s.Start(fmt.Sprintf(":%s", s.cfg.Port))
		if err != nil {
			log.Fatalf("Error starting server: %v", err)
		}
	}()

}

func (s *Server) GracefulShutdown() {
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)

	<-quit

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	go func() {
		if err := s.Shutdown(ctx); err != nil {
			s.Logger.Fatal(err)
		}
	}()
}
