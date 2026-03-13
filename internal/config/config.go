package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	DB     DBConfig
	Server ServerConfig
	App    AppConfig
}

type DBConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	Name     string
	SSLMode  string
}

type ServerConfig struct {
	Port string
	Host string
}

type AppConfig struct {
	Environment string
}

func Load() (*Config, error) {
	_ = godotenv.Load()

	cfg := &Config{
		DB: DBConfig{
			Host:     os.Getenv("DB_HOST"),
			Port:     os.Getenv("DB_PORT"),
			User:     os.Getenv("DB_USER"),
			Password: os.Getenv("DB_PASSWORD"),
			Name:     os.Getenv("DB_NAME"),
			SSLMode:  os.Getenv("DB_SSLMODE"),
		},
		Server: ServerConfig{
			Port: os.Getenv("SERVER_PORT"),
			Host: os.Getenv("SERVER_HOST"),
		},
		App: AppConfig{
			Environment: os.Getenv("APP_ENV"),
		},
	}
	if err := cfg.validate(); err != nil {
		return nil, fmt.Errorf("config validation failed: %w", err)
	}
	return cfg, nil
}

func (c *Config) validate() error {
	if c.DB.Host == "" {
		return fmt.Errorf("DB_HOST is required")
	}
	if c.DB.User == "" {
		return fmt.Errorf("DB_USER is required")
	}
	if c.DB.Name == "" {
		return fmt.Errorf("DB_NAME is required")
	}
	return nil
}
