package main

import (
	"fmt"
	"net/http"
	"strings"
	"log"
)

//http://localhost:9090/?name=abc&age=222
func sayhelloName(w http.ResponseWriter, r *http.Request) {
	r. ParseForm()  //解析参数，默认是不会解析的
	fmt.Println("---------")
	fmt.Println("request args:",r.Form)  //请求参数，map类型；request args: map[name:["abc"] age:[222]]
	fmt.Println("Path:",r.URL.Path)  //打印请求路径；Path: /
	fmt.Println("scheme:",r.URL.Scheme)  //打印请求协议(http、https,ftp)
	fmt.Println("request args:",r.Form["name"]) //读取name参数的值；request args: [abc]
	for k,v := range r.Form {  //循环请求字段，取每个字段的key和value
		fmt.Println("key:",k)
		fmt.Println("val:",strings.Join(v,""))
	}
	fmt.Fprintf(w,"hello golang!")  //这个写入到w的是输出到客户端的

}

func main() {
	http.HandleFunc("/",sayhelloName)  //定义访问的路由
	err := http.ListenAndServe(":9090",nil)  //设置监听的端口
	if err != nil {   //如果监听异常，记录日志..
		log.Fatal("ListenAndServer:",err)
	}
}