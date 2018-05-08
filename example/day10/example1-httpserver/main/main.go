package main

import (
	"net/http"
	"fmt"
)

//路由处理函数（接收http.ResponseWriter和http.Request参数即可）
func Hello(w http.ResponseWriter,r *http.Request) {
	fmt.Println("handle hello")
	fmt.Fprintf(w,"hello!")
}

//路由处理函数
func login(w http.ResponseWriter,r *http.Request) {
	fmt.Println("handle login")
	fmt.Fprintf(w,"login ")
}


func main() {
	//路由注册
	http.HandleFunc("/",Hello)
	http.HandleFunc("/user/login",login)
	//监听并初始化
	err := http.ListenAndServe("0.0.0.0:8080",nil)
	if err != nil {
		fmt.Println("http listen failed")
	}

}
