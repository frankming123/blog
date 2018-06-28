package main

import (
	_ "beeblog/routers"
	"beeblog/models"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

func init(){
	models.RegisterDB()
}

func main() {
	orm.Debug=true
	orm.RunSyncdb("default",false,true)
	beego.Run()
}

