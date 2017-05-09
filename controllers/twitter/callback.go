package controllersTwitter

import (
	"github.com/astaxie/beego"
)

// CallbackController コールバックコントローラ
type CallbackController struct {
	beego.Controller
}

// CallbackRequest コールバックリクエスト
type CallbackRequest struct {
	Token  string `form:"code"`
	Secret int    `form:"state"`
}

// Get コールバックする
func (c *CallbackController) Get() {

	c.TplName = "twitter/callback.tpl"
}
