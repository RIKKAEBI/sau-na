package controller

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func Home(c echo.Context) error {
	data := map[string]string{
		"name": "sau-na",
	}

	return c.Render(http.StatusOK, "home", data)
}
