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
