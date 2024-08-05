package config

import (
	"log"
	"os"
	"strconv"
	"time"

	"github.com/joho/godotenv"

	_ "github.com/lib/pq"
)

type Config struct {
	Server   ServerConfig
	Database DatabaseConfig
	Redis    RedisConfig
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

type RedisConfig struct {
	Addr     string
	Password string
	DB       int
}

type JWTConfig struct {
	Secret          string
	AccessTokenExp  time.Duration
	RefreshTokenExp time.Duration
}

func (c *Config) Load() error {
	if err := godotenv.Load(); err != nil {
		// Log a warning but don't fail if the .env file is not found
		log.Println("Warning: .env file not found or failed to load.")
	}

	c.Server.Port = ":" + getEnv("SERVER_PORT", "4444")
	c.Database.Host = getEnv("DB_HOST", "localhost")
	c.Database.Port = getEnv("DB_PORT", "5544")
	c.Database.User = getEnv("DB_USER", "postgres")
	c.Database.Password = getEnv("DB_PASSWORD", "1234")
	c.Database.DBName = getEnv("DB_NAME", "olympydb")

	c.Redis.Addr = getEnv("REDIS_ADDR", "localhost:6379")
	c.Redis.Password = getEnv("REDIS_PASSWORD", "")
	c.Redis.DB = 0

	ac := getEnv("JWT_ACCESS_TOKEN_EXP", "15")
	intacc, _ := strconv.Atoi(ac)
	re := getEnv("JWT_REFRESH_TOKEN_EXP", "10080") // Default to 7 days in minutes
	intref, _ := strconv.Atoi(re)
	c.JWT.Secret = getEnv("JWT_SECRET", "your_jwt_secret")
	c.JWT.AccessTokenExp = time.Duration(intacc) * time.Minute
	c.JWT.RefreshTokenExp = time.Duration(intref) * time.Minute

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
