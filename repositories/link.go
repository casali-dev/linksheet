package repositories

import (
	"database/sql"

	"github.com/casali-dev/linksheet/models"
)

type LinkRepository interface {
	GetAll() ([]models.Link, error)
	GetByAuthor(authorID string) ([]models.Link, error)
	Insert(link models.Link) error
}

type SQLiteLinkRepository struct {
	DB *sql.DB
}

func NewLinkRepository(db *sql.DB) LinkRepository {
	return &SQLiteLinkRepository{DB: db}
}

func (r *SQLiteLinkRepository) GetAll() ([]models.Link, error) {
	rows, err := r.DB.Query(`
		SELECT id, name, description, url, is_public, author_id, created_at, updated_at
		FROM links
		WHERE is_public = 1
	`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var links []models.Link
	for rows.Next() {
		var l models.Link
		if err := rows.Scan(
			&l.ID,
			&l.Name,
			&l.Description,
			&l.URL,
			&l.IsPublic,
			&l.AuthorID,
			&l.CreatedAt,
			&l.UpdatedAt,
		); err != nil {
			return nil, err
		}
		links = append(links, l)
	}

	if links == nil {
		links = []models.Link{}
	}

	return links, nil
}

func (r *SQLiteLinkRepository) GetByAuthor(authorID string) ([]models.Link, error) {
	rows, err := r.DB.Query(`
		SELECT id, name, description, url, is_public, author_id, created_at, updated_at
		FROM links
		WHERE author_id = ?
	`, authorID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var links []models.Link
	for rows.Next() {
		var link models.Link
		err := rows.Scan(
			&link.ID,
			&link.Name,
			&link.Description,
			&link.URL,
			&link.IsPublic,
			&link.AuthorID,
			&link.CreatedAt,
			&link.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		links = append(links, link)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return links, nil
}

func (r *SQLiteLinkRepository) Insert(link models.Link) error {
	_, err := r.DB.Exec(`
	INSERT INTO links (id, name, description, url, is_public, author_id, created_at, updated_at)
	VALUES (?, ?, ?, ?, ?, ?, ?, ?)
`, link.ID, link.Name, link.Description, link.URL, link.IsPublic, link.AuthorID, link.CreatedAt, link.UpdatedAt)

	return err
}
