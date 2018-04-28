package main

import "github.com/astaxie/beego"


type HomeController struct {
	beego.Controller
}
func (p HomeController) Get() {
	p.Ctx.WriteString("hello!world")
}

func main() {
	beego.Router("/",&HomeController{})
	beego.Run()
}
