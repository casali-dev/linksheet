package repositories

import (
	"database/sql"

	"github.com/casali-dev/linksheet/models"
)

type LinkRepository interface {
	GetAll() ([]models.Link, error)
	Insert(models.Link) error
}

type SQLiteLinkRepository struct {
	DB *sql.DB
}

func NewLinkRepository(db *sql.DB) LinkRepository {
	return &SQLiteLinkRepository{DB: db}
}

func (r *SQLiteLinkRepository) GetAll() ([]models.Link, error) {
	rows, err := r.DB.Query(`
		SELECT id, name, description, url, created_at, updated_at FROM links
	`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var links []models.Link
	for rows.Next() {
		var l models.Link
		if err := rows.Scan(&l.ID, &l.Name, &l.Description, &l.URL, &l.CreatedAt, &l.UpdatedAt); err != nil {
			return nil, err
		}
		links = append(links, l)
	}

	if links == nil {
		links = []models.Link{}
	}

	return links, nil
}

func (r *SQLiteLinkRepository) Insert(link models.Link) error {
	_, err := r.DB.Exec(`
	INSERT INTO links (id, name, description, url, created_at, updated_at)
	VALUES (?, ?, ?, ?, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP)
`, link.ID, link.Name, link.Description, link.URL)

	return err
}
