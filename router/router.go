package router

import (
	"sau-na/controller"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func InitRouter(e *echo.Echo, db *gorm.DB) {
	healthController := controller.NewHealthController(db)
	e.GET("/hc", healthController.HC)
}
