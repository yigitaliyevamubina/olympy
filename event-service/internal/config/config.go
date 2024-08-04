package config

import (
	"os"
	"strconv"
	"time"

	"github.com/joho/godotenv"
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
	err := godotenv.Load()
	if err != nil {
		return err
	}

	c.Server.Port = ":" + os.Getenv("SERVER_PORT")
	c.Database.Host = os.Getenv("DB_HOST")
	c.Database.Port = os.Getenv("DB_PORT")
	c.Database.User = os.Getenv("DB_USER")
	c.Database.Password = os.Getenv("DB_PASSWORD")
	c.Database.DBName = os.Getenv("DB_NAME")

	c.Redis.Addr = os.Getenv("REDIS_ADDR")
	c.Redis.Password = os.Getenv("REDIS_PASSWORD")
	c.Redis.DB = 0

	ac := os.Getenv("JWT_ACCESS_TOKEN_EXP")
	intacc, _ := strconv.Atoi(ac)
	re := os.Getenv("JWT_REFRESH_TOKEN_EXP")
	intref, _ := strconv.Atoi(re)
	c.JWT.Secret = os.Getenv("JWT_SECRET")
	c.JWT.AccessTokenExp = time.Duration(intacc) * time.Minute
	c.JWT.RefreshTokenExp = time.Duration(intref) * time.Minute

	return nil
}

func New() (*Config, error) {
	config := &Config{}
	err := config.Load()
	if err != nil {
		return nil, err
	}
	return config, nil
}
