package service

import (
	"context"
	"database/sql"
	"log"
	"net/http"

	"github.com/idkwhyureadthis/url-shortener/shortener/internal/db"
	"github.com/idkwhyureadthis/url-shortener/shortener/internal/models"
)

type RedirectService struct {
	conn *db.Queries
}

func NewRedirectService(connUrl string) *RedirectService {
	conn, err := sql.Open("pgx", connUrl)
	if err != nil {
		log.Fatal(err)
	}
	service := RedirectService{}
	service.conn = db.New(conn)
	return &service
}

func (r *RedirectService) GetFullLink(id string) models.RedirectLinkData {
	data := models.RedirectLinkData{}
	link, err := r.conn.GetLink(context.Background(), id)
	if err != nil {
		data.Code = http.StatusInternalServerError
		data.Err = err
	} else if link == "" {
		data.Code = http.StatusNotFound
		data.Err = models.ErrNotFound
	}
	data.FullLink = link
	return data
}
