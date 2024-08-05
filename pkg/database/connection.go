package database

import (
	"fmt"

	"github.com/RayhanAsauqi/blog-app/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func ConnectMysql(cfg *config.Config) (*gorm.DB, error) {
	// Debug log to check configuration values
	fmt.Println("MySQL User:", cfg.MySQL.User)
	fmt.Println("MySQL Password:", cfg.MySQL.Password)
	fmt.Println("MySQL Host:", cfg.MySQL.Host)
	fmt.Println("MySQL Port:", cfg.MySQL.Port)
	fmt.Println("MySQL Database:", cfg.MySQL.Database)

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		cfg.MySQL.User,
		cfg.MySQL.Password,
		cfg.MySQL.Host,
		cfg.MySQL.Port,
		cfg.MySQL.Database,
	)

	// Debug log to check DSN
	fmt.Println("DSN:", dsn)

	var log logger.Interface

	if cfg.Env == "dev" {
		log = logger.Default.LogMode(logger.Info)
	} else {
		log = logger.Default.LogMode(logger.Silent)
	}

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: log,
	})

	if err != nil {
		fmt.Println("Failed to connect to database:", err)
		return nil, err
	}

	return db, nil
}
