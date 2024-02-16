package middleware

import (
	"net/http"
	"sau-na/common"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

// cross origin resource secure の設定
func Cors() echo.MiddlewareFunc {
	API_URL := common.Env.PROTOCOL + "://" + common.Env.DOMAIN

	return middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{API_URL},
		AllowMethods: []string{http.MethodGet, http.MethodPut, http.MethodPost, http.MethodDelete},
	})
}
