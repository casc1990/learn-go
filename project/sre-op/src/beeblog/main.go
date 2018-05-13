package main

import (
	_ "beeblog/routers"
	"github.com/astaxie/beego"
	"beeblog/models"
	"github.com/astaxie/beego/orm"

)

func init() {
	//注册model
	models.RegisterDB()
}

func main() {
	//打印orm详细日志
	orm.Debug = true
	//model更新了，自动sync
	orm.RunSyncdb("default",false,true)
	beego.Run()
}

