package main

import (
	"net/http"
	"io"
)

func main() {
	http.HandleFunc("/",Cookie)
	http.HandleFunc("/2",Cookie2)
	http.ListenAndServe(":8080",nil)
}

//获取cookie方法一（问题：cookie设置到response了，request里第一次拿不到）
func Cookie(w http.ResponseWriter,r *http.Request) {
	//实例化cookie对象
	ck := &http.Cookie{
		Name: "myCookie",
		Value: "myValue",
		Path: "/",
		Domain: "localhost",
		MaxAge: 1200,
	}
	//设置cookice
	http.SetCookie(w,ck)
	//查询cookie
	ck2,err := r.Cookie("myCookie")
	if err != nil {
		io.WriteString(w,err.Error())
	}else {
		io.WriteString(w,ck2.Value)
	}

}

//获取cookie方法二(使用response的Header设置，注意：Header不能识别空格，需要自己处理)
func Cookie2(w http.ResponseWriter,r *http.Request) {
	//实例化cookie对象
	ck := &http.Cookie{
		Name: "myCookie",
		Value: "myValue",
		Path: "/",
		Domain: "localhost",
		MaxAge: 1200,
	}
	//设置cookice
	w.Header().Set("Set-Cookie",ck.String())
	//查询cookie
	ck2,err := r.Cookie("myCookie")
	if err != nil {
		io.WriteString(w,err.Error())
	}else {
		io.WriteString(w,ck2.Value)
	}

}

