package config

import (
	"errors"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/project-platypus-perry/v8/internal/constants"
)

type Config struct {
	Env         string
	DatabaseURL string
	LogLevel    string
	Port        string
}

func Load() *Config {
	// Load from .env
	err := godotenv.Load(".env")
	if err != nil {
		log.Println("No .env file found, using system environment variables")
	}
	env := os.Getenv("ENV")
	if env == "" {
		env = string(constants.EnvDevelopment)
	}
	dbUrl := os.Getenv("DATABASE_URL")
	if dbUrl == "" {
		log.Fatal("DATABASE_URL is not set")
	}
	logLevel := os.Getenv("LOG_LEVEL") // if not set, default to "info"
	if logLevel == "" {
		logLevel = "info"
	}
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	return &Config{
		Env:         env,
		DatabaseURL: dbUrl,
		LogLevel:    logLevel,
		Port:        port,
	}
}

func (c *Config) Validate() error {
	if c.DatabaseURL == "" {
		return errors.New("database URL is required")
	}
	return nil
}
