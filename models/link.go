package models

import (
	"github.com/google/uuid"
)

type Link struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	URL         string `json:"url"`
}

func NewLink(name, description, url string) Link {
	return Link{
		ID:          uuid.New().String(),
		Name:        name,
		Description: description,
		URL:         url,
	}
}