package routers

import (
	"beeblog/controllers"
	"github.com/astaxie/beego"
)

func init() {
	//首页路由
    beego.Router("/", &controllers.HomeController{})
    //登陆路由
    beego.Router("/login",&controllers.LoginController{})
    //分类路由
    beego.Router("/category",&controllers.CategoryController{})
    //文章路由
    beego.Router("/topic",&controllers.TopicController{})
    //文章自动路由
    beego.AutoRouter(&controllers.TopicController{})

}
