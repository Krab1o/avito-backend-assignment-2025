-- +goose Up
-- +goose StatementBegin
CREATE TABLE user_transaction (
    id SERIAL PRIMARY KEY,
    id_sender INTEGER NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    id_receiver INTEGER NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    amount INTEGER NOT NULL CHECK (amount > 0)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE user_transaction;
-- +goose StatementEnd
