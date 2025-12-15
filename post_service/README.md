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

в раннере нужно прописать в поле environment строчку "CONFIG_PATH=./config/local.yaml"

- создать пост
- удалить пост
- редактировать пост
- получить все посты - по передаче page
  page (optional, default: 1) - номер страницы
  limit (optional, default: 10) - количество статей на странице
  Response 200 OK:
- /api/articles/{id} - Получение конкретной статьи
json
{
"articles": [
{
"id": "123e4567-e89b-12d3-a456-426614174000",
"title": "Getting Started with Go",
"content": "Go is a statically typed, compiled programming language...",
"author_id": "550e8400-e29b-41d4-a716-446655440000",
"author_username": "johndoe",
"created_at": "2025-12-15T10:30:00Z",
"updated_at": "2025-12-15T10:30:00Z"
}
],
"pagination": {
"page": 1,
"limit": 10,
"total": 42,
"total_pages": 5
}
}