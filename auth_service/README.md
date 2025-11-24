auth-service/
├── cmd/
│   └── main.go
├── internal/
│   ├── handler/
│   ├── service/
│   ├── repository/
│   └── models/
├── migrations/
│   ├── 000001_create_users_table.up.sql
│   └── 000001_create_users_table.down.sql
├── proto/
├── .env
├── .env.example
├── go.mod
├── go.sum
├── Dockerfile
└── README.md

router -> handler -> service -> repository.