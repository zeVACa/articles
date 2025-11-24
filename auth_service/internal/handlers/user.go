package handlers

import (
	"articles/internal/service"
	"encoding/json"
	"fmt"
	"net/http"
)

type RegisterUserHandlerDI struct {
	service service.Service
}

func NewRegisterUserHandler(service service.Service) *RegisterUserHandlerDI {
	return &RegisterUserHandlerDI{
		service: service,
	}
}

func (s *RegisterUserHandlerDI) RegisterUser(w http.ResponseWriter, r *http.Request) {
	var req RegisterUserRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, fmt.Sprintf("Invalid request body: %v", err), http.StatusBadRequest)
	}

	u, err := s.service.Register(req.Email, req.Username, req.Password)
	if err != nil {
		fmt.Println(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	err = json.NewEncoder(w).Encode(u)
	if err != nil {
		fmt.Println(err)
	}
}
