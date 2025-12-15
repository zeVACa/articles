package storage

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v5"
)

func InitDatabase() (db *pgx.Conn) {
	conn, err := pgx.Connect(context.Background(), "postgres://postgres:postgres@localhost:5433/postdb?sslmode=disable")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}

	return conn
}
