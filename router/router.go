package router

import (
	"github.com/labstack/echo/v4"
)

func InitRouter(e *echo.Echo) {
	setStatic(e)
}

func setStatic(e *echo.Echo) {
	e.File("/", "./components/dist/index.html")
	e.File("/favicon.ico", "./components/dist/favicon.ico")
	e.File("/assets/main.js", "./components/dist/assets/main.js")
	e.File("/assets/style.css", "./components/dist/assets/style.css")
}
