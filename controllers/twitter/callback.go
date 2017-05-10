package controllersTwitter

import (
	"example-golang-oauth2/lib/twitter"

	"github.com/astaxie/beego"
	"github.com/garyburd/go-oauth/oauth"
)

// CallbackController コールバックコントローラ
type CallbackController struct {
	beego.Controller
}

// CallbackRequest コールバックリクエスト
type CallbackRequest struct {
	Token    string `form:"oauth_token"`
	Verifier string `form:"oauth_verifier"`
}

// Get コールバックする
func (c *CallbackController) Get() {
	c.StartSession()

	request := CallbackRequest{}
	if err := c.ParseForm(&request); err != nil {
		panic(err)
	}

	at, err := twitter.GetAccessToken(
		&oauth.Credentials{
			Token:  c.CruSession.Get("request_token").(string),
			Secret: c.CruSession.Get("request_token_secret").(string),
		},
		request.Verifier,
	)
	if err != nil {
		panic(err)
	}

	account := twitter.Account{}
	if err = twitter.GetMe(at, &account); err != nil {
		panic(err)
	}

	c.Data["ID"] = account.ID
	c.Data["ProfileImageURL"] = account.ProfileImageURL
	c.Data["ScreenName"] = account.ScreenName
	c.Data["Email"] = account.Email
	c.TplName = "twitter/callback.tpl"
}
