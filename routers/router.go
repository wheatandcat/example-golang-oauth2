package routers

import (
	"example-golang-oauth2/controllers"
	"example-golang-oauth2/controllers/google"

	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/google/oauth2", &controllersGoogle.Oauth2Controller{})
	beego.Router("/google/callback", &controllersGoogle.CallbackController{})
	beego.Router("/", &controllers.MainController{})
}
