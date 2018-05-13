package controllers

import (
	"github.com/astaxie/beego"
	"beeblog/models"
	"fmt"
)

type CategoryController struct {
	beego.Controller
}

func (this *CategoryController) Get() {
	//获取操作类型
	op := this.Input().Get("op")
	fmt.Println("目前的操作是:",op)
	switch op {
	//如果是添加操作，获取用户提交的值，为空，退出
	case "add":
		name := this.Input().Get("name")  //name等于用户提交那个记录的值
		if len(name) == 0 {
			break
		}
		//数据入库
		err := models.AddCategory(name)
		if err != nil {
			fmt.Println("入库失败")
			beego.Error(err)
		}

		//添加完了，刷新页面
		this.Redirect("/category",301)
		return

	//如果是del，拿到id
	case "del":
		id := this.Input().Get("id")
		fmt.Println("要删除的ID：",id)
		if len(id) == 0 {
			break
		}
		err := models.DelCategory(id)
		if err != nil {
			beego.Error(err)
		}
		//删除完了，刷新页面
		this.Redirect("/category",301)
		return
	}
	
	
	this.TplName = "category.html"
	//开启高亮
	this.Data["IsCategory"] = true
	//调用验证cookie函数验证，验证是否是登陆状态，并将值传递给模板，决定显示登陆还是退出
	this.Data["IsLogin"] = checkAccount(this.Ctx)
	//获取全部的分类列表
	var err error
	this.Data["Categories"],err = models.GetAllCategory()
	if err != nil {
		beego.Error(err)
	}

}
