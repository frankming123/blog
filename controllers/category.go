package controllers

import (
	"beeblog/models"

	"github.com/astaxie/beego"
)

type CategoryController struct {
	beego.Controller
}

func (c *CategoryController) Get() {
	isLogin := checkAccount(c.Ctx)
	c.Data["IsCategory"] = true
	var err error
	c.Data["Categories"], err = models.GetAllCategories()
	c.Data["IsLogin"] = isLogin
	c.TplName = "category.html"

	if err != nil {
		beego.Error(err)
	}
	if !isLogin {
		return
	}

	op := c.Input().Get("op")
	switch op {
	case "add":
		name := c.Input().Get("name")
		if len(name) == 0 {
			break
		}
		err := models.AddCategory(name)
		if err != nil {
			beego.Error(err)
		}
		c.Redirect("/category", 302)
		return
	case "del":
		id := c.Input().Get("id")
		if len(id) == 0 {
			break
		}
		err := models.DelCategories(id)
		if err != nil {
			beego.Error(err)
		}
		c.Redirect("/category", 302)
		return
	}
}
