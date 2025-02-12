package handler

import (
	"github.com/idkwhyureadthis/url-shortener/internal/service"
	"github.com/labstack/echo/v4"
)

type Handler struct {
	echo    *echo.Echo
	service service.Service
}

func New(connUrl string) *Handler {
	e := echo.New()
	setupHandlers(e)
	h := Handler{}
	h.service = service.New(connUrl)
	h.echo = e
	return &h
}

func (h *Handler) Start(addr string) error {
	return h.echo.Start(addr)
}

func setupHandlers(e *echo.Echo) {
}
