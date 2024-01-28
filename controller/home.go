package controller

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func Home(c echo.Context) error {
	return c.HTML(http.StatusOK, "<script src='https://unpkg.com/htmx.org@1.9.10' integrity='sha384-D1Kt99CQMDuVetoL1lrYwg5t+9QdHe7NLX/SoJYkXDFfX37iInKRy5xLSi8nO7UC' crossorigin='anonymous'></script><link rel='stylesheet' crossorigin href='/assets/style.css'><script type='module' crossorigin src='/assets/main.js'></script><sau-button hx-post='/api/users/me'>hello world!</sau-button><br><sau-button outline hx-post='/api/users/me'>クリックして！！</sau-button>")
}
