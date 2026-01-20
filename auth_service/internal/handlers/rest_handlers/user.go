package rest_handlers

import (
	"articles/internal/kafka"
	"articles/internal/service"
	"articles/pkg/jsonPkg"
	"articles/pkg/jwtPkg"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type AuthHandlerDI struct {
	service       service.Service
	kafkaProducer *kafka.Producer
}

func NewAuthHandler(service service.Service, kafkaProducer *kafka.Producer) *AuthHandlerDI {
	return &AuthHandlerDI{
		service:       service,
		kafkaProducer: kafkaProducer,
	}
}

func (s *AuthHandlerDI) RegisterUser(w http.ResponseWriter, r *http.Request) {
	var req RegisterRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		jsonPkg.SendError(w, fmt.Sprintf("Invalid request body: %v", err), "Невалидное тело запроса", http.StatusBadRequest)
		return
	}

	userID, statusCode, err, userErrorMessage := s.service.Register(req.Email, req.Username, req.Password)
	if err != nil {
		log.Println(err)
		jsonPkg.SendError(w, err.Error(), userErrorMessage, statusCode)
		return
	}
	
	err = s.kafkaProducer.SendUserRegisteredEvent(userID, req.Email, req.Username)
	if err != nil {
		log.Printf("Failed to send Kafka event: %v", err)
	}

	jsonPkg.SendJSON(w, http.StatusCreated, RegisterResponse{Success: true})
}

func (s *AuthHandlerDI) LoginUser(w http.ResponseWriter, r *http.Request) {
	var req LoginRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		jsonPkg.SendError(w, err.Error(), "Неверный формат данных", http.StatusBadRequest)
		return
	}

	statusCode, err, userErrorMessage := s.service.Login(req.Email, req.Password)
	if err != nil {
		jsonPkg.SendError(w, err.Error(), userErrorMessage, statusCode)
		return
	}

	token, err := jwtPkg.GenerateJWT(req.Email)
	if err != nil {
		jsonPkg.SendError(w, err.Error(), "Ошибка сервера", http.StatusInternalServerError)
		return
	}

	jsonPkg.SendJSON(w, http.StatusOK, LoginResponse{Token: token})
}

func (s *AuthHandlerDI) Verify(w http.ResponseWriter, r *http.Request) {
	jsonPkg.SendJSON(w, http.StatusOK, UserVerifyResponse{IsUserVerified: true})
}
