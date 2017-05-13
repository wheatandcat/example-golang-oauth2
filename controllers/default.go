package controllers

import (
	"github.com/astaxie/beego"
)

// MainController メインコントローラー
type MainController struct {
	beego.Controller
}

// Get メイン
func (c *MainController) Get() {
	c.Data["Message"] = c.GetString("message")
	c.TplName = "index.tpl"
}
