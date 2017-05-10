package controllersGoogle

import (
	"context"
	"errors"
	"example-golang-oauth2/lib/google"

	"github.com/astaxie/beego"
	v2 "google.golang.org/api/oauth2/v2"
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

	config := google.GetConnect()

	context := context.Background()

	tok, err := config.Exchange(context, request.Code)
	if err != nil {
		panic(err)
	}

	if tok.Valid() == false {
		panic(errors.New("vaild token"))
	}

	service, _ := v2.New(config.Client(context, tok))
	tokenInfo, _ := service.Tokeninfo().AccessToken(tok.AccessToken).Context(context).Do()

	c.Data["ID"] = tokenInfo.UserId
	c.Data["Email"] = tokenInfo.Email
	c.TplName = "google/callback.tpl"
}
