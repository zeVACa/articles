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
	GetAllPosts() ([]models.Post, error)
	GetPostById(id int64) (models.Post, error)
	CreatePost(authorId int64, title, content string) (models.Post, error)
}

func (r *PostsRepositoryDI) GetAllPosts() ([]models.Post, error) {
	var posts []models.Post

	rows, err := r.conn.Query(context.Background(), "SELECT * from posts")
	if err != nil {
		return nil, fmt.Errorf("Query failed: %w", err)
	}
	defer rows.Close()

	for rows.Next() {
		var p models.Post

		err = rows.Scan(&p.ID, &p.AuthorID, &p.Title, &p.Content, &p.CreatedAt, &p.UpdatedAt)
		if err != nil {
			return nil, fmt.Errorf("Scan failed: %w", err)
		}

		posts = append(posts, p)
	}
	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("Rows iteration failed: %w", err)
	}

	return posts, nil
}

func (r *PostsRepositoryDI) GetPostById(id int64) (models.Post, error) {
	var p models.Post

	err := r.conn.QueryRow(context.Background(),
		"SELECT * FROM posts WHERE id=$1", id).Scan(&p.ID, &p.AuthorID, &p.Title, &p.Content, &p.CreatedAt, &p.UpdatedAt)
	if err != nil {
		return models.Post{}, fmt.Errorf("QueryRow failed: %s\n", err)
	}

	return p, nil
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
