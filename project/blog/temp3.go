package main

import (
	"net/http"
	"time"
	"log"
	"io"
)

//自定义http.Server

var mux map[string]func(w http.ResponseWriter,r *http.Request)

//定义结构体，实现ServeHTTP方法
type myhandle struct {}
func (p *myhandle) ServeHTTP(w http.ResponseWriter,r *http.Request) {
	if h,ok :=mux[r.URL.String()]; ok {
		h(w,r)
		return
	}
	io.WriteString(w,"url:" + r.URL.String())
}

func sayHello(w http.ResponseWriter,r *http.Request) {
	io.WriteString(w, "sayHello! this is version 3!")
}

func sayBey(w http.ResponseWriter,r *http.Request) {
	io.WriteString(w, "sayBye! this is version 3!")
}

func main() {
	//自定义http.Server，Handler字段写我们自定义的handle
	server := http.Server{
		Addr: ":8080",
		Handler: &myhandle{},   //使用我们自定义的Handle
		ReadHeaderTimeout: 5 * time.Second,  //超时5s
	}

	mux = make(map[string]func(w http.ResponseWriter,r *http.Request))
	mux["/hello"] = sayHello
	mux["/bye"] = sayBey

	err := server.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}



