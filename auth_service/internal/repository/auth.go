package repository

import (
	"articles/internal/models"
	"context"
	"fmt"
	"github.com/jackc/pgx/v5"
)

//type AuthRepository interface {
//	RegisterUser() (models.User, error)
//	LoginUser() (models.User, error)
//}

type AuthRepositoryDI struct {
	conn *pgx.Conn
}

func NewAuthRepository(conn *pgx.Conn) *AuthRepositoryDI {
	return &AuthRepositoryDI{conn: conn}
}

type AuthRepository interface {
	RegisterUser(email, username, passwordHash string) (*models.User, error)
}

func (r *AuthRepositoryDI) RegisterUser(email, username, passwordHash string) (*models.User, error) {
	var u models.User

	err := r.conn.QueryRow(context.Background(),
		"INSERT INTO users (email, username, password_hash) VALUES ($1, $2, $3) RETURNING id, email, username, password_hash, email_verified, created_at, updated_at", email, username, passwordHash).
		Scan(&u.ID, &u.Email, &u.Username, &u.PasswordHash, &u.EmailVerified, &u.CreatedAt, &u.UpdatedAt)

	if err != nil {
		return nil, fmt.Errorf("QueryRow failed: %v\n", err)
	}

	return &u, nil
}
