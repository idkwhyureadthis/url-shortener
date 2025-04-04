-- +goose Up
CREATE TABLE users(
    id BIGSERIAL PRIMARY KEY,
    login TEXT UNIQUE NOT NULL,
    name TEXT NOT NULL,
    crypted_refresh TEXT,
    crypted_password TEXT NOT NULL
);

-- +goose Down
DROP TABLE users;