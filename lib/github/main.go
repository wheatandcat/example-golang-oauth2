package github

import (
	"github.com/astaxie/beego"
	"golang.org/x/oauth2"
	githuboauth "golang.org/x/oauth2/github"
)

// GetConnect 接続を取得する
func GetConnect() *oauth2.Config {
	config := &oauth2.Config{
		ClientID:     beego.AppConfig.String("githubClientID"),
		ClientSecret: beego.AppConfig.String("githubClientSecret"),
		Endpoint:     githuboauth.Endpoint,
		Scopes:       []string{"user:email", "repo"},
		RedirectURL:  "http://localhost:8080/github/callback",
	}

	return config
}
