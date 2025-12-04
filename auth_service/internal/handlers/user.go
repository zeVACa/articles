package handlers

import (
	"articles/internal/service"
	"articles/pgk/myjson"
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
			"Ошибка создания пользователя. Email должен быть уникальным",
			http.StatusBadRequest)
		return
	}

	u, err := (*s.service).Register(req.Email, req.Username, req.Password)
	if err != nil {
		myjson.SendError(
			w,
			fmt.Sprintf("Filed to create user: %s", err),
			"Ошибка создания пользователя. Email должен быть уникальным",
			http.StatusBadRequest)
		return
	}

	myjson.SendJSON(w, http.StatusCreated, u)
}
