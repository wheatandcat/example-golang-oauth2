package google

import (
	"github.com/astaxie/beego"
	"golang.org/x/oauth2"
)

const (
	authorizeEndpoint = "https://accounts.google.com/o/oauth2/v2/auth"
	tokenEndpoint     = "https://www.googleapis.com/oauth2/v4/token"
)

// GetConnect 接続を取得する
func GetConnect() *oauth2.Config {
	config := &oauth2.Config{
		ClientID:     beego.AppConfig.String("googleClientID"),
		ClientSecret: beego.AppConfig.String("googleClientSecret"),
		Endpoint: oauth2.Endpoint{
			AuthURL:  authorizeEndpoint,
			TokenURL: tokenEndpoint,
		},
		Scopes:      []string{"openid", "email", "profile"},
		RedirectURL: "http://localhost:8080/google/callback",
	}

	return config
}
