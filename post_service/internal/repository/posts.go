package repository

import (
	"articles/internal/models"
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"
)

type PostsRepositoryDI struct {
	conn *pgx.Conn
}

func NewPostsRepository(conn *pgx.Conn) *PostsRepositoryDI {
	return &PostsRepositoryDI{conn: conn}
}

type PostsRepository interface {
	//RegisterUser(email, username, hashedPassword string) error
	//LoginUser(email string) (models.User, error)
	CreatePost(authorId int64, title, content string) (models.Post, error)
}

func (r *PostsRepositoryDI) CreatePost(authorId int64, title, content string) (models.Post, error) {
	var p models.Post

	err := r.conn.QueryRow(context.Background(),
		"INSERT INTO posts (author_id, title, content) VALUES ($1, $2, $3) RETURNING id, author_id, title, content, created_at, updated_at", authorId, title, content).
		Scan(&p.ID, &p.AuthorID, &p.Title, &p.Content, &p.CreatedAt, &p.UpdatedAt)
	if err != nil {
		return models.Post{}, fmt.Errorf("QueryRow failed: %s\n", err)
	}

	return p, nil
}
