package services

import (
	"errors"

	"github.com/casali-dev/linksheet/models"
	"github.com/casali-dev/linksheet/repositories"
)

type LinkService interface {
	GetAll() ([]models.Link, error)
	Create(name, description, url string) (models.Link, error)
}

type DefaultLinkService struct {
	Repo repositories.LinkRepository
}

func NewLinkService(repo repositories.LinkRepository) LinkService {
	return &DefaultLinkService{Repo: repo}
}

func (s *DefaultLinkService) GetAll() ([]models.Link, error) {
	return s.Repo.GetAll()
}

func (s *DefaultLinkService) Create(name, description, url string) (models.Link, error) {
	if name == "" || url == "" {
		return models.Link{}, errors.New("os campos 'name' e 'url' são obrigatórios")
	}

	link := models.NewLink(name, description, url)

	if err := s.Repo.Insert(link); err != nil {
		return models.Link{}, err
	}

	return link, nil
}
