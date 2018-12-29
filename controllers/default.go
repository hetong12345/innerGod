package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/hetong12345/innerGod/models"
	"os"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	c.TplName = "index.tpl"
}
func (c *MainController) Post() {
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	//fmt.Println(c.Ctx.Request.PostForm)
	c.Data["Fname"] = c.Ctx.Request.PostForm.Get("fname")
	c.TplName = "index.tpl"
}

type LoginController struct {
	beego.Controller
}

func (c *LoginController) Get() {
	conn := models.Pool.Get()
	reply, err := conn.Do("ping")
	if err != nil {
		fmt.Println("conn redis err", err)
		os.Exit(2)
	}
	fmt.Println(reply)
	c.TplName = "login.tpl"
}
