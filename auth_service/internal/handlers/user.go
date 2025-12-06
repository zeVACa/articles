package handlers

import (
	"articles/internal/service"
	"articles/pgk/myjson"
	"articles/pgk/myjwt"
	"encoding/json"
	"fmt"
	"log"
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
		log.Println(err)
		myjson.SendError(w, err.Error(), userErrorMessage, statusCode)
		return
	}

	myjson.SendJSON(w, http.StatusCreated, RegisterResponse{Success: true})
}

func (s *AuthHandlerDI) LoginUser(w http.ResponseWriter, r *http.Request) {
	var req LoginRequest
	json.NewDecoder(r.Body).Decode(&req)

	//myjwt.GenerateJWT()

	statusCode, err, userErrorMessage := (*s.service).Login(req.Email, req.Password)
	if err != nil {
		myjson.SendError(w, err.Error(), userErrorMessage, statusCode)
		return
	}

	token, err := myjwt.GenerateJWT(req.Email)
	if err != nil {
		myjson.SendError(w, err.Error(), "Ошибка сервера", http.StatusInternalServerError)
		return
	}

	myjson.SendJSON(w, http.StatusOK, LoginResponse{Token: token})
}
