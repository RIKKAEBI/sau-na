package middleware

import (
	"sau-na/common"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

// Googleのプロジェクト設定
func NewGoogleOauthConfig(redirectUrl string) *oauth2.Config {

	conf := &oauth2.Config{
		RedirectURL:  redirectUrl,
		ClientID:     common.GOOGLE_CLIENT_ID,
		ClientSecret: common.GOOGLE_CLIENT_SECRET,
		Scopes:       []string{"https://www.googleapis.com/auth/userinfo.email", "https://www.googleapis.com/auth/userinfo.profile"},
		Endpoint:     google.Endpoint,
	}

	return conf
}
