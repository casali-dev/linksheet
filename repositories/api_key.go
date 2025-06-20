package repositories

import (
	"database/sql"
)

type APIKeyRepository interface {
	GetAuthorIDByKey(key string) (string, error)
}

type SQLiteAPIKeyRepository struct {
	DB *sql.DB
}

func NewAPIKeyRepository(db *sql.DB) APIKeyRepository {
	return &SQLiteAPIKeyRepository{DB: db}
}

func (r *SQLiteAPIKeyRepository) GetAuthorIDByKey(key string) (string, error) {
	var authorID string
	err := r.DB.QueryRow("SELECT author_id FROM api_keys WHERE key = ?", key).Scan(&authorID)
	if err != nil {
		if err == sql.ErrNoRows {
			return "", nil // não encontrada
		}
		return "", err
	}
	return authorID, nil
}
