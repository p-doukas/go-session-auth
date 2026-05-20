package models

import (
	"database/sql"
	"fmt"
	"io/fs"
	"os"

	_ "github.com/jackc/pgx/v4/stdlib"
	"github.com/pressly/goose/v3"
)

type PostgresConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	Database string
	SSLMode  string
}

func (cfg PostgresConfig) String() string {
	// If Host starts with /, it's a Unix socket (Cloud Run / Cloud SQL)
	if len(cfg.Host) > 0 && cfg.Host[0] == '/' {
		return fmt.Sprintf("host=%s user=%s password=%s dbname=%s sslmode=%s",
			cfg.Host, cfg.User, cfg.Password, cfg.Database, cfg.SSLMode)
	}

	// Otherwise normal TCP connection
	return fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		cfg.Host, cfg.Port, cfg.User, cfg.Password, cfg.Database, cfg.SSLMode)
}

// Open will open a SQL connection with the provided Postgres database.
// Callers of Open need to ensure that the connection is eventually
// closed via the db.Close() method.
func Open(config PostgresConfig) (*sql.DB, error) {
	db, err := sql.Open("pgx", config.String())
	if err != nil {
		return nil, fmt.Errorf("open: %w", err)
	}
	return db, nil
}

func DefaultPostgresConfig() PostgresConfig {
	return PostgresConfig{
		Host:     os.Getenv("PG_HOST"),
		Port:     os.Getenv("PG_PORT"),
		User:     os.Getenv("PG_USER"),
		Password: os.Getenv("PG_PASSWORD"),
		Database: os.Getenv("PG_DATABASE"),
		SSLMode:  os.Getenv("PG_SSLMODE"),
	}
}

func Migrate(db *sql.DB, dir string) error {
	err := goose.SetDialect("postgres")
	if err != nil {
		return fmt.Errorf("migrate: %w", err)
	}

	err = goose.Up(db, dir)
	if err != nil {
		return fmt.Errorf("migrate: %w", err)
	}

	return nil
}

func MigrateFS(db *sql.DB, migrationFS fs.FS, dir string) error {
	if dir == "" {
		dir = "."
	}

	// Configure goose to read migrations from the provided filesystem.
	// This allows migrations to be loaded from embedded or non-disk sources.
	goose.SetBaseFS(migrationFS)

	// Reset the base filesystem after running migrations to avoid
	// affecting other goose operations.
	defer func() {
		goose.SetBaseFS(nil)
	}()

	return Migrate(db, dir)
}
