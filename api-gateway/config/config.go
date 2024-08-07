package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

type (
	Config struct {
		AuthHost      string
		EventHost     string
		MedalHost     string
		AthleteHost   string
		AiService     string
		ServerAddress string
		StreamHost    string
	}
)
const (
	OtpSecret = "some_secret"
	SignKey   = "nodirbek"
)
func (c *Config) Load() error {
	if err := godotenv.Load(); err != nil {
		fmt.Println("AAAAAAAAAAAAAAAAAAAAAA", err)
		return err
	}
	c.AuthHost = os.Getenv("AUTH_HOST")
	c.EventHost = os.Getenv("EVENT_HOST")
	c.MedalHost = os.Getenv("MEDAL_HOST")
	c.AthleteHost = os.Getenv("ATHLETE_HOST")
	c.AiService = os.Getenv("AI_SERVICE")
	c.StreamHost = os.Getenv("STREAM_HOST")
	c.ServerAddress = os.Getenv("SERVER_ADDRESS")
	return nil
}

func New() (*Config, error) {
	var cnfg Config
	if err := cnfg.Load(); err != nil {
		return nil, err
	}
	return &cnfg, nil
}
