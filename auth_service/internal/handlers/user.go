package handlers

import (
	"articles/internal/service"
	"articles/pgk/myjson"
	"encoding/json"
	"fmt"
	"net/http"
)

type AuthHandlerDI struct {
	service *service.Service
}

func NewAuthHandler(service service.Service) *AuthHandlerDI {
	return &AuthHandlerDI{
		service: &service,
	}
}

func (s *AuthHandlerDI) RegisterUser(w http.ResponseWriter, r *http.Request) {
	var req RegisterRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		myjson.SendError(w, fmt.Sprintf("Invalid request body: %v", err), "Невалидное тело запроса", http.StatusBadRequest)
		return
	}

	statusCode, err, userErrorMessage := (*s.service).Register(req.Email, req.Username, req.Password)
	if err != nil {
		myjson.SendError(w, err.Error(), userErrorMessage, statusCode)
		return
	}

	res := RegisterResponse{
		Success: true,
	}

	myjson.SendJSON(w, http.StatusCreated, res)
}

func (s *AuthHandlerDI) LoginUser(w http.ResponseWriter, r *http.Request) {
	var req LoginRequest
	json.NewDecoder(r.Body).Decode(&req)

	//myjwt.GenerateJWT()

	myjson.SendJSON(w, http.StatusOK, req)

	fmt.Println(req)
}
