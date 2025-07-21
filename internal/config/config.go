package config

import (
	"computer-manager/internal/helpers"
	"time"
)

type Config struct {
	ProjectRoot string
	Server      ServerConfig
	Database    DatabaseConfig
	Cache       CacheConfig
}

type ServerConfig struct {
	Port int
}

type DatabaseConfig struct {
	Host     string
	Port     int
	User     string
	Password string
	Name     string
	SSLMode  string
	MaxConns int
	Timeout  time.Duration
}

type CacheConfig struct {
	Host          string
	Port          string
	DB            int
	Password      string
	DefaultTTL    time.Duration
	CleanupPeriod time.Duration
}

func NewConfig() (*Config, error) {
	return &Config{
		ProjectRoot: helpers.GetEnv("PROJECT_ROOT", "/app"),
		Server: ServerConfig{
			Port: helpers.GetEnv("PORT", 8000),
		},
		Database: DatabaseConfig{
			Host:     helpers.GetEnv("DB_HOST", "postgres"),
			Port:     helpers.GetEnv("DB_PORT", 5432),
			User:     helpers.GetEnv("DB_USER", "postgres"),
			Password: helpers.GetEnv("DB_PASSWORD", "postgres"),
			Name:     helpers.GetEnv("DB_NAME", "postgres"),
			SSLMode:  helpers.GetEnv("DB_SSLMODE", "disable"),
			MaxConns: helpers.GetEnv("DB_MAX_CONNS", 10),
			Timeout:  helpers.GetEnv("DB_TIMEOUT", 5*time.Second),
		},
		Cache: CacheConfig{
			Host:     helpers.GetEnv("REDIS_HOST", "redis"),
			Port:     helpers.GetEnv("REDIS_PORT", "6379"),
			DB:       helpers.GetEnv("REDIS_DB", 0),
			Password: helpers.GetEnv("REDIS_PASSWORD", ""),
		},
	}, nil
}
