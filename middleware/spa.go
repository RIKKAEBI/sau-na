package middleware

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

// SPAモードでファイルを配置
func Spa(path string) echo.MiddlewareFunc {
	return middleware.StaticWithConfig(middleware.StaticConfig{
		HTML5: true,
		Root:  path,
	})
}
