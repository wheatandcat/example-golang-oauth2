package controllersGithub

import (
	"example-golang-oauth2/lib/github"

	"github.com/astaxie/beego"
	gh "github.com/google/go-github/github"
	"golang.org/x/oauth2"
)

// CallbackController コールバックコントローラ
type CallbackController struct {
	beego.Controller
}

// CallbackRequest コールバックリクエスト
type CallbackRequest struct {
	Code  string `form:"code"`
	State int    `form:"state"`
}

// Get コールバックする
func (c *CallbackController) Get() {
	request := CallbackRequest{}
	if err := c.ParseForm(&request); err != nil {
		panic(err)
	}

	config := github.GetConnect()

	tok, err := config.Exchange(oauth2.NoContext, request.Code)
	if err != nil {
		panic(err)
	}

	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: tok.AccessToken},
	)
	tc := oauth2.NewClient(oauth2.NoContext, ts)

	client := gh.NewClient(tc)

	u, _, err := client.Users.Get(oauth2.NoContext, "")
	if err != nil {
		panic(err)
	}

	c.Data["ID"] = u.GetID()
	c.Data["AvatarURL"] = u.GetAvatarURL()
	c.Data["Name"] = u.GetName()
	c.Data["ReposURL"] = u.GetReposURL()

	c.TplName = "github/callback.tpl"
}
