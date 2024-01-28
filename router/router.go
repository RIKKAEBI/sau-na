package router

import (
	"sau-na/controller"

	"github.com/labstack/echo/v4"
)

func InitRouter(e *echo.Echo) {
	e.File("/assets/main.js", "./components/dist/assets/main.js")
	e.File("/assets/style.css", "./components/dist/assets/style.css")

	e.GET("/", controller.Home)
}
