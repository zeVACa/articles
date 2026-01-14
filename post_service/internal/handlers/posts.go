package handlers

import (
	//"articles/internal/models"
	"articles/internal/service"
	"articles/pkg/jsonPkg"
	"encoding/json"
	"fmt"
	"net/http"
	//"time"
)

type PostsHandlerDI struct {
	service service.Service
}

func NewPostsHandler(service service.Service) *PostsHandlerDI {
	return &PostsHandlerDI{
		service: service,
	}
}

func (p *PostsHandlerDI) CreatePost(w http.ResponseWriter, r *http.Request) {
	var req CreatePostRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		jsonPkg.SendError(w, fmt.Sprintf("Invalid request body: %v", err), "Невалидное тело запроса", http.StatusBadRequest)
		return
	}

	post, statusCode, err, errMessage := p.service.CreatePost(req.AuthorId, req.Title, req.Content)
	if err != nil {
		jsonPkg.SendError(w, errMessage, "кккая-то ошибка", statusCode)
		fmt.Println(err)
		return
	}

	jsonPkg.SendJSON(w, 200, CreatePostResponse{Post: post})

	//jsonPkg.SendJSON(w, http.StatusOK, CreatePostResponse{Post: models.Post{
	//	ID:        1,
	//	AuthorID:  2,
	//	Title:     "hello",
	//	Content:   "some long text",
	//	CreatedAt: time.Now(),
	//	UpdatedAt: time.Now(),
	//}})

	// мне нужено проверять авторизацию каждый запрос

	// 1) есть ли такой юзер 2) валидный ли токен
	// Если токен не валидный, то редиректить на логин

	// в сервисе posts должна быть мидлвара которая ходит в сервис auth

	//statusCode, err, userErrorMessage := p.service.Register(req.Email, req.Username, req.Password)
	//if err != nil {
	//	log.Println(err)
	//	jsonPkg.SendError(w, err.Error(), userErrorMessage, statusCode)
	//	return
	//}
	//
	//jsonPkg.SendJSON(w, http.StatusCreated, RegisterResponse{Success: true})
}

func (p *PostsHandlerDI) GetAllPosts(w http.ResponseWriter, r *http.Request) {
	//var req LoginRequest
	//if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
	//	jsonPkg.SendError(w, err.Error(), "Неверный формат данных", http.StatusBadRequest)
	//	return
	//}
	//
	//statusCode, err, userErrorMessage := p.service.Login(req.Email, req.Password)
	//if err != nil {
	//	jsonPkg.SendError(w, err.Error(), userErrorMessage, statusCode)
	//	return
	//}
	//
	//token, err := jwtPkg.GenerateJWT(req.Email)
	//if err != nil {
	//	jsonPkg.SendError(w, err.Error(), "Ошибка сервера", http.StatusInternalServerError)
	//	return
	//}
	//
	//jsonPkg.SendJSON(w, http.StatusOK, LoginResponse{Token: token})
}

func (p *PostsHandlerDI) Test(w http.ResponseWriter, r *http.Request) {
	type Ok struct {
		Ok string `json:"ok"`
	}
	jsonPkg.SendJSON(w, http.StatusOK, Ok{Ok: "hello"})
}
