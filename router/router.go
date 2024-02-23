package router

import (
	"sau-na/controller"
	"sau-na/global"
	"sau-na/middleware"

	"github.com/labstack/echo/v4"
)

type (
	Host struct {
		Echo *echo.Echo
	}
)

func Router() {
	// Hosts
	hosts := map[string]*Host{}

	// API
	api := echo.New()
	api.Use(middleware.Cors())
	hosts["api."+global.Env.DOMAIN] = &Host{api}
	// routing
	controller.ApiIndex(api)

	// Storybook
	sb := echo.New()
	sb.Use(middleware.Spa("./components/storybook-static"))
	hosts["storybook."+global.Env.DOMAIN] = &Host{sb}

	// Website
	site := echo.New()
	site.Use(middleware.Spa("./components/dist"))
	hosts[global.Env.DOMAIN] = &Host{site}
	// routing
	controller.SiteAuth(site)

	// Database
	db := echo.New()
	hosts["database."+global.Env.DOMAIN] = &Host{db}

	// Server
	e := echo.New()
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
