package db

import (
	"database/sql"
	"embed"
	"log"

	"github.com/pressly/goose/v3"
)

//go:embed migrations/*.sql
var embeddedMigrations embed.FS

func RunMigrations(db *sql.DB) error {
	log.Println("[DB] Iniciando migrations...")

	goose.SetLogger(log.Default())

	if err := goose.SetDialect("sqlite"); err != nil {
		return err
	}

	goose.SetBaseFS(embeddedMigrations)

	if err := goose.Up(db, "migrations"); err != nil {
		return err
	}

	log.Println("[DB] As migrations estão em dia!")
	return nil
}
