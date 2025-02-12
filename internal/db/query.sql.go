// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.28.0
// source: query.sql

package db

import (
	"context"
)

const createLink = `-- name: CreateLink :exec
INSERT INTO links(id, refers_to, created_by)
VALUES ($1, $2, $3)
`

type CreateLinkParams struct {
	ID        string
	RefersTo  string
	CreatedBy int32
}

func (q *Queries) CreateLink(ctx context.Context, arg CreateLinkParams) error {
	_, err := q.db.ExecContext(ctx, createLink, arg.ID, arg.RefersTo, arg.CreatedBy)
	return err
}

const createUser = `-- name: CreateUser :one
INSERT INTO users(login, name, crypted_password)
VALUES ($1, $2, $3)
RETURNING id
`

type CreateUserParams struct {
	Login           string
	Name            string
	CryptedPassword string
}

func (q *Queries) CreateUser(ctx context.Context, arg CreateUserParams) (int32, error) {
	row := q.db.QueryRowContext(ctx, createUser, arg.Login, arg.Name, arg.CryptedPassword)
	var id int32
	err := row.Scan(&id)
	return id, err
}

const getLinks = `-- name: GetLinks :many
SELECT id, refers_to, created_by, visits, created_at FROM links
WHERE created_by = $1
`

func (q *Queries) GetLinks(ctx context.Context, createdBy int32) ([]Link, error) {
	rows, err := q.db.QueryContext(ctx, getLinks, createdBy)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Link
	for rows.Next() {
		var i Link
		if err := rows.Scan(
			&i.ID,
			&i.RefersTo,
			&i.CreatedBy,
			&i.Visits,
			&i.CreatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}
