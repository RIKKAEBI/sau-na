package controller

import (
	"net/http"
	"sau-na/global"
	"sau-na/services"

	"github.com/labstack/echo/v4"
	"golang.org/x/oauth2"
)

func SiteAuth(e *echo.Echo) {
	e.GET("/register", func(c echo.Context) error {
		code := c.QueryParams().Get("code")

		// codeのクエリがついている場合はユーザー情報を取得しにいく
		if code != "" {
			_, err := services.Register(c.Request().Context(), code)

			if err != nil {
				return c.Redirect(http.StatusFound, global.Env.URL+"?error=auth")
			} else {
				return c.Redirect(http.StatusFound, global.Env.URL+"/mypage")
			}
		}

		// Googleログインページにリダイレクトさせる
		url := global.GoogleAuthConf.AuthCodeURL("state", oauth2.AccessTypeOffline)
		return c.Redirect(http.StatusFound, url)
	})
}
