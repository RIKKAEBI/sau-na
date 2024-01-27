package controller

import (
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func InitController(e *echo.Echo, db *gorm.DB) {
	healthController := NewHealthController(db)
	e.GET("/hc", healthController.HC)
}
