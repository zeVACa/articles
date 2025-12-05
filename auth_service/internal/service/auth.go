package service

import (
	"articles/internal/repository"
	"articles/pgk/validation"
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"net/http"
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
	Register(email, username, passwordHash string) (int, error, string)
}

func (r *AuthDI) Register(email, username, password string) (int, error, string) {
	if !validation.IsEmailValid(email) {
		return http.StatusBadRequest, fmt.Errorf("Invalid email:"), "Некорректный email."
	}
	if len(username) < 4 {
		return http.StatusBadRequest, fmt.Errorf("Username is too short"), "Ошибка. Ваше имя пользователя меньше 4 символов"
	}
	if len(password) < 6 {
		return http.StatusBadRequest, fmt.Errorf("Password is too short"), "Ошибка. Ваш пароль меньше 4 символов"
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return http.StatusInternalServerError, fmt.Errorf("Passwrd encrypting error"), "Внутренняя ошибка сервера"
	}

	err = r.repo.RegisterUser(email, username, string(hashedPassword))
	if err != nil {
		return http.StatusConflict, err, "Данный Email уже зарегистрирован. Попробуйте авторизоваться в системе"
	}

	return http.StatusOK, fmt.Errorf(""), ""
}
