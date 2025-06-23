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
	IsPublic    bool      `json:"isPublic"`
	AuthorID    string    `json:"authorId"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
}

type PublicLink struct {
	ID          string    `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	URL         string    `json:"url"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
}

func NewLink(name, description, url string, isPublic bool, authorID string) Link {
	now := time.Now()

	return Link{
		ID:          uuid.New().String(),
		Name:        name,
		Description: description,
		URL:         url,
		IsPublic:    isPublic,
		AuthorID:    authorID,
		CreatedAt:   now,
		UpdatedAt:   now,
	}
}
