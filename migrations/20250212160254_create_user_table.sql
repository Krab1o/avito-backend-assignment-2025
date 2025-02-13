-- +goose Up
CREATE TABLE user_ (
    id SERIAL PRIMARY KEY,
    username VARCHAR(100) UNIQUE NOT NULL,
    password_hash TEXT NOT NULL,
    coins INTEGER NOT NULL
);

-- +goose Down
DROP TABLE user_;
