-- +goose Up
-- +goose StatementBegin
CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    nickname VARCHAR(100) UNIQUE NOT NULL,
    coins INTEGER NOT NULL DEFAULT 0  -- Users start with 0 coins
);

-- Insert the system user (id = 1)
INSERT INTO users (id, nickname, coins) VALUES (1, 'SYSTEM', 0) 
ON CONFLICT DO NOTHING;  -- Ensures the system user is only created once
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE users;
-- +goose StatementEnd
