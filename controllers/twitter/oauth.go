package controllersTwitter

import (
	"example-golang-oauth2/lib/twitter"
	"log"

	"github.com/astaxie/beego"
)

// Oauth2Controller Oauth2コントローラー
type Oauth2Controller struct {
	beego.Controller
}

// Get 認証する
func (c *Oauth2Controller) Get() {
	config := twitter.GetConnect()
	rt, err := config.RequestTemporaryCredentials(nil, "http://localhost:8080/twitter/callback", nil)
	if err != nil {
		log.Println("----------------------")
		panic(err)
	}

	log.Println(rt)

	url := config.AuthorizationURL(rt, nil)

	c.Redirect(url, 302)
}
