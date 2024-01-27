package controller

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type HealthController interface {
	HC(c echo.Context) error
	HelloPage(c echo.Context) error
}

type healthController struct {
	db *gorm.DB
}

func NewHealthController(db *gorm.DB) *healthController {
	return &healthController{
		db: db,
	}
}

func (h *healthController) HC(c echo.Context) error {
	return c.JSON(http.StatusOK, "OK")
}

func (a *healthController) HelloPage(c echo.Context) error {
	return c.Render(http.StatusOK, "index.html", nil)
}
