package service

import (
	"articles/internal/models"
	"articles/internal/repository"
	"net/http"
)

type PostsDI struct {
	repo repository.PostsRepository
}

func NewPostsService(repo repository.PostsRepository) *PostsDI {
	return &PostsDI{
		repo: repo,
	}
}

type Service interface {
	GetAllPosts() ([]models.Post, int, error, string)
	GetPostById(id int64) (models.Post, int, error, string)
	CreatePost(authorId int64, title, content string) (models.Post, int, error, string)
}

func (p *PostsDI) GetAllPosts() ([]models.Post, int, error, string) {
	posts, err := p.repo.GetAllPosts()
	if err != nil {
		return []models.Post{}, http.StatusInternalServerError, err, "Ошибка сервера"
	}

	return posts, 0, nil, ""
}

func (p *PostsDI) GetPostById(id int64) (models.Post, int, error, string) {
	post, err := p.repo.GetPostById(id)
	if err != nil {
		return models.Post{}, http.StatusInternalServerError, err, "Ошибка сервера"
	}

	return post, 0, nil, ""
}

func (p *PostsDI) CreatePost(authorId int64, title, content string) (models.Post, int, error, string) {
	post, err := p.repo.CreatePost(authorId, title, content)
	if err != nil {
		return models.Post{}, http.StatusInternalServerError, err, "Ошибка сервера"
	}

	return post, http.StatusOK, nil, ""
}

//func (a *PostsDI) Register(email, username, password string) (int, error, string) {
//	if !validation.IsEmailValid(email) {
//		return http.StatusBadRequest, fmt.Errorf("Invalid email:"), "Некорректный email."
//	}
//	if len(username) < 4 {
//		return http.StatusBadRequest, fmt.Errorf("Username is too short"), "Ошибка. Ваше имя пользователя меньше 4 символов"
//	}
//	if len(password) < 6 {
//		return http.StatusBadRequest, fmt.Errorf("Password is too short"), "Ошибка. Ваш пароль меньше 6 символов"
//	}
//
//	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
//	if err != nil {
//		return http.StatusInternalServerError, fmt.Errorf("Passwrd encrypting error"), "Внутренняя ошибка сервера"
//	}
//
//	err = a.repo.RegisterUser(email, username, string(hashedPassword))
//	if err != nil {
//		return http.StatusConflict, err, "Данный Email уже зарегистрирован. Попробуйте авторизоваться в системе"
//	}
//
//	return http.StatusOK, nil, ""
//}

//func (a *PostsDI) Login(email, password string) (int, error, string) {
//	u, err := a.repo.LoginUser(email)
//	if err != nil {
//		return http.StatusInternalServerError, err, "Ошибка сервера"
//	}
//
//	err = bcrypt.CompareHashAndPassword([]byte(u.PasswordHash), []byte(password))
//	if err != nil {
//		return http.StatusUnauthorized, fmt.Errorf("wrong login or password"), "Неправильный логин или пароль"
//	}
//
//	return http.StatusOK, nil, ""
//}
