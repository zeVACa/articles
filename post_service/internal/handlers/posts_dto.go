package handlers

import (
	"articles/internal/models"
)

type CreatePostRequest struct {
	AuthorId int64  `json:"author_id"`
	Title    string `json:"title"`
	Content  string `json:"content"`
}

type CreatePostResponse struct {
	Post models.Post `json:"post"`
}
