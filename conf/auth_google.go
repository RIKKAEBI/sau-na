package conf

import (
	"sau-na/global"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

var GoogleAuthScopeEmail = "https://www.googleapis.com/auth/userinfo.email"

var GoogleAuthScopeProfile = "https://www.googleapis.com/auth/userinfo.profile"

var GoogleOAuthScopeUserInfo = "https://www.googleapis.com/oauth2/v2/userinfo"

// Googleのプロジェクト設定
func GoogleAuthInit() *oauth2.Config {
	conf := &oauth2.Config{
		RedirectURL:  global.Env.URL + "/register",
		ClientID:     global.Env.GOOGLE_CLIENT_ID,
		ClientSecret: global.Env.GOOGLE_CLIENT_SECRET,
		Scopes:       []string{GoogleAuthScopeEmail, GoogleAuthScopeProfile},
		Endpoint:     google.Endpoint,
	}

	return conf
}
