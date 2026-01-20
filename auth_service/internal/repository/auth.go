package repository

import (
	"articles/internal/models"
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"
)

type AuthRepositoryDI struct {
	conn *pgx.Conn
}

func NewAuthRepository(conn *pgx.Conn) *AuthRepositoryDI {
	return &AuthRepositoryDI{conn: conn}
}

type AuthRepository interface {
	RegisterUser(email, username, hashedPassword string) (int64, error)
	LoginUser(email string) (models.User, error)
}

func (r *AuthRepositoryDI) RegisterUser(email, username, hashedPassword string) (int64, error) {
	var userID int64
	err := r.conn.QueryRow(context.Background(),
		"INSERT INTO users (email, username, password_hash) VALUES ($1, $2, $3) RETURNING id",
		email, username, hashedPassword).Scan(&userID)

	if err != nil {
		return 0, fmt.Errorf("QueryRow failed: %s\n", err)
	}

	return userID, nil
}

func (r *AuthRepositoryDI) LoginUser(email string) (models.User, error) {
	var u models.User

	err := r.conn.QueryRow(context.Background(),
		"SELECT * FROM users WHERE email=$1", email).
		Scan(&u.ID, &u.Email, &u.Username, &u.PasswordHash, &u.EmailVerified, &u.CreatedAt, &u.UpdatedAt)
	if err != nil {
		return models.User{}, err
	}

	return u, nil
}
