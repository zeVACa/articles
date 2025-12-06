package service

import (
	"articles/internal/repository"
	"articles/pgk/validation"
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
	Register(email, username, password string) (int, error, string)
	Login(email, password string) (int, error, string)
}

func (a *AuthDI) Register(email, username, password string) (int, error, string) {
	if !validation.IsEmailValid(email) {
		return http.StatusBadRequest, fmt.Errorf("Invalid email:"), "Некорректный email."
	}
	if len(username) < 4 {
		return http.StatusBadRequest, fmt.Errorf("Username is too short"), "Ошибка. Ваше имя пользователя меньше 4 символов"
	}
	if len(password) < 6 {
		return http.StatusBadRequest, fmt.Errorf("Password is too short"), "Ошибка. Ваш пароль меньше 6 символов"
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return http.StatusInternalServerError, fmt.Errorf("Passwrd encrypting error"), "Внутренняя ошибка сервера"
	}

	err = a.repo.RegisterUser(email, username, string(hashedPassword))
	if err != nil {
		return http.StatusConflict, err, "Данный Email уже зарегистрирован. Попробуйте авторизоваться в системе"
	}

	return http.StatusOK, nil, ""
}

func (a *AuthDI) Login(email, password string) (int, error, string) {
	u, err := a.repo.LoginUser(email)
	if err != nil {
		return http.StatusInternalServerError, err, "Ошибка сервера2"
	}

	err = bcrypt.CompareHashAndPassword([]byte(u.PasswordHash), []byte(password))
	fmt.Println("passwords", u.PasswordHash, "---", password)
	if err != nil {
		return http.StatusUnauthorized, fmt.Errorf("Wrong login or password"), "Неправильный логин или пароль"
	}

	return http.StatusOK, nil, ""
}
