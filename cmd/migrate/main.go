package main

import (
	"log"

	"github.com/casali-dev/linkhub/internal/config"
	"github.com/casali-dev/linkhub/internal/db"
	_ "github.com/casali-dev/linkhub/scripts"
)

func main() {
	config.LoadEnv()

	db.Connect()
	defer db.Close()

	if err := db.RunMigrations(db.DB); err != nil {
		log.Fatalf("[MIGRATE] Failed to apply migrations: %v", err)
	}

	log.Println("[MIGRATE] Migrations applied successfully.")
}
