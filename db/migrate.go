package db

import (
	"embed"
	"fmt"
	"log"
)

//go:embed migrations/*.sql
var migrationFiles embed.FS

func RunMigrations() error {
	entries, err := migrationFiles.ReadDir("migrations")
	if err != nil {
		return fmt.Errorf("erro lendo arquivos de migration: %w", err)
	}

	for _, entry := range entries {
		content, err := migrationFiles.ReadFile("migrations/" + entry.Name())
		if err != nil {
			return fmt.Errorf("erro lendo conteúdo da migration %s: %w", entry.Name(), err)
		}

		if _, err := DB.Exec(string(content)); err != nil {
			return fmt.Errorf("erro executando migration %s: %w", entry.Name(), err)
		}

		log.Printf("[DB] Migration %s executada com sucesso", entry.Name())
	}

	return nil
}
