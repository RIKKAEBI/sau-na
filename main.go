package main

import (
	"os"
	"sau-na/database"
	"sau-na/router"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	db, err := database.ConnectDatabase()
	if err != nil {
		os.Exit(1)
	}
	router.InitRouter(e, db)

	if err := e.Start(":3000"); err != nil {
	}
}
