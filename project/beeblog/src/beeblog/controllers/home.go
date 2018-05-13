package controllers

import (
	"github.com/astaxie/beego"
	"beeblog/models"
)

type HomeController struct {
	beego.Controller
}

func (this *HomeController) Get() {
	this.TplName = "index.html"
	//访问首页时，首页按钮高亮（传递变量给模板，模板判断传递的值，决定是否高亮）
	this.Data["IsHome"] = true
	//调用验证cookie函数验证，验证是否是登陆状态，并将值传递给模板，决定显示登陆还是退出
	this.Data["IsLogin"] = checkAccount(this.Ctx)

	//显示所有文章
	topics,err := models.GetAllTopics(true)
	if err != nil {
		beego.Error(err)
	}else {
		this.Data["Topics"] = topics
	}

}
