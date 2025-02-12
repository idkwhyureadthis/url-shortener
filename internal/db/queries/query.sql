-- name: CreateUser :one
INSERT INTO users(login, name, crypted_password)
VALUES ($1, $2, $3)
RETURNING id;


-- name: GetLinks :many
SELECT * FROM links
WHERE created_by = $1;


-- name: CreateLink :exec
INSERT INTO links(id, refers_to, created_by)
VALUES ($1, $2, $3);