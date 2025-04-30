package config

type EmailConfig struct {
	SMTPHost     string
	SMTPPort     int
	SMTPUser     string
	SMTPPassword string
	FromEmail    string
	AppBaseURL   string
}

func (c *Config) InitEmailConfig() {
	c.Email = &EmailConfig{
		SMTPHost:     getEnv("SMTP_HOST", "smtp.gmail.com"),
		SMTPPort:     getEnvAsInt("SMTP_PORT", 587),
		SMTPUser:     getEnv("SMTP_USER", ""),
		SMTPPassword: getEnv("SMTP_PASSWORD", ""),
		FromEmail:    getEnv("FROM_EMAIL", ""),
		AppBaseURL:   getEnv("APP_BASE_URL", ""),
	}
}
