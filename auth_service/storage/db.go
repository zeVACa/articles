package storage

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5"
	"os"
)

func InitDatabase() {
	conn, err := pgx.Connect(context.Background(), "postgres://postgres:postgres@localhost:5432/postgres?sslmode=disable")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	defer conn.Close(context.Background())

	//err = InitDatabase(conn, "storage/migrations/000001_create_users_table.up.sql")
}
