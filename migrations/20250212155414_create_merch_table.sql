-- +goose Up
-- +goose StatementBegin
CREATE TABLE merch (
    id SERIAL PRIMARY KEY,
    title VARCHAR(255) NOT NULL,
    price INTEGER NOT NULL
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE merch;
-- +goose StatementEnd
