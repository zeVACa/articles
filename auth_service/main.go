package main

import (
	"articles/internal/router"
	"articles/storage"
	"net/http"
)

func main() {

	router.ImplementRouter()
	storage.InitDatabase()

	http.ListenAndServe("localhost:8080", nil)
}

//package main
//
//import (
//"context"
//"fmt"
//"github.com/jackc/pgx/v5"
//"os"
//)
//
//func main() {
//	conn, err := pgx.Connect(context.Background(), "postgres://postgres:postgres@localhost:5432/postgres?sslmode=disable")
//	if err != nil {
//		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
//		os.Exit(1)
//	}
//	defer conn.Close(context.Background())
//
//	/*var email, username, passwordHash string
//
//	err = conn.QueryRow(context.Background(),
//		"SELECT email, username, password_hash FROM users WHERE id=$1", 1).
//		Scan(&email, &username, &passwordHash)
//	if err != nil {
//		fmt.Fprintf(os.Stderr, "QueryRow failed: %v\n", err)
//		os.Exit(1)
//	}
//
//	fmt.Printf("Email: %s, Username: %s, Password Hash: %s\n", email, username, passwordHash)*/
//}
