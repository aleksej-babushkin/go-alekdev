package main

import (
	"context"
	"log"
	"net/http"
	"time"

	"app/internal/config"
	"app/internal/database"
	"app/internal/routes"
)

func main() {
	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("Configuration error: %v", err)
	}

	ctx, cancel := context.WithTimeout(
		context.Background(),
		5*time.Second,
	)
	defer cancel()

	db, err := database.NewPostgresPool(ctx, cfg.DatabaseURL)
	if err != nil {
		log.Fatalf("Database connection failed: %v", err)
	}
	defer db.Close()

	server := &http.Server{
		Addr:              ":" + cfg.Port,
		Handler:           routes.New(db),
		ReadHeaderTimeout: 5 * time.Second,
	}

	log.Printf("Server listening on %s", server.Addr)

	if err := server.ListenAndServe(); err != nil &&
		err != http.ErrServerClosed {
		log.Fatalf("HTTP server failed: %v", err)
	}
}
