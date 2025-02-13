-- +goose Up
CREATE TABLE inventory_ (
    id SERIAL PRIMARY KEY,
    id_user INTEGER NOT NULL REFERENCES user_(id) ON DELETE CASCADE,
    id_merch INTEGER NOT NULL REFERENCES merch_(id) ON DELETE CASCADE,
    quantity INTEGER NOT NULL DEFAULT 1
);

INSERT INTO merch_ (title, price)
VALUES 
    ('t-shirt', 80),
    ('cup', 20),
    ('book', 50),
    ('pen', 10),
    ('powerbank', 200),
    ('hoody', 300),
    ('umbrella', 200),
    ('socks', 10),
    ('wallet', 50),
    ('pink-hoody', 500)
ON CONFLICT (title) DO NOTHING;

-- +goose Down
DELETE FROM merch_ WHERE title IN ('t-shirt', 'cup', 'book', 'pen', 'powerbank', 'hoody', 'umbrella', 'socks', 'wallet', 'pink-hoody');
DROP TABLE inventory_;
