package handlers

import (
	//"articles/internal/models"
	"articles/internal/service"
	"articles/pkg/jsonPkg"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	//"time"

	"github.com/go-chi/chi/v5"
)

type PostsHandlerDI struct {
	service service.Service
}

func NewPostsHandler(service service.Service) *PostsHandlerDI {
	return &PostsHandlerDI{
		service: service,
	}
}

func (p *PostsHandlerDI) GetAllPosts(w http.ResponseWriter, r *http.Request) {
	posts, statusCode, err, errMessage := p.service.GetAllPosts()
	if err != nil {
		jsonPkg.SendError(w, errMessage, "Ошибка сервера", statusCode)
	}

	jsonPkg.SendJSON(w, http.StatusOK, GetAllPostsResponse{Posts: posts})
}

func (p *PostsHandlerDI) GetPostById(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")

	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		http.Error(w, "Invalid post ID", http.StatusBadRequest)
		return
	}

	post, statusCode, err, errMessage := p.service.GetPostById(id)
	if err != nil {
		jsonPkg.SendError(w, errMessage, "Ошибка сервера", statusCode)
		return
	}

	jsonPkg.SendJSON(w, http.StatusOK, GetPostByIdResponse{Post: post})
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
		jsonPkg.SendError(w, errMessage, "Ошибка сервера", statusCode)
		fmt.Println(err)
		return
	}

	jsonPkg.SendJSON(w, 200, CreatePostResponse{Post: post})
}

func (p *PostsHandlerDI) Test(w http.ResponseWriter, r *http.Request) {
	type Ok struct {
		Ok string `json:"ok"`
	}
	jsonPkg.SendJSON(w, http.StatusOK, Ok{Ok: "hello"})
}
