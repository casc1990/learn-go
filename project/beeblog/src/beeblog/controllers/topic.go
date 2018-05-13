package controllers

import (
	"github.com/astaxie/beego"
	"beeblog/models"
)

type TopicController struct {
	beego.Controller
}

func (this *TopicController) Get() {
	this.TplName = "topic.html"
	//访问文章时，文章按钮高亮（传递变量给模板，模板判断传递的值，决定是否高亮）
	this.Data["IsTopic"] = true
	//调用验证cookie函数验证，验证是否是登陆状态，并将值传递给模板，决定显示登陆还是退出
	this.Data["IsLogin"] = checkAccount(this.Ctx)
	//显示所有文章
	topics,err := models.GetAllTopics(false)
	if err != nil {
		beego.Error(err)
	}else {
		this.Data["Topics"] = topics
	}

}

func (this *TopicController) Post() {
	//检查用户cookie，判断用户是否是登录状态
	if !checkAccount(this.Ctx) {
		this.Redirect("/login",302)
		return
	}

	//检查通过，获取信息
	title := this.Input().Get("title")
	content := this.Input().Get("content")
	//入库
	var err error
	err = models.AddTopic(title,content)
	if err != nil {
		beego.Error(err)
	}

	//跳转
	this.Redirect("/topic",302)
	return

}

func (this *TopicController) Add() {
	this.TplName = "topic_add.html"
	//this.Ctx.WriteString("add")
}
