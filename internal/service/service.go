package service

import (
	"database/sql"
	"log"

	"github.com/idkwhyureadthis/url-shortener/internal/db"

	_ "github.com/jackc/pgx/stdlib"
)

type Service struct {
	conn *db.Queries
}

func New(connUrl string) Service {
	conn, err := sql.Open("pgx", connUrl)
	if err != nil {
		log.Fatal(err)
	}
	s := Service{}

	s.conn = db.New(conn)

	return s
}
