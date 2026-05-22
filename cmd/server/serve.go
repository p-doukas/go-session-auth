package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/p-doukas/go-session-auth/internal/config"
	"github.com/p-doukas/go-session-auth/internal/models"
)

func serve() {
	cfg, err := config.LoadEnvConfig()
	if err != nil {
		log.Fatal("failed to load config: ", err)
	}

	db, err := models.Open(cfg.PSQL)
	if err != nil {
		log.Fatal("failed to open database: ", err)
	}
	defer db.Close()

	// Set up router and routes
	r := chi.NewRouter()

	r.Get("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"status":"ok"}`))
	})

	fmt.Printf("Starting the server on %s...\n", cfg.Server.Address)
	err = http.ListenAndServe(cfg.Server.Address, r)
	if err != nil {
		log.Fatal("server failed to start: ", err)
	}
}
