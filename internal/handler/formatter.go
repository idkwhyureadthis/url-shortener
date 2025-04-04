package handler

import "github.com/labstack/echo/v4"

func jsonifyError(c echo.Context, err error, code int) error {
	return c.JSON(code, map[string]string{
		"message": err.Error(),
	})
}
