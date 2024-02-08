package router

import (
	"net/http"
	"sau-na/common"

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
	origin := common.LoadEnv()

	// Hosts
	hosts := map[string]*Host{}

	// API
	api := echo.New()
	hosts["api."+origin] = &Host{api}
	api.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "API")
	})

	// Storybook
	sb := spaBinding("./components/storybook-static")
	hosts["storybook."+origin] = &Host{sb}

	// Website
	site := spaBinding("./components/dist")
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

// SPAモードでファイルを配置
func spaBinding(path string) *echo.Echo {
	e := echo.New()
	e.Use(middleware.StaticWithConfig(middleware.StaticConfig{
		HTML5: true,
		Root:  path,
	}))

	return e
}
