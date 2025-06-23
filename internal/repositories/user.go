package repositories

import (
	"database/sql"
	"errors"

	"github.com/casali-dev/linkhub/internal/models"
)

type UserRepository interface {
	FindByEmail(email string) (*models.User, error)
	Insert(u models.User) error
}

type SQLiteUserRepository struct {
	DB *sql.DB
}

func NewUserRepository(db *sql.DB) UserRepository {
	return &SQLiteUserRepository{DB: db}
}

func (r *SQLiteUserRepository) FindByEmail(email string) (*models.User, error) {
	row := r.DB.QueryRow(`
		SELECT id, email, password_hash, role, created_at, updated_at
		FROM users
		WHERE email = ?
	`, email)

	var u models.User
	if err := row.Scan(
		&u.ID,
		&u.Email,
		&u.PasswordHash,
		&u.Role,
		&u.CreatedAt,
		&u.UpdatedAt,
	); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}
	return &u, nil
}

func (r *SQLiteUserRepository) Insert(u models.User) error {
	_, err := r.DB.Exec(`
		INSERT INTO users (id, email, password_hash, role, created_at, updated_at)
		VALUES (?, ?, ?, ?, ?, ?)
	`, u.ID, u.Email, u.PasswordHash, u.Role, u.CreatedAt, u.UpdatedAt)
	return err
}
