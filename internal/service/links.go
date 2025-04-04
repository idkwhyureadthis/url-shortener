package service

import (
	"context"
	"database/sql"
	"log"
	"net/url"
	"strings"

	"github.com/idkwhyureadthis/url-shortener/shortener/internal/db"
	"github.com/idkwhyureadthis/url-shortener/shortener/internal/models"
	"github.com/idkwhyureadthis/url-shortener/shortener/pkg/linkgen"
	"github.com/idkwhyureadthis/url-shortener/shortener/pkg/linkverify"

	_ "github.com/jackc/pgx/v5/stdlib"
)

type LinksService struct {
	redirectServiceUrl string
	conn               *db.Queries
}

func NewLinksService(connUrl, redirectUrl string) *LinksService {
	conn, err := sql.Open("pgx", connUrl)
	if err != nil {
		log.Fatal(err)
	}
	s := LinksService{}
	s.redirectServiceUrl = redirectUrl
	s.conn = db.New(conn)

	return &s
}

func (s *LinksService) CreateShortLink(data models.CreateLinkData) (*models.CreateLinkResponse, error) {
	_, err := url.Parse(data.InitialLink)
	if err != nil || len(data.InitialLink) == 0 {
		return nil, models.ErrWrongInitialLink
	}
	if len(data.CustomLink) == 0 {
		return s.createRandomLink(data)
	}
	return s.createCustomLink(data)
}

func (s *LinksService) createRandomLink(data models.CreateLinkData) (*models.CreateLinkResponse, error) {
	retries := 0
	startingLink := data.InitialLink
	for {
		if retries > 10 {
			return nil, models.ErrCreationFailed
		}
		shortLink, err := linkgen.GenerateShortLink(data.InitialLink)
		if err != nil {
			return nil, err
		}
		url, err := s.conn.CreateLink(context.Background(), db.CreateLinkParams{
			ID:        shortLink,
			RefersTo:  startingLink,
			CreatedBy: data.UserId,
		})

		if err == nil {
			data := models.CreateLinkResponse{ShortLink: url}
			return &data, nil
		} else if !strings.Contains(err.Error(), "SQLSTATE 23505") {
			return nil, err
		}
		data.InitialLink += data.InitialLink
		retries++
	}
}

func (s *LinksService) createCustomLink(data models.CreateLinkData) (*models.CreateLinkResponse, error) {
	if !(linkverify.VerifyLink(data.CustomLink)) {
		return nil, models.ErrBadCustomLink
	}

	url, err := s.conn.CreateLink(context.Background(), db.CreateLinkParams{
		ID:        data.CustomLink,
		CreatedBy: data.UserId,
		RefersTo:  data.InitialLink,
	})

	if err == nil {
		data := models.CreateLinkResponse{ShortLink: url}
		return &data, nil
	} else if strings.Contains(err.Error(), "(SQLSTATE 23505)") {
		return nil, models.ErrAlreadyOccupied
	} else {
		return nil, err
	}
}
