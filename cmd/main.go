package main

import (
	"log"

	"github.com/RayhanAsauqi/blog-app/config"
	"github.com/RayhanAsauqi/blog-app/internal/entity"
	"github.com/RayhanAsauqi/blog-app/pkg/database"
	"github.com/RayhanAsauqi/blog-app/pkg/server"
)

func main() {
	// Load configuration
	cfg, err := config.NewConfig(".env")
	if err != nil {
		log.Fatalf("Failed to load configuration: %v", err)
	}

	// Connect to the database
	db, err := database.ConnectMysql(cfg)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	// Auto migrate the database
	err = db.AutoMigrate(&entity.Article{}, &entity.User{})
	if err != nil {
		log.Fatalf("Failed to migrate database: %v", err)
	}

	// Create and run the server
	server := server.NewServer(cfg, db)
	server.Run()

	// Handle graceful shutdown
	server.GracefulShutdown()
}
