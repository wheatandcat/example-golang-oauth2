package controllersGoogle

import (
	"example-golang-oauth2/lib/google"
	"strconv"

	"github.com/astaxie/beego"
)

// Oauth2Controller Oauth2コントローラー
type Oauth2Controller struct {
	beego.Controller
}

// Get 認証する
func (c *Oauth2Controller) Get() {
	config := google.GetConnect()

	url := config.AuthCodeURL(strconv.Itoa(1))

	c.Redirect(url, 302)
}
