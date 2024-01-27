package main

import (
	"sau-na/router"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	router.InitRouter(e)

	if err := e.Start(":3000"); err != nil {

	}
}
