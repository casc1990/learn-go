package main

import (
	"io"
	"log"
	"net/http"
)
//Sayhello要作为参数传给HandleFunc函数，所以必须要接受ResponseWriter和Request参数
func Sayhello(w http.ResponseWriter,r *http.Request) {
	io.WriteString(w,"this is version 1")
}


func main() {
	//设置路由
	http.HandleFunc("/",Sayhello)
	//运行
	err := http.ListenAndServe(":8080",nil)
	if err != nil {
		log.Fatal(err)
	}
}
