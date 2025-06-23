package db

import (
	"context"
	"database/sql"
	"errors"
	"time"

	"github.com/casali-dev/linksheet/config"
	"github.com/google/uuid"
	"github.com/pressly/goose/v3"
	"golang.org/x/crypto/bcrypt"
)

func init() {
	goose.AddMigrationContext(upSeedAdminUser, downSeedAdminUser)
}

func upSeedAdminUser(ctx context.Context, tx *sql.Tx) error {
	email := config.MustGet("LINKHUB_ADMIN_EMAIL")
	password := config.MustGet("LINKHUB_ADMIN_PASSWORD")

	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	now := time.Now().UTC()
	_, err = tx.ExecContext(ctx, `
		INSERT INTO users (id, email, password_hash, role, created_at, updated_at)
		VALUES (?, ?, ?, 'admin', ?, ?)
	`, uuid.NewString(), email, string(hash), now, now)

	return err
}

func downSeedAdminUser(ctx context.Context, tx *sql.Tx) error {
	email := config.MustGet("LINKHUB_ADMIN_EMAIL")
	if email == "" {
		return errors.New("LINKHUB_ADMIN_EMAIL must be set to roll back")
	}

	_, err := tx.ExecContext(ctx, `DELETE FROM users WHERE email = ?`, email)
	return err
}
