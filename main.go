package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	// controller.InitController(e, db)
	e.GET("/", func(c echo.Context) error {
		return c.JSON(http.StatusOK, "OK")
	})
}
