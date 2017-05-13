package controllersTwitter

import (
	"example-golang-oauth2/lib/twitter"

	"github.com/astaxie/beego"
	"github.com/garyburd/go-oauth/oauth"
)

// TweetController Tweetコントローラ
type TweetController struct {
	beego.Controller
}

// Get ツイートする
func (c *TweetController) Get() {
	c.StartSession()

	at := oauth.Credentials{
		Secret: c.CruSession.Get("oauth_secret").(string),
		Token:  c.CruSession.Get("oauth_token").(string),
	}

	if err := twitter.PostTweet(&at); err != nil {
		panic(err)
	}

	c.Redirect("http://localhost:8080/?message=投稿しました", 302)
}
