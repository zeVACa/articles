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

в раннере нужно прописать в поле environment строчку "CONFIG_PATH=./config/local.yaml

Ниже перечислены ручки каждого из сервисов

1) Auth
    - register/
    - login/
2) Post
   - create/
   - update/
   - delete/
   - getOne/
   - GetAll/ - pagination
   
    Models
   - id
   - author_id  
   - title
   - content
   - created_at
   - updated_at
3) Subscription
   - subscribe/ - 1 to many
   - unsubscribe/
4) Notification
5) Comments