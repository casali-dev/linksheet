package models

import (
	"time"

	"github.com/google/uuid"
)

type Link struct {
	ID          string    `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	URL         string    `json:"url"`
	CreatedAt   time.Time `json:"createdAt" db:"created_at"`
	UpdatedAt   time.Time `json:"updatedAt" db:"updated_at"`
}

func NewLink(name, description, url string) Link {
	now := time.Now()

	return Link{
		ID:          uuid.New().String(),
		Name:        name,
		Description: description,
		URL:         url,
		CreatedAt:   now,
		UpdatedAt:   now,
	}
}
