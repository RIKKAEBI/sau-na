package controller

import (
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
)

func ApiIndex(e *echo.Echo) {
	e.GET("/", indexUrl)
}

type IndexDate struct {
	Date int64 `json:"date"`
}

func indexUrl(c echo.Context) error {
	date := IndexDate{Date: time.Now().Unix()}

	return c.JSON(http.StatusOK, date)
}
