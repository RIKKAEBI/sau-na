package conf

import (
	"sau-na/common"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

// Googleのプロジェクト設定
func GoogleAuthConf(redirectUrl string) *oauth2.Config {
	conf := &oauth2.Config{
		RedirectURL:  redirectUrl,
		ClientID:     common.Env.GOOGLE_CLIENT_ID,
		ClientSecret: common.Env.GOOGLE_CLIENT_SECRET,
		Scopes:       []string{"https://www.googleapis.com/auth/userinfo.email", "https://www.googleapis.com/auth/userinfo.profile"},
		Endpoint:     google.Endpoint,
	}

	return conf
}