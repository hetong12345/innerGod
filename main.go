package main

import (
	"github.com/astaxie/beego"
	_ "github.com/hetong12345/innerGod/models"
	_ "github.com/hetong12345/innerGod/routers"
)

func main() {
	beego.Run()
}
