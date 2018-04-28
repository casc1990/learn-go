package main

//mux设置路由

import (
	"net/http"
	"io"
	_"fmt"
	"log"
)
//定义结构体，如果要注册到mux.Handle中，必须实现ServeHTTP方法
type MyHandle struct {}
func (p *MyHandle) ServeHTTP(w http.ResponseWriter,r *http.Request) {
	io.WriteString(w,"url:"+ r.URL.String())
}

//定义handle函数
func sayhello(w http.ResponseWriter,r *http.Request) {
	io.WriteString(w,"this is version 2!")

}

func main() {
	//mux实现handle进行路由注册
	mux := http.NewServeMux()
	mux.Handle("/",&MyHandle{})  //注册定义的结构体
	mux.HandleFunc("/hello",sayhello) //注册函数
	//将自己定义的mux类型的handle注册到ListenAndServe
	err := http.ListenAndServe(":8080",mux)
	if err != nil {
		log.Fatal(err)
	}

}