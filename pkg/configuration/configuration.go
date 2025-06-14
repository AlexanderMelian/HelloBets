package configuration

import (
	"fmt"
	"log"
	"os"
)

type Config struct {
	DBUser      string
	DBPassword  string
	DBHost      string
	DBPort      string
	DBName      string
	PatternMail string
}

func LoadConfig() (*Config, error) {
	var cfg = &Config{
		DBUser:      os.Getenv("POSTGRES_USER"),
		DBPassword:  os.Getenv("POSTGRES_PASSWORD"),
		DBHost:      os.Getenv("POSTGRES_HOST"),
		DBPort:      os.Getenv("POSTGRES_PORT"),
		DBName:      os.Getenv("POSTGRES_DB"),
		PatternMail: os.Getenv("PATTERN_MAIL"),
	}

	log.Println("Configuration loaded:", cfg)

	if cfg.DBUser == "" || cfg.DBPassword == "" {
		return nil, fmt.Errorf("missing critical DB env vars")
	}
	return cfg, nil
}
