package main

import (
	"fmt"
	"os"
)

func main() {
	//os.Getenv获取环境变量
	goos := os.Getenv("GOOS")
	gopath := os.Getenv("GOPATH")
	fmt.Println("os type:",goos)  //os type: windows
	fmt.Println("gopath:",gopath) //gopath: D:\go-project

	//路径相关
	dir,_ := os.Getwd()
	fmt.Println("当前路径:",dir) //当前路径
	fmt.Println(os.Chdir("..")) //切换目录。到上一级



}
