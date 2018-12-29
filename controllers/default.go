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
func (c *LoginController) Post() {
	c.Data["username"] = c.Ctx.Request.PostForm.Get("username")
	c.Data["password"] = c.Ctx.Request.PostForm.Get("password")
	c.TplName = "tmp.tpl"
}

//func (this *LoginController) Get() {
//	v := this.GetSession("asta")
//	if v == nil {
//		this.SetSession("asta", int(1))
//		this.Data["num"] = 0
//	} else {
//		this.SetSession("asta", v.(int)+1)
//		this.Data["num"] = v.(int)
//	}
//	this.TplName = "index.tpl"
//}

type SignController struct {
	beego.Controller
}

func (c *SignController) Get() {
	c.TplName = "sign.tpl"
}
func (c *SignController) Post() {
	c.Data["username"] = c.Ctx.Request.PostForm.Get("username")
	c.Data["password"] = c.Ctx.Request.PostForm.Get("password")
	c.Data["Massage"] = ""
	conn := models.Pool.Get()
	defer conn.Close()
	reply, err := conn.Do("SETNX", c.Data["username"], c.Data["password"])
	if err != nil {
		fmt.Println("conn err", err)
	} else {
		if reply.(int64) == 1 {
			c.Data["Massage"] = "sign success"
		} else {
			c.Data["Massage"] = "username is exist"
		}
	}
	c.TplName = "sign.tpl"
}
