package config

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
	"github.com/p-doukas/go-session-auth/internal/models"
)

type config struct {
	PSQL    models.PostgresConfig
	SMTP    models.SMTPConfig
	BaseURL string
	CSRF    struct {
		Key    string
		Secure bool
	}
	Server struct {
		Address string
	}
}

func LoadEnvConfig() (config, error) {
	var cfg config

	err := godotenv.Load()
	if err != nil {
		log.Println("No .env file found, using environment variables")
	}

	cfg.PSQL = models.DefaultPostgresConfig()

	// See: docs/notes.md#go-assignment-rules
	cfg.SMTP, err = models.DefaultSMTPConfig()
	if err != nil {
		return cfg, err
	}

	cfg.BaseURL = os.Getenv("BASE_URL")
	if cfg.BaseURL == "" {
		cfg.BaseURL = "http://localhost:8080"
	}

	cfg.CSRF.Key = os.Getenv("CSRF_KEY")
	cfg.CSRF.Secure, err = strconv.ParseBool(os.Getenv("CSRF_SECURE"))
	if err != nil {
		return cfg, fmt.Errorf("invalid CSRF_SECURE value: %w", err)
	}

	cfg.Server.Address = os.Getenv("SERVER_ADDRESS")
	if cfg.Server.Address == "" {
		cfg.Server.Address = ":8080"
	}

	return cfg, nil
}
