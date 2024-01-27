package router

import (
	"sau-na/controller"

	"github.com/labstack/echo/v4"
)

func InitRouter(e *echo.Echo) {
	// middleware.InitRenderer(e)

	e.GET("/", controller.Home)
}
