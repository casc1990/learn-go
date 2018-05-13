package controllers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
	"fmt"
)

type LoginController struct {
	beego.Controller
}

func (this *LoginController) Get() {
	//判断是否是退出请求(退出发送的url：/login?exit=true)
	isExit := this.Input().Get("exit")
	//如果退出，清除cookie
	if isExit == "true" {
		this.Ctx.SetCookie("uname","",-1,"/")
		this.Ctx.SetCookie("pwd","",-1,"/")

		//重定向到首页
		this.Redirect("/",302)
		return
	}
	//渲染模板
	this.TplName = "login.html"
}

func (this *LoginController) Post() {
	//.input方法里保存了用户的输入信息
	fmt.Println(this.Input())
	//Input().Get获取request的某个属性
	uname := this.Input().Get("uname")
	pwd := this.Input().Get("pwd")
	authlogin := this.Input().Get("authlogin") == "on"

	//验证用户名、密码
	if beego.AppConfig.String("uname") == uname && beego.AppConfig.String("pwd") ==pwd {
		//cookie保存时间设置(默认是0，浏览器关闭，cookie就失效)
		maxAge := 0

		//如果是自动登陆，cookie永久保存
		if authlogin {
			maxAge = 1 << 31 -1
		}
		//设置cookie保存时间
		this.Ctx.SetCookie("uname",uname,maxAge,"/")
		this.Ctx.SetCookie("pwd",pwd,maxAge,"/")

	}

	//登陆完成，重定向到首页
	this.Redirect("/",302)
	return
}

//验证上下文（context.Context用来解析用户输入或者输出的信息，是this.Input()或者this.Output()的封装）
func checkAccount(ctx *context.Context) bool {
	//获取用户名和密码的cookie
	ch,err := ctx.Request.Cookie("uname")
	if err !=nil {
		return false
	}
	uname := ch.Value  //获取用户名的cookie的值

	ch,err = ctx.Request.Cookie("pwd")
	if err !=nil {
		return false
	}
	pwd := ch.Value  //获取密码的cookie的值

	//返回验证结果
	return beego.AppConfig.String("uname") == uname &&
			beego.AppConfig.String("pwd") ==pwd
}