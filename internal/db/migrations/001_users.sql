-- +goose Up
CREATE TABLE USERS(
    id SERIAL PRIMARY KEY,
    login TEXT UNIQUE NOT NULL,
    name TEXT NOT NULL,
    crypted_password TEXT NOT NULL,
    subscription_type TEXT NOT NULL DEFAULT 'FREE'
);



-- +goose Down
DROP TABLE users;