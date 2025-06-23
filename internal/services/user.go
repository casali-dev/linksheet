package services

import (
	"errors"
	"regexp"

	"github.com/casali-dev/linksheet/internal/models"
	"github.com/casali-dev/linksheet/internal/repositories"
	"golang.org/x/crypto/bcrypt"
)

var emailRegex = regexp.MustCompile(`^[^@\s]+@[^@\s]+\.[^@\s]+$`)

type UserService interface {
	Signup(email, password string) (models.User, error)
}

type DefaultUserService struct {
	Repo repositories.UserRepository
}

func NewUserService(repo repositories.UserRepository) UserService {
	return &DefaultUserService{Repo: repo}
}

func (s *DefaultUserService) Signup(email, password string) (models.User, error) {
	if !emailRegex.MatchString(email) {
		return models.User{}, errors.New("invalid email format")
	}
	if len(password) < 8 {
		return models.User{}, errors.New("password must be at least 8 characters")
	}

	existing, err := s.Repo.FindByEmail(email)
	if err != nil {
		return models.User{}, err
	}
	if existing != nil {
		return models.User{}, errors.New("email already registered")
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return models.User{}, err
	}

	user := models.NewUser(email, string(hash))
	if err := s.Repo.Insert(user); err != nil {
		return models.User{}, err
	}
	return user, nil
}
