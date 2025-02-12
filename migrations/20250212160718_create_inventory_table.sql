-- +goose Up
-- +goose StatementBegin
CREATE TABLE inventory (
    id SERIAL PRIMARY KEY,
    id_user INTEGER NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    id_merch INTEGER NOT NULL REFERENCES merch(id) ON DELETE CASCADE,
    quantity INTEGER NOT NULL DEFAULT 1
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE inventory;
-- +goose StatementEnd
