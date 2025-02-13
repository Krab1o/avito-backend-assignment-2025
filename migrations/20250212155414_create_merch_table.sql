-- +goose Up
CREATE TABLE merch_ (
    id SERIAL PRIMARY KEY,
    title VARCHAR(255) UNIQUE NOT NULL,
    price INTEGER NOT NULL
);

-- +goose Down
DROP TABLE merch_;
