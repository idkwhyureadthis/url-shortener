-- name: GetLinks :many
SELECT * FROM links
WHERE created_by = $1;


-- name: CreateLink :one
INSERT INTO links(id, refers_to, created_by)
VALUES ($1, $2, $3)
RETURNING id;

-- name: GetLink :one
SELECT refers_to FROM links
WHERE id = $1;