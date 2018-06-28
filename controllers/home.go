package controllers

import (
	"beeblog/models"

	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	c.Data["IsHome"] = true
	c.Data["IsLogin"] = checkAccount(c.Ctx)
	c.TplName = "home.html"

	topics, err := models.GetAllTopics(true, c.Input().Get("cate"))
	if err != nil {
		beego.Error(err)
	}
	c.Data["Topics"] = topics

	categories, err := models.GetAllCategories()
	if err != nil {
		beego.Error(err)
	}
	c.Data["Categories"] = categories
}
