package handlers

import (
	"articles/internal/models"
)

type GetAllPostsResponse struct {
	Posts []models.Post `json:"posts"`
}

type GetPostByIdResponse struct {
	Post models.Post `json:"post"`
}

type CreatePostRequest struct {
	Title   string `json:"title"`
	Content string `json:"content"`
}

type CreatePostResponse struct {
	Post models.Post `json:"post"`
}
