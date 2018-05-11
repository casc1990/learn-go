package main

import (
	_ "sre-op/routers"
	"github.com/astaxie/beego"
)

func main() {
	beego.SetStaticPath("/static", "static")

	beego.Run()
}

