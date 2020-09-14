package main

import (
	"myapp/controllers"
	"myapp/models"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

func init() {
	models.RegisterDB()
}

func main() {
	orm.Debug = true
	orm.RunSyncdb("default", false, true)

	beego.Router("/", &controllers.HomeController{})
	beego.Router("/login", &controllers.LoginController{})
	beego.Run()
}
