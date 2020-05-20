package controller

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func HomePage(c echo.Context) error {
	type HomePage struct {
		Message string `json:"message"`
	}
	return c.JSON(http.StatusOK, HomePage{Message: "API SERVER IS WORK!"})
}
