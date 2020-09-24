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

	topics, err := models.GetAllTopics(true)
	if err != nil {
		beego.Error(err)
	}
	this.Data["Topics"] = topics
}
