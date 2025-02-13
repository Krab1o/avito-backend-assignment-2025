-- +goose Up
CREATE TABLE merch (
    id SERIAL PRIMARY KEY,
    title VARCHAR(255) NOT NULL,
    price INTEGER NOT NULL
);

-- +goose Down
DROP TABLE merch;
