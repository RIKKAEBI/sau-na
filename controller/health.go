package controller

import (
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
)

type Date struct {
	date int64
}

func Health(c echo.Context) error {
	date := Date{date: time.Now().Unix()}

	return c.JSON(http.StatusOK, date)
}

// TODO: 一旦本番で動かないのでコメントアウト
// package controller

// import (
// 	"net/http"

// 	"github.com/labstack/echo/v4"
// 	"gorm.io/gorm"
// )

// type HealthController interface {
// 	HC(c echo.Context) error
// 	HelloPage(c echo.Context) error
// }

// type healthController struct {
// 	db *gorm.DB
// }

// func NewHealthController(db *gorm.DB) *healthController {
// 	return &healthController{
// 		db: db,
// 	}
// }

// func (h *healthController) HC(c echo.Context) error {
// 	return c.JSON(http.StatusOK, "OK")
// }
