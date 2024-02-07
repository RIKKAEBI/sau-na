package router

import (
	"net/http"
	"os"
	"path/filepath"
	"sau-na/common"
	"strings"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type (
	Host struct {
		Echo *echo.Echo
	}
)

func Router() {
	origin := common.LoadEnv()

	// Hosts
	hosts := map[string]*Host{}

	// API
	api := echo.New()
	api.Use(middleware.Logger())
	api.Use(middleware.Recover())

	hosts["api."+origin] = &Host{api}

	api.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "API")
	})

	// Storybook
	storybook := echo.New()
	storybook.Use(middleware.Logger())
	storybook.Use(middleware.Recover())

	hosts["storybook."+origin] = &Host{storybook}
	storybookHandler(storybook)

	// Website
	site := echo.New()
	site.Use(middleware.Logger())
	site.Use(middleware.Recover())

	hosts[origin] = &Host{site}
	appHandler(site)

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

func appHandler(e *echo.Echo) {
	e.File("/", "./components/dist/index.html")
	e.File("/favicon.ico", "./components/dist/favicon.ico")
	e.File("/assets/main.js", "./components/dist/assets/main.js")
	e.File("/assets/style.css", "./components/dist/assets/style.css")

	// TODO: 頭悪い配置やめたい時間ないからとりあえず
	e.File("/assets/style2.css", "./components/dist/assets/style2.css")
	e.File("/assets/style3.css", "./components/dist/assets/style3.css")
}

// TODO: 要リファクタ
// アプリの部分もこれで置き換えれそう
func storybookHandler(e *echo.Echo) {
	err := filepath.Walk("./components/storybook-static", func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		dir := strings.Replace(path, "components/storybook-static", "", 1)

		if dir == "./" {
			return err
		}

		if dir == "/index.html" {
			e.File("/", path)
		} else {
			e.File(dir, path)
		}

		return nil
	})

	if err != nil {
		panic(err)
	}
}
