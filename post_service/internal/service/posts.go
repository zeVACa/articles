package service

import (
	"articles/internal/models"
	"articles/internal/repository"
	"context"
	"log"
	"net/http"
)

type PostsDI struct {
	repo  repository.PostsRepository
	cache *CacheService
}

func NewPostsService(repo repository.PostsRepository, cache *CacheService) *PostsDI {
	return &PostsDI{
		repo:  repo,
		cache: cache,
	}
}

type Service interface {
	GetAllPosts() ([]models.Post, int, error, string)
	GetPostById(id int64) (models.Post, int, error, string)
	CreatePost(authorId int64, title, content string) (models.Post, int, error, string)
}

func (p *PostsDI) GetAllPosts() ([]models.Post, int, error, string) {
	ctx := context.Background()

	cachedPosts, err := p.cache.GetAllPosts(ctx)
	if err != nil {
		log.Printf("cache error: %v", err)
	}
	if cachedPosts != nil {
		log.Println("cache HIT: all posts")
		return cachedPosts, 0, nil, ""
	}

	log.Println("cache MISS: all posts")
	posts, err := p.repo.GetAllPosts()
	if err != nil {
		return []models.Post{}, http.StatusInternalServerError, err, "Ошибка сервера"
	}

	if err := p.cache.SetAllPosts(ctx, posts); err != nil {
		log.Printf("failed to cache all posts: %v", err)
	}

	return posts, 0, nil, ""
}

func (p *PostsDI) GetPostById(id int64) (models.Post, int, error, string) {
	ctx := context.Background()

	cachedPost, err := p.cache.GetPost(ctx, id)
	if err != nil {
		log.Printf("cache error: %v", err)
	}
	if cachedPost != nil {
		log.Printf("cache HIT: post %d", id)
		return *cachedPost, 0, nil, ""
	}

	log.Printf("cache MISS: post %d", id)
	post, err := p.repo.GetPostById(id)
	if err != nil {
		return models.Post{}, http.StatusInternalServerError, err, "Ошибка сервера"
	}

	if err := p.cache.SetPost(ctx, post); err != nil {
		log.Printf("failed to cache post %d: %v", id, err)
	}

	return post, 0, nil, ""
}

func (p *PostsDI) CreatePost(authorId int64, title, content string) (models.Post, int, error, string) {
	ctx := context.Background()

	post, err := p.repo.CreatePost(authorId, title, content)
	if err != nil {
		return models.Post{}, http.StatusInternalServerError, err, "Ошибка сервера"
	}

	if err := p.cache.InvalidateAllPosts(ctx); err != nil {
		log.Printf("failed to invalidate cache: %v", err)
	}

	return post, http.StatusOK, nil, ""
}
