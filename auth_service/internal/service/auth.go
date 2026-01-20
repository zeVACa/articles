package service

import (
	"articles/internal/repository"
	"articles/pkg/validation"
	"fmt"
	"net/http"

	"golang.org/x/crypto/bcrypt"
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
	Register(email, username, password string) (int64, int, error, string)
	Login(email, password string) (int, error, string)
}

func (a *AuthDI) Register(email, username, password string) (int64, int, error, string) {
	if !validation.IsEmailValid(email) {
		return 0, http.StatusBadRequest, fmt.Errorf("Invalid email:"), "Некорректный email."
	}
	if len(username) < 4 {
		return 0, http.StatusBadRequest, fmt.Errorf("Username is too short"), "Ошибка. Ваше имя пользователя меньше 4 символов"
	}
	if len(password) < 6 {
		return 0, http.StatusBadRequest, fmt.Errorf("Password is too short"), "Ошибка. Ваш пароль меньше 6 символов"
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return 0, http.StatusInternalServerError, fmt.Errorf("Passwrd encrypting error"), "Внутренняя ошибка сервера"
	}

	userID, err := a.repo.RegisterUser(email, username, string(hashedPassword))
	if err != nil {
		return 0, http.StatusConflict, err, "Данный Email уже зарегистрирован. Попробуйте авторизоваться в системе"
	}

	return userID, http.StatusOK, nil, ""
}

func (a *AuthDI) Login(email, password string) (int, error, string) {
	u, err := a.repo.LoginUser(email)
	if err != nil {
		return http.StatusInternalServerError, err, "Ошибка сервера"
	}

	err = bcrypt.CompareHashAndPassword([]byte(u.PasswordHash), []byte(password))
	if err != nil {
		return http.StatusUnauthorized, fmt.Errorf("wrong login or password"), "Неправильный логин или пароль"
	}

	return http.StatusOK, nil, ""
}
