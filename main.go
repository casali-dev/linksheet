package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/casali-dev/linksheet/config"
	"github.com/casali-dev/linksheet/db"
	"github.com/casali-dev/linksheet/router"
)

func main() {
	config.LoadEnv()

	db.Connect()
	defer db.Close()

	if err := db.RunMigrations(db.DB); err != nil {
		log.Fatalf("[DB] Failed to run migrations: %v", err)
	}

	port := config.Get("LINKHUB_PORT", "3333")

	server := &http.Server{
		Addr:    ":" + port,
		Handler: router.Handler(),
	}

	serverErrors := make(chan error, 1)

	go func() {
		log.Printf("[Server] Listening on port %s", port)
		serverErrors <- server.ListenAndServe()
	}()

	sigint := make(chan os.Signal, 1)
	signal.Notify(sigint, syscall.SIGINT, syscall.SIGTERM)

	select {
	case err := <-serverErrors:
		log.Fatalf("[Server] Fatal error: %v", err)

	case sig := <-sigint:
		log.Printf("[Server] Signal caught: %v. Shutting down gracefully...", sig)

		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		if err := server.Shutdown(ctx); err != nil {
			log.Printf("[Server] Shutdown error: %v", err)
		} else {
			log.Println("[Server] Shutdown completed successfully")
		}
	}
}
