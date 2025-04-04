package handler

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"github.com/idkwhyureadthis/url-shortener/shortener/internal/models"
	"github.com/idkwhyureadthis/url-shortener/shortener/internal/service"
	"github.com/labstack/echo/v4"
)

type Handler struct {
	echo            *echo.Echo
	linksService    *service.LinksService
	redirectService *service.RedirectService
}

func New(connUrl, redirectUrl string) *Handler {
	e := echo.New()
	h := Handler{}
	h.linksService = service.NewLinksService(connUrl, redirectUrl)
	h.redirectService = service.NewRedirectService(connUrl)
	h.echo = e
	h.setupHandlers()
	return &h
}

func (h *Handler) Start(addr string) error {
	return h.echo.Start(":" + addr)
}

func (h *Handler) setupHandlers() {
	h.echo.Use()
	h.echo.GET("/available", h.checkAvailability)
	h.echo.POST("/new", h.createLink)
	h.echo.GET("/:link", h.getLink)
}

func (h *Handler) checkAvailability(c echo.Context) error {
	return nil
}

func (h *Handler) createLink(c echo.Context) error {
	data := models.CreateLinkData{}
	data.CustomLink = c.QueryParam("cl")
	data.InitialLink = c.QueryParam("il")
	idString := c.QueryParam("id")
	id, err := strconv.ParseInt(idString, 10, 64)
	if err != nil {
		return jsonifyError(c, err, http.StatusBadRequest)
	}
	data.UserId = id
	resp, err := h.linksService.CreateShortLink(data)
	if errors.Is(err, models.ErrWrongInitialLink) {
		return jsonifyError(c, err, http.StatusBadRequest)
	} else if errors.Is(err, models.ErrCreationFailed) || errors.Is(err, models.ErrAlreadyOccupied) {
		return jsonifyError(c, err, http.StatusConflict)
	} else if err != nil {
		return jsonifyError(c, err, http.StatusInternalServerError)
	}
	return c.JSON(http.StatusCreated, resp)
}

func (h *Handler) getLink(c echo.Context) error {
	id := c.Param("link")
	fmt.Println(id)
	data := h.redirectService.GetFullLink(id)
	if data.Err != nil {
		return jsonifyError(c, data.Err, data.Code)
	}
	fmt.Println(data.FullLink)
	return c.Redirect(http.StatusMovedPermanently, data.FullLink)
}
