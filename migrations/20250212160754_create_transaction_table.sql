-- +goose Up
CREATE TABLE user_transaction_ (
    id SERIAL PRIMARY KEY,
    id_sender INTEGER NOT NULL REFERENCES user_(id) ON DELETE CASCADE,
    id_receiver INTEGER NOT NULL REFERENCES user_(id) ON DELETE CASCADE,
    amount INTEGER NOT NULL CHECK (amount > 0)
);

-- +goose Down
DROP TABLE user_transaction_;
