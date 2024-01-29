package router

import (
	"sau-na/controller"

	"github.com/labstack/echo/v4"
)

func InitRouter(e *echo.Echo) {
	setStatic(e)

	// API用のエンドポイントは /api から始める
	e.GET("/api/health", controller.Health)
}

func setStatic(e *echo.Echo) {
	e.File("/", "./components/dist/index.html")
	e.File("/favicon.ico", "./components/dist/favicon.ico")
	e.File("/assets/main.js", "./components/dist/assets/main.js")
	e.File("/assets/style.css", "./components/dist/assets/style.css")

	// TODO: 頭悪い配置やめたい時間ないからとりあえず
	e.File("/assets/style2.css", "./components/dist/assets/style2.css")
	e.File("/assets/style3.css", "./components/dist/assets/style3.css")
}
