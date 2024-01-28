package router

import (
	"sau-na/controller"

	"github.com/labstack/echo/v4"
)

func Init(e *echo.Echo) {
	// 静的ファイル
	e.File("/", "./components/dist/index.html")
	e.File("/favicon.ico", "./components/dist/favicon.ico")
	e.File("/assets/main.js", "./components/dist/assets/main.js")
	e.File("/assets/style.css", "./components/dist/assets/style.css")

	// API
	e.GET("/api/health", controller.Health)
}
