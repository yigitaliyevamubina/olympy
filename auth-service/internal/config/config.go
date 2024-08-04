package config

import (
	"fmt"
	"os"
	"time"

	"github.com/joho/godotenv"
)

type Config struct {
	Server   ServerConfig
	Database DatabaseConfig
	JWT      JWTConfig
}

type ServerConfig struct {
	Port string
}

type DatabaseConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	DBName   string
}

type JWTConfig struct {
	Secret          string
	AccessTokenExp  time.Duration // Duration for access token
	RefreshTokenExp time.Duration // Duration for refresh token
}

func (c *Config) Load() error {
	// Load environment variables from a .env file if it exists
	if err := godotenv.Load(); err != nil {
		return fmt.Errorf("error loading .env file: %w", err)
	}

	// Load server configuration
	serverPort := os.Getenv("SERVER_PORT")
	if serverPort == "" {
		return fmt.Errorf("SERVER_PORT environment variable is required")
	}
	c.Server.Port = ":" + serverPort

	// Load database configuration
	c.Database.Host = os.Getenv("DB_HOST")
	c.Database.Port = os.Getenv("DB_PORT")
	c.Database.User = os.Getenv("DB_USER")
	c.Database.Password = os.Getenv("DB_PASSWORD")
	c.Database.DBName = os.Getenv("DB_NAME")

	// Check for required database configuration
	if c.Database.Host == "" || c.Database.Port == "" || c.Database.User == "" || c.Database.Password == "" || c.Database.DBName == "" {
		return fmt.Errorf("one or more database configuration variables are missing")
	}

	// Load JWT configuration
	c.JWT.Secret = os.Getenv("JWT_SECRET")
	if c.JWT.Secret == "" {
		return fmt.Errorf("JWT_SECRET environment variable is required")
	}

	// Parse token expiration times
	accessTokenExp, err := time.ParseDuration(os.Getenv("ACCESS_TOKEN_EXP"))
	if err != nil {
		return fmt.Errorf("invalid ACCESS_TOKEN_EXP value: %w", err)
	}
	c.JWT.AccessTokenExp = accessTokenExp

	refreshTokenExp, err := time.ParseDuration(os.Getenv("REFRESH_TOKEN_EXP"))
	if err != nil {
		return fmt.Errorf("invalid REFRESH_TOKEN_EXP value: %w", err)
	}
	c.JWT.RefreshTokenExp = refreshTokenExp

	return nil
}

func New() (*Config, error) {
	config := &Config{}
	if err := config.Load(); err != nil {
		return nil, err
	}
	return config, nil
}
