package main

import (
	"net/http"
	"fmt"
	"io/ioutil"
)

func main() {
	//http.Get调用get方法请求指定的url
	result,err := http.Get("http://www.baidu.com/")
	if err != nil {
		fmt.Println("get request err:",err)
		return
	}

	//获取请求内容(get请求返回的是Response对象，Body方法获取相应的内容)
	data,err := ioutil.ReadAll(result.Body)  //ioutil.ReadAll全部读取
	if err != nil {
		fmt.Println("get data err:",err)
	}

	//打印结果（bytes类型转成字符串）
	fmt.Printf("get data: %s",string(data))

}

