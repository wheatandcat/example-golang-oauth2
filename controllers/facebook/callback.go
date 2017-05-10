package controllersFacebook

import (
	"errors"
	"example-golang-oauth2/lib/facebook"

	"github.com/astaxie/beego"
	fb "github.com/huandu/facebook"
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

	config := facebook.GetConnect()

	tok, err := config.Exchange(oauth2.NoContext, request.Code)
	if err != nil {
		panic(err)
	}

	if tok.Valid() == false {
		panic(errors.New("vaild token"))
	}

	client := config.Client(oauth2.NoContext, tok)
	session := &fb.Session{
		Version:    "v2.8",
		HttpClient: client,
	}

	res, err := session.Get("/me?fields=id,name,email", nil)
	if err != nil {
		panic(err)
	}

	c.Data["Email"] = res["email"]
	c.TplName = "facebook/callback.tpl"
}
