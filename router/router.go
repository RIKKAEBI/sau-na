package router

import (
	"net/http"
	"sau-na/common"
	"sau-na/controller"
	sauna_middleware "sau-na/middleware"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type (
	Host struct {
		Echo *echo.Echo
	}
)

func Router() {
	// サイトのオリジンを取得
	ORIGIN := common.ORIGIN
	URL := common.URL

	// Hosts
	hosts := map[string]*Host{}

	// API
	api := echo.New()
	hosts["api."+ORIGIN] = &Host{api}
	controller.ApiIndex(api)

	// Storybook
	sb := sauna_middleware.SpaBinding("./components/storybook-static")
	hosts["storybook."+ORIGIN] = &Host{sb}

	// Website
	site := sauna_middleware.SpaBinding("./components/dist")
	controller.SiteAuth(site, URL)
	hosts[ORIGIN] = &Host{site}

	// Server
	e := echo.New()
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"https://sau-na.com", "http://localhost:4000", "http://localhost:3000"},
		AllowMethods: []string{http.MethodGet, http.MethodPut, http.MethodPost, http.MethodDelete},
	}))

	e.Any("/*", func(c echo.Context) (err error) {
		req := c.Request()
		res := c.Response()
		host := hosts[req.Host]

		if host == nil {
			err = echo.ErrNotFound
		} else {
			host.Echo.ServeHTTP(res, req)
		}

		return
	})
	e.Logger.Fatal(e.Start(":3000"))
}
