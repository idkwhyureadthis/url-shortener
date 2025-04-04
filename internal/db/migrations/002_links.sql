-- +goose Up
CREATE TABLE links(
    id TEXT PRIMARY KEY,
    refers_to TEXT NOT NULL,
    created_by BIGINT REFERENCES users(id),
    visits BIGINT NOT NULL DEFAULT 0,
    created_at DATE NOT NULL DEFAULT CURRENT_DATE
);


-- +goose Down
DROP TABLE links;