package main

import (
	"fmt"
	"log"

	"github.com/p-doukas/go-session-auth/internal/config"
	"github.com/p-doukas/go-session-auth/internal/models"
	"github.com/p-doukas/go-session-auth/migrations"
)

func migrate() {
	cfg, err := config.LoadEnvConfig()
	if err != nil {
		log.Fatal("failed to load config: ", err)
	}

	db, err := models.Open(cfg.PSQL)
	if err != nil {
		log.Fatal("failed to open database: ", err)
	}
	defer db.Close()

	err = models.MigrateFS(db, migrations.FS, ".")
	if err != nil {
		log.Fatal("failed to run migrations: ", err)
	}

	fmt.Println("Migrations complete.")
}
