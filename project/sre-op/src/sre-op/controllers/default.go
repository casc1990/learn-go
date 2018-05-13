package controllers

import (
	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	//自定义数据，传递给模板
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	//获取config文件变量内容
	c.Ctx.WriteString("AppName:" + beego.AppConfig.String("appname"))
	//设置日志级别
	beego.SetLevel(beego.LevelInformational)

	//c.TplName = "index.html"
}
