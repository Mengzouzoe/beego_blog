package controllers

import (
	"myapp/models"

	"github.com/astaxie/beego"
)

type HomeController struct {
	beego.Controller
}

func (this *HomeController) Get() {
	this.TplName = "home.html"
	this.Data["IsHome"] = true
	this.Data["IsLogin"] = checkAccount(this.Ctx)

	//Testing
	//beego.Alert("!!!!" + this.Input().Get("label"))
	topics, err := models.GetAllTopics(
		this.Input().Get("cate"), this.Input().Get("label"), true)
	if err != nil {
		beego.Error(err)
	}
	this.Data["Topics"] = topics

	categories, err := models.GetAllCategories()
	if err != nil {
		beego.Error(err)
	}
	this.Data["Categories"] = categories
}
