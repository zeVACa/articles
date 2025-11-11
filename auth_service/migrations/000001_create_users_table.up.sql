CREATE TABLE iF NOT EXISTS users (
    id serial primary key,
    email varchar(255) UNIQUE NOT NULL,
    username varchar(255) not null,
    password_hash varchar(255) not null,
    created_at timestamp default CURRENT_TIMESTAMP
)