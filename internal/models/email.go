package models

import (
	"fmt"
	"os"
	"strconv"

	"github.com/go-mail/mail/v2"
)

const (
	DefaultSender = "doukas.py@gmail.com"
)

type SMTPConfig struct {
	Host     string
	Port     int
	Username string
	Password string
}

type EmailService struct {
	DefaultSender string
	BaseURL       string
	// Dialer allows connecting to the email server
	// Dialer allows connecting to the email server
	// (e.g Mailtrap) and send emails.
	dialer *mail.Dialer
	// TODO: Add templates
}

func DefaultSMTPConfig() (SMTPConfig, error) {
	port, err := strconv.Atoi(os.Getenv("SMTP_PORT"))
	if err != nil {
		return SMTPConfig{}, fmt.Errorf("Invalid SMTP_PORT: %w", err)
	}

	return SMTPConfig{
		Host:     os.Getenv("SMTP_HOST"),
		Port:     port,
		Username: os.Getenv("SMTP_USERNAME"),
		Password: os.Getenv("SMTP_PASSWORD"),
	}, nil
}

func NewEmailService(config SMTPConfig) *EmailService {
	return &EmailService{
		dialer: mail.NewDialer(config.Host, config.Port, config.Username, config.Password),
	}
}
