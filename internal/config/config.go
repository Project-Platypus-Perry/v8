package config

import (
	"errors"
	"os"
	"strconv"
	"time"

	"github.com/joho/godotenv"
	"github.com/project-platypus-perry/v8/internal/constants"
)

type Config struct {
	Env         string
	Port        string
	DatabaseURL string
	LogLevel    string
	JWT         *JWTConfig
	RateLimiter *RateLimiterConfig
	Email       *EmailConfig
}

type RateLimiterConfig struct {
	Requests int           // Number of requests allowed
	Duration time.Duration // Time window for rate limiting
}

func Load() *Config {
	if err := godotenv.Load(); err != nil {
		panic(err)
	}

	cfg := &Config{
		Env:         getEnv("ENV", string(constants.EnvDevelopment)),
		Port:        getEnv("PORT", "8080"),
		DatabaseURL: getEnv("DATABASE_URL", ""),
		LogLevel:    getEnv("LOG_LEVEL", "info"),
	}

	cfg.InitJWTConfig()
	cfg.InitRateLimiterConfig()
	cfg.InitEmailConfig()

	return cfg
}

func getEnv(key string, defaultVal string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultVal
}

func getEnvAsInt(key string, defaultVal int) int {
	if value, exists := os.LookupEnv(key); exists {
		if intVal, err := strconv.Atoi(value); err == nil {
			return intVal
		}
	}
	return defaultVal
}

func (c *Config) InitRateLimiterConfig() {
	c.RateLimiter = &RateLimiterConfig{
		Requests: getEnvAsInt("RATE_LIMIT_REQUESTS", 100),                            // 100 requests
		Duration: time.Duration(getEnvAsInt("RATE_LIMIT_DURATION", 1)) * time.Minute, // per minute
	}
}

func (c *Config) Validate() error {
	if c.DatabaseURL == "" {
		return errors.New("database URL is required")
	}
	return nil
}
