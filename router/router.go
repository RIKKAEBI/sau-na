package router

import (
	"sau-na/controller"
	"sau-na/middleware"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func InitRouter(e *echo.Echo, db *gorm.DB) {
	middleware.InitRenderer(e)

	healthController := controller.NewHealthController(db)
	e.GET("/hc", healthController.HC)
	e.GET("/hello", healthController.HelloPage)
}
