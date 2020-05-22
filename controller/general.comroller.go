package controller

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func HomePage(e echo.Context) error {
	type HomePage struct {
		Message string `json:"message"`
	}
	return e.JSON(http.StatusOK, HomePage{Message: "API SERVER IS WORK!"})
}

func TestAPI(e echo.Context) error {
	return e.String(http.StatusOK, "Test API")
}
