-- +goose Up
CREATE TABLE links(
    id TEXT NOT NULL,
    refers_to TEXT NOT NULL,
    created_by INTEGER REFERENCES users(id) NOT NULL,
    visits BIGINT NOT NULL DEFAULT 0,
    created_at DATE NOT NULL DEFAULT CURRENT_DATE
);


-- +goose Down
DROP TABLE links