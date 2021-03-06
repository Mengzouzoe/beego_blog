package main

import (
	"myapp/controllers"
	"myapp/models"
	"os"

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
	beego.Router("/category", &controllers.CategoryController{})
	beego.Router("/topic", &controllers.TopicController{})
	beego.AutoRouter(&controllers.TopicController{})
	beego.Router("/reply", &controllers.ReplyController{})
	beego.Router("/reply/add", &controllers.ReplyController{}, "post:Add")
	beego.Router("/reply/delete", &controllers.ReplyController{}, "get:Delete")

	//创建附件目录
	os.Mkdir("attachment", os.ModePerm)
	//作为静态文件
	//beego.SetStaticPath("/attachment", attachment)

	//作为一个单独的控制器来处理
	beego.Router("/attachment/:all", &controllers.AttachController{})

	beego.Run()
}
