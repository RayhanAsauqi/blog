package config

import (
	"github.com/caarlos0/env/v6"
	"github.com/joho/godotenv"
)

type Config struct {
	Env          string      `env:"ENV" envDefault:"dev"`
	Port         string      `env:"PORT" envDefault:"8080"`
	MySQL        MysqlConfig `envPrefix:"MYSQL_"`
	JWTSecret    string      `env:"JWT_SECRET_KEY" envDefault:"secret"`
}

type MysqlConfig struct {
	Host     string `env:"HOST" envDefault:"localhost"`
	Port     string `env:"PORT" envDefault:"3306"`
	User     string `env:"USER" envDefault:"root"`
	Password string `env:"PASSWORD" envDefault:""`
	Database string `env:"DATABASE" envDefault:"blog"`
}

func NewConfig(envPath string) (*Config, error) {
	// Load environment variables from file
	if err := godotenv.Load(envPath); err != nil {
		return nil, err
	}

	cfg := new(Config)
	// Parse environment variables into Config struct
	if err := env.Parse(cfg); err != nil {
		return nil, err
	}

	return cfg, nil
}