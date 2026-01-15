package rest_handlers

import (
	"articles/internal/service"
	authClient "articles/pkg/grpc/auth_client"
	"articles/pkg/jsonPkg"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/go-chi/chi/v5"
)

type PostsHandlerDI struct {
	service        service.Service
	authGrpcClient *authClient.AuthClient
}

func NewPostsHandler(service service.Service, authGrpcClient *authClient.AuthClient) *PostsHandlerDI {
	return &PostsHandlerDI{
		service:        service,
		authGrpcClient: authGrpcClient,
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

	authHeader := r.Header.Get("Authorization")
	if authHeader == "" {
		log.Println("Missing Authorization header")
		jsonPkg.SendError(w, "Authorization header is required", "Требуется авторизация", http.StatusUnauthorized)
		return
	}

	token := strings.TrimPrefix(authHeader, "Bearer ")
	if token == authHeader {
		log.Println("Invalid Authorization format")
		jsonPkg.SendError(w, "Invalid Authorization format", "Неверный формат токена", http.StatusUnauthorized)
		return
	}

	isTokenVerified, userID, err := p.authGrpcClient.VerifyToken(token)
	if err != nil {
		log.Printf("gRPC error: %v", err)
		jsonPkg.SendError(w, err.Error(), "Сервис авторизации недоступен", http.StatusServiceUnavailable)
		return
	}

	if !isTokenVerified {
		jsonPkg.SendError(w, "Token is invalid", "Невалидный токен", http.StatusUnauthorized)
		return
	}

	post, statusCode, err, errMessage := p.service.CreatePost(userID, req.Title, req.Content)
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
