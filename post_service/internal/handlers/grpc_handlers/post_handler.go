package grpc_handlers

//
//import (
//	"encoding/json"
//	"log"
//	"net/http"
//	"strings"
//
//	"articles/internal/models"
//	"articles/internal/service"
//	"articles/pkg/authclient" // ← Импорт gRPC клиента
//)
//
//type PostHandler struct {
//	service    *service.PostService
//	authClient *authclient.AuthClient // ← Добавили gRPC клиент
//}
//
//func NewPostHandler(s *service.PostService, ac *authclient.AuthClient) *PostHandler {
//	return &PostHandler{
//		service:    s,
//		authClient: ac,
//	}
//}
//
//// CreatePost - создание поста (защищенный маршрут)
//func (h *PostHandler) CreatePost(w http.ResponseWriter, r *http.Request) {
//	// 1. Получаем токен из заголовка
//	authHeader := r.Header.Get("Authorization")
//	if authHeader == "" {
//		log.Println("❌ Missing Authorization header")
//		http.Error(w, "Missing Authorization header", http.StatusUnauthorized)
//		return
//	}
//
//	token := strings.TrimPrefix(authHeader, "Bearer ")
//	if token == authHeader {
//		log.Println("❌ Invalid Authorization format")
//		http.Error(w, "Invalid Authorization format. Use: Bearer <token>", http.StatusUnauthorized)
//		return
//	}
//
//	// 2. Проверяем токен через gRPC вызов к auth_service
//	log.Printf("🔍 Verifying token via gRPC...")
//	valid, userID, err := h.authClient.VerifyToken(token)
//	if err != nil {
//		log.Printf("❌ gRPC error: %v", err)
//		http.Error(w, "Auth service unavailable", http.StatusServiceUnavailable)
//		return
//	}
//
//	if !valid {
//		log.Println("❌ Invalid token")
//		http.Error(w, "Invalid token", http.StatusUnauthorized)
//		return
//	}
//
//	log.Printf("✅ Token valid, user_id: %d", userID)
//
//	// 3. Парсим тело запроса
//	var req model.CreatePostRequest
//	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
//		log.Printf("❌ Invalid request body: %v", err)
//		http.Error(w, "Invalid request body", http.StatusBadRequest)
//		return
//	}
//
//	// 4. Создаем пост
//	post := &model.Post{
//		Title:   req.Title,
//		Content: req.Content,
//		UserID:  userID, // ← Используем user_id из токена, а не из запроса
//	}
//
//	if err := h.service.CreatePost(post); err != nil {
//		log.Printf("❌ Failed to create post: %v", err)
//		http.Error(w, "Failed to create post", http.StatusInternalServerError)
//		return
//	}
//
//	log.Printf("✅ Post created: ID=%d, UserID=%d, Title=%s", post.ID, post.UserID, post.Title)
//
//	w.Header().Set("Content-Type", "application/json")
//	w.WriteHeader(http.StatusCreated)
//	json.NewEncoder(w).Encode(post)
//}
//
//// GetPosts - получение всех постов (открытый маршрут)
//func (h *PostHandler) GetPosts(w http.ResponseWriter, r *http.Request) {
//	posts, err := h.service.GetAllPosts()
//	if err != nil {
//		log.Printf("❌ Failed to get posts: %v", err)
//		http.Error(w, "Failed to get posts", http.StatusInternalServerError)
//		return
//	}
//
//	log.Printf("✅ Retrieved %d posts", len(posts))
//
//	w.Header().Set("Content-Type", "application/json")
//	json.NewEncoder(w).Encode(posts)
//}
