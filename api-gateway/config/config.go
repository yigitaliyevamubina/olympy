package config

import (
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
	}
)

func (c *Config) Load() error {
	if err := godotenv.Load(); err != nil {
		return err
	}
	c.AuthHost = os.Getenv("AUTH_HOST")
	c.EventHost = os.Getenv("EVENT_HOST")
	c.MedalHost = os.Getenv("MEDAL_HOST")
	c.AthleteHost = os.Getenv("ATHLETE_HOST")
	c.AiService = os.Getenv("AI_SERVICE")
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
