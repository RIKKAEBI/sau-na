package main

import (
	"sau-na/middleware"
	"sau-na/router"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	router.InitRouter(e)
	middleware.InitRenderer(e)

	e.Logger.Fatal(e.Start(":3000"))
}
