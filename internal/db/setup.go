package db

import (
	"database/sql"

	"github.com/pressly/goose/v3"
)

func SetupMigrations(connUrl string) error {
	conn, err := sql.Open("pgx", connUrl)
	if err != nil {
		return err
	}
	goose.SetDialect("postgres")
	err = goose.Up(conn, "internal/db/migrations")
	if err != nil {
		return err
	}
	return nil
}
