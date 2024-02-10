package controller

import (
	"context"
	"encoding/json"
	"net/http"
	"sau-na/middleware"

	"github.com/labstack/echo/v4"
	"golang.org/x/oauth2"
)

func SiteAuth(e *echo.Echo, URL string) {
	register(e, URL)
}

func register(e *echo.Echo, URL string) {
	e.GET("/register", func(c echo.Context) error {
		code := c.QueryParams().Get("code")

		// codeのクエリがついている場合はユーザー情報を取得しにいく
		if code != "" {
			// codeからtokenを取得しにいく
			googleOauthConfig := middleware.NewGoogleOauthConfig(URL + "/register")
			token, err := googleOauthConfig.Exchange(c.Request().Context(), code)

			if err != nil {
				return c.NoContent(404)
			}

			// tokenの有効性を確認
			client := googleOauthConfig.Client(context.Background(), token)
			response, err := client.Get("https://www.googleapis.com/oauth2/v2/userinfo")
			if err != nil {
				return c.NoContent(404)
			}

			// jsonに変換
			userInfo := make(map[string]interface{})
			err = json.NewDecoder(response.Body).Decode(&userInfo)
			if err != nil {
				return c.NoContent(404)
			}

			// cookieにセット
			c.SetCookie(&http.Cookie{
				Name:   "cookie",
				Value:  userInfo["name"].(string),
				MaxAge: 86400,
				Path:   "/",
			})

			return c.Redirect(http.StatusFound, URL+"/mypage")
		}

		// Googleログインページにリダイレクトさせる
		googleOauthConfig := middleware.NewGoogleOauthConfig(URL + "/register")
		url := googleOauthConfig.AuthCodeURL("state", oauth2.AccessTypeOffline)

		return c.Redirect(http.StatusFound, url)
	})
}
