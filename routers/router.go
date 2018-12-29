package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
	"github.com/hetong12345/innerGod/controllers"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/login", &controllers.LoginController{})
	beego.Router("/sign", &controllers.SignController{})
	beego.Post("/tmp", func(ctx *context.Context) {

	})
}
