package middleware

import (
	"net/http"
	"sau-na/global"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

// cross origin resource secure の設定
func Cors() echo.MiddlewareFunc {
	return middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{global.Env.URL},
		AllowMethods: []string{http.MethodGet, http.MethodPut, http.MethodPost, http.MethodDelete},
	})
}
