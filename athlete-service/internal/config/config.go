package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"

	_ "github.com/lib/pq"
)

type Config struct {
	Server   ServerConfig
	Database DatabaseConfig
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

func (c *Config) Load() error {

	err := godotenv.Load()
	if err != nil {
		fmt.Println(".env file not found, proceeding with other environment variables")
	}

	c.Server.Port = ":" + getEnv("SERVER_PORT", "6666")
	c.Database.Host = getEnv("DB_HOST", "postgres")
	c.Database.Port = getEnv("DB_PORT", "5432")
	c.Database.User = getEnv("DB_USER", "postgres")
	c.Database.Password = getEnv("DB_PASSWORD", "aaaa")
	c.Database.DBName = getEnv("DB_NAME", "olympydb")

	return nil
}

func getEnv(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}

func New() (*Config, error) {
	config := &Config{}
	err := config.Load()
	if err != nil {
		return nil, err
	}
	return config, nil
}
