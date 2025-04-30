package emailService

import (
	"fmt"
	"net/smtp"

	"github.com/project-platypus-perry/v8/internal/config"
)

type EmailService interface {
	SendInviteEmail(to string, name string, email string, password string) error
	SendPasswordResetEmail(to string, resetToken string) error
}

type emailService struct {
	smtpHost     string
	smtpPort     int
	smtpUsername string
	smtpPassword string
	fromEmail    string
	appBaseURL   string
}

func NewEmailService(cfg *config.EmailConfig) EmailService {
	return &emailService{
		smtpHost:     cfg.SMTPHost,
		smtpPort:     cfg.SMTPPort,
		smtpUsername: cfg.SMTPUser,
		smtpPassword: cfg.SMTPPassword,
		fromEmail:    cfg.FromEmail,
		appBaseURL:   cfg.AppBaseURL,
	}
}

func (s *emailService) SendInviteEmail(to string, name string, email string, password string) error {
	subject := "Welcome to Platform - Your Account Details"
	body := fmt.Sprintf(`
		Dear %s,

		Welcome to our platform! Your account has been created with the following credentials:

		Email: %s
		Password: %s

		Please login at %s/login and change your password immediately.

		Best regards,
		Your Platform Team
	`, name, email, password, s.appBaseURL)

	return s.sendEmail(to, subject, body)
}

func (s *emailService) SendPasswordResetEmail(to string, resetToken string) error {
	subject := "Password Reset Request"
	resetLink := fmt.Sprintf("%s/reset-password?token=%s", s.appBaseURL, resetToken)
	body := fmt.Sprintf(`
		Hello,

		You have requested to reset your password. Please click the link below to reset your password:

		%s

		If you did not request this, please ignore this email.

		Best regards,
		Your Platform Team
	`, resetLink)

	return s.sendEmail(to, subject, body)
}

func (s *emailService) sendEmail(to string, subject string, body string) error {
	auth := smtp.PlainAuth("", s.smtpUsername, s.smtpPassword, s.smtpHost)

	msg := fmt.Sprintf("From: %s\r\n"+
		"To: %s\r\n"+
		"Subject: %s\r\n"+
		"MIME-Version: 1.0\r\n"+
		"Content-Type: text/plain; charset=utf-8\r\n"+
		"\r\n"+
		"%s", s.fromEmail, to, subject, body)

	return smtp.SendMail(
		fmt.Sprintf("%s:%d", s.smtpHost, s.smtpPort),
		auth,
		s.fromEmail,
		[]string{to},
		[]byte(msg),
	)
}
