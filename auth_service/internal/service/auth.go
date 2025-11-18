package service

import (
	"articles/internal/models"
	"articles/internal/repository"
	"fmt"
)

type AuthDI struct {
	repo repository.AuthRepository
}

func NewAuthService(repo repository.AuthRepository) *AuthDI {
	return &AuthDI{
		repo: repo,
	}
}

func (r *AuthDI) Register() (models.User, error) {
	user, err := r.repo.RegisterUser()
	if err != nil {
		fmt.Println(err)
	}

	return user, nil
}
