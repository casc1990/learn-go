package main

import (
	"net/http"
	"io"
	"log"
)
//表单
const form  = `
<html lang="en">
	<head>
		<meta charset="UTF-8">
		<title>form</title>
	</head>
	<body>
		<form action="#" method="post" name="bar">
			username: <input type="text" name="username" />
			password: <input type="text" name="password" />
			<input type="submit" value="Submit" />

		</form>
	</body>
</html>`

//处理函数
func SimpleServer(w http.ResponseWriter,r *http.Request) {
	io.WriteString(w,"<h1>hello,world! </h1>")
}

//表单处理函数
func FormServer(w http.ResponseWriter,r *http.Request) {
	w.Header().Set("Content-Type","text/html")
	//判断请求类型
	switch r.Method {
	case "GET":
		io.WriteString(w,form)  //get请求，将表单显示给客户端
	case "POST":
		r.ParseForm()  //解析表单
		io.WriteString(w,r.Form["username"][0]) //获取字段
		io.WriteString(w,"\n")
		io.WriteString(w,r.FormValue("password"))

	}
}

//定义异常
func logPanics(handle http.HandlerFunc) http.HandlerFunc {
	return func(write http.ResponseWriter,request *http.Request) {
		defer func() {
			if x := recover();x != nil {
				log.Panicf("[%v] caught panic: %v",request.RemoteAddr,x)
			}
		}()
		//处理请求
		handle(write,request)
	}
}

func main() {
	http.HandleFunc("/test1",logPanics(SimpleServer))
	http.HandleFunc("/test2",logPanics(FormServer))
	if err := http.ListenAndServe(":8080",nil); err != nil {

	}

	}