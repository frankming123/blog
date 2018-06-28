package controllers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
)

type LoginController struct {
	beego.Controller
}

func (c *LoginController) Get() {
	isExit := c.Input().Get("exit") == "true"
	if isExit {
		c.Ctx.SetCookie("username", "", -1, "/")
		c.Ctx.SetCookie("password", "", -1, "/")
		c.Redirect("/", 302)
		return
	}
	c.Data["IsLoginPage"]=true
	c.TplName = "login.html"
}

func (c *LoginController) Post() {
	username := c.Input().Get("username")
	password := c.Input().Get("password")
	autoLogin := c.Input().Get("autologin") == "on"
	if beego.AppConfig.String("username") == username && beego.AppConfig.String("password") == password {
		maxAge := 0
		if autoLogin {
			maxAge = 1<<31 - 1
		}
		c.Ctx.SetCookie("username", username, maxAge, "/")
		c.Ctx.SetCookie("password", password, maxAge, "/")
	}
	c.Redirect("/", 302)
	return
}

func checkAccount(ctx *context.Context) bool {
	ck, err := ctx.Request.Cookie("username")
	if err != nil {
		return false
	}
	username := ck.Value

	ck, err = ctx.Request.Cookie("password")
	if err != nil {
		return false
	}
	password := ck.Value

	return beego.AppConfig.String("username") == username &&
		beego.AppConfig.String("password") == password
}
