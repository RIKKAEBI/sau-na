package controller

import (
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
)

type User struct {
	Date int64
}

func Health(c echo.Context) error {
	user := User{Date: time.Now().Unix()}

	return c.JSON(http.StatusOK, user)
}
