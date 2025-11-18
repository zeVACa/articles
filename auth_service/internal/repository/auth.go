package repository

import (
	"articles/internal/models"
	"context"
	"fmt"
	"github.com/jackc/pgx/v5/pgxpool"
	"os"
)

//type AuthRepository interface {
//	RegisterUser() (models.User, error)
//	LoginUser() (models.User, error)
//}

type AuthRepositoryDI struct {
	pool *pgxpool.Pool
}

func NewAuthRepository(pool *pgxpool.Pool) *AuthRepositoryDI {
	return &AuthRepositoryDI{pool: pool}
}

type AuthRepository interface {
	RegisterUser() (models.User, error)
}

func (r *AuthRepositoryDI) RegisterUser() {
	var email, username, passwordHash string

	err := r.pool.QueryRow(context.Background(),
		"SELECT email, username, password_hash FROM users WHERE id=$1", 1).
		Scan(&email, &username, &passwordHash)
	if err != nil {
		fmt.Fprintf(os.Stderr, "QueryRow failed: %v\n", err)
		os.Exit(1)
	}

	//err = conn.QueryRow(context.Background(),
	//		"SELECT email, username, password_hash FROM users WHERE id=$1", 1).
	//		Scan(&email, &username, &passwordHash)
	//	if err != nil {
	//		fmt.Fprintf(os.Stderr, "QueryRow failed: %v\n", err)
	//		os.Exit(1)
	//	}
}
