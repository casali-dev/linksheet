package services

import (
	"errors"
	"regexp"
	"strings"

	"github.com/casali-dev/linksheet/models"
	"github.com/casali-dev/linksheet/repositories"
)

type LinkService interface {
	GetPublic() ([]models.PublicLink, error)
	GetAllByAuthor(authorID string) ([]models.Link, error)
	Create(name, description, url string, isPublic bool, authorID string) (models.Link, error)
}

type DefaultLinkService struct {
	Repo repositories.LinkRepository
}

func NewLinkService(repo repositories.LinkRepository) LinkService {
	return &DefaultLinkService{Repo: repo}
}

func (s *DefaultLinkService) GetPublic() ([]models.PublicLink, error) {
	all, err := s.Repo.GetAll()
	if err != nil {
		return nil, err
	}

	public := make([]models.PublicLink, 0, len(all))
	for _, link := range all {
		public = append(public, models.PublicLink{
			ID:          link.ID,
			Name:        link.Name,
			Description: link.Description,
			URL:         link.URL,
			CreatedAt:   link.CreatedAt,
			UpdatedAt:   link.UpdatedAt,
		})
	}

	return public, nil
}

func (s *DefaultLinkService) GetAllByAuthor(authorID string) ([]models.Link, error) {
	return s.Repo.GetByAuthor(authorID)
}

var urlRegex = regexp.MustCompile(`^https?://[^\s]+$`)

func (s *DefaultLinkService) Create(name, description, url string, isPublic bool, authorID string) (models.Link, error) {
	name = strings.TrimSpace(name)
	url = strings.TrimSpace(url)

	switch {
	case name == "" || len(name) > 100:
		return models.Link{}, errors.New("name must be 1-100 characters")
	case url == "" || len(url) > 500 || !urlRegex.MatchString(url):
		return models.Link{}, errors.New("invalid url")
	case len(description) > 300:
		return models.Link{}, errors.New("description max 300 characters")
	case authorID == "":
		return models.Link{}, errors.New("authorID is required")
	}

	link := models.NewLink(name, description, url, isPublic, authorID)

	if err := s.Repo.Insert(link); err != nil {
		return models.Link{}, err
	}

	return link, nil
}
