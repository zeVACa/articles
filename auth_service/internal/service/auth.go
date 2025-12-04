package service

import (
	"articles/internal/models"
	"articles/internal/repository"
)

type AuthDI struct {
	repo repository.AuthRepository
}

func NewAuthService(repo repository.AuthRepository) *AuthDI {
	return &AuthDI{
		repo: repo,
	}
}

type Service interface {
	Register(email, username, passwordHash string) (*models.User, error)
}

func (r *AuthDI) Register(email, username, passwordHash string) (*models.User, error) {
	user, err := r.repo.RegisterUser(email, username, passwordHash)
	if err != nil {
		return nil, err
	}

	return user, nil
}
