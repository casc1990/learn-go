package main

import (
	 "fmt"
	"html/template"
	"net/http"
)

//定义一个模板变量
var myTemplate *template.Template

//定义一个结构体，用于将结构体的字段传递给表单
type Person struct {
	Name string
	Age int
	Score float32
	Title string
}

//初始化模板函数,（解析模板文件：template.ParseFiles）
func initTemplate(tep string) (err error) {
	myTemplate,err = template.ParseFiles(tep)
	if err != nil {
		fmt.Println("parse html file err:",err)
		return
	}
	return
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

//路由处理函数（接收http.ResponseWriter和http.Request参数即可）
func UserInfo(w http.ResponseWriter,r *http.Request) {
	//fmt.Println("handle hello")
	//fmt.Fprintf(w,"hello!")

	//实例化结构体，将结构体传递给表单渲染
	arr := []Person{}
	p := Person{Name:"stu01",Age:28,Score:89,Title:"我的个人网站"}
	p1 := Person{Name:"stu02",Age:29,Score:90,Title:"sut02的网站"}
	p3 := Person{Name:"stu03",Age:30,Score:91,Title:"sut03的网站"}
	arr = append(arr,p)
	arr = append(arr,p1)
	arr = append(arr,p3)
	myTemplate.Execute(w,arr) //渲染模板
}



func main() {

	initTemplate("D:/go-project/src/learn-go/example/day10/example5-template/main/index.html")
	//路由注册
	http.HandleFunc("/user/info",UserInfo)
	//监听并初始化
	err := http.ListenAndServe("0.0.0.0:8080",nil)
	if err != nil {
		fmt.Println("http listen failed",err)
	}

}