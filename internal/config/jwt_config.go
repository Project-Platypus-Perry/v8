package config

import "time"

type JWTConfig struct {
	AccessTokenSecret  string        `mapstructure:"ACCESS_TOKEN_SECRET"`
	RefreshTokenSecret string        `mapstructure:"REFRESH_TOKEN_SECRET"`
	AccessTokenExpiry  time.Duration `mapstructure:"ACCESS_TOKEN_EXPIRY"`
	RefreshTokenExpiry time.Duration `mapstructure:"REFRESH_TOKEN_EXPIRY"`
}

func (c *Config) InitJWTConfig() {
	c.JWT = &JWTConfig{
		AccessTokenSecret:  getEnv("ACCESS_TOKEN_SECRET", "your-256-bit-secret"),
		RefreshTokenSecret: getEnv("REFRESH_TOKEN_SECRET", "your-256-bit-refresh-secret"),
		AccessTokenExpiry:  time.Duration(getEnvAsInt("ACCESS_TOKEN_EXPIRY", 15)) * time.Minute,
		RefreshTokenExpiry: time.Duration(getEnvAsInt("REFRESH_TOKEN_EXPIRY", 7*24)) * time.Hour,
	}
}
