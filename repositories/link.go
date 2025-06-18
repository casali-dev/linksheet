package repositories

import (
	"database/sql"

	"github.com/casali-dev/linksheet/models"
)

type LinkRepository struct {
	DB *sql.DB
}

func NewLinkRepository(db *sql.DB) *LinkRepository {
	return &LinkRepository{DB: db}
}

func (r *LinkRepository) GetAll() ([]models.Link, error) {
	rows, err := r.DB.Query(`SELECT id, name, description, url FROM links`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var links []models.Link
	for rows.Next() {
		var l models.Link
		if err := rows.Scan(&l.ID, &l.Name, &l.Description, &l.URL); err != nil {
			return nil, err
		}
		links = append(links, l)
	}

	return links, nil
}

func (r *LinkRepository) Insert(link models.Link) error {
	_, err := r.DB.Exec(`
		INSERT INTO links (id, name, description, url)
		VALUES (?, ?, ?, ?)
	`, link.ID, link.Name, link.Description, link.URL)

	return err
}
