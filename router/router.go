package router

import (
	"sau-na/common"
	"sau-na/controller"
	"sau-na/middleware"

	"github.com/labstack/echo/v4"
)

type (
	Host struct {
		Echo *echo.Echo
	}
)

func Router() {
	// サイトのオリジンを取得
	origin, URL := common.LoadEnv()

	// Hosts
	hosts := map[string]*Host{}

	// API
	api := echo.New()
	hosts["api."+origin] = &Host{api}
	controller.ApiIndex(api)

	// Storybook
	sb := middleware.SpaBinding("./components/storybook-static")
	hosts["storybook."+origin] = &Host{sb}

	// Website
	site := middleware.SpaBinding("./components/dist")
	controller.SiteAuth(site, URL)
	hosts[origin] = &Host{site}

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
