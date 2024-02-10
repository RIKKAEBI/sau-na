package middleware

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

// SPAモードでファイルを配置
func SpaBinding(path string) *echo.Echo {
	e := echo.New()
	e.Use(middleware.StaticWithConfig(middleware.StaticConfig{
		HTML5: true,
		Root:  path,
	}))

	return e
}
