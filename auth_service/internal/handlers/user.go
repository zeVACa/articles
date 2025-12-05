package handlers

import (
	"articles/internal/service"
	"articles/pgk/myjson"
	"articles/pgk/validation"
	"encoding/json"
	"fmt"
	"net/http"
)

type RegisterUserHandlerDI struct {
	service *service.Service
}

func NewRegisterUserHandler(service service.Service) *RegisterUserHandlerDI {
	return &RegisterUserHandlerDI{
		service: &service,
	}
}

func (s *RegisterUserHandlerDI) RegisterUser(w http.ResponseWriter, r *http.Request) {
	var req RegisterUserRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		myjson.SendError(
			w,
			fmt.Sprintf("Invalid request body: %v", err),
			"Невалидное тело запроса",
			http.StatusBadRequest)
		return
	}

	if !validation.IsEmailValid(req.Email) {
		myjson.SendError(
			w,
			"Invalid email:",
			"Некорректный email.",
			http.StatusBadRequest)
		return
	}

	if len(req.Password) < 6 {
		myjson.SendError(
			w,
			"Password too short",
			"Ошибка. Ваш пароль меньше 6 символов",
			http.StatusBadRequest)
		return
	}

	if len(req.Username) < 4 {
		myjson.SendError(
			w,
			"Username is too short",
			"Ошибка. Ваше имя пользователя меньше 4 символов",
			http.StatusBadRequest)
		return
	}

	_, err = (*s.service).Register(req.Email, req.Username, req.Password)
	if err != nil {
		myjson.SendError(
			w,
			fmt.Sprintf("Filed to create user: %s", err),
			"Ошибка создания пользователя. Email должен быть уникальным",
			http.StatusConflict)
		return
	}

	res := RegisterUserResponse{
		Success: true,
	}

	myjson.SendJSON(w, http.StatusCreated, res)
}
