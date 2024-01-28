package main

import (
	"sau-na/router"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	router.Init(e)
	e.Logger.Fatal(e.Start(":3000"))
}
