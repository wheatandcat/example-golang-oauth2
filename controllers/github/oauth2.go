package controllersGithub

import (
	"example-golang-oauth2/lib/github"

	"github.com/astaxie/beego"
)

// Oauth2Controller Oauth2コントローラー
type Oauth2Controller struct {
	beego.Controller
}

// Get 認証する
func (c *Oauth2Controller) Get() {
	config := github.GetConnect()

	url := config.AuthCodeURL("")

	c.Redirect(url, 302)
}
