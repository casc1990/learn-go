package main

import (
	"net/http"
	"fmt"
	"net"
	"time"
)

var url = []string{
	"http://www.baidu.com",
	"http://www.google.com",
	"http://www.taobao.com",
}

func main() {
	for _,k := range url {
		//自定义客户端，定义超时时间（修改默认client的超时时间）
		c := http.Client{
			Transport: &http.Transport{
				Dial: func(network,addr string) (net.Conn,error) {
					timeout := time.Second *2
					return  net.DialTimeout(network,addr,timeout)
				},
			},
		}

		//head请求（请求头部）
		result,err := c.Head(k)
		if err != nil {
			fmt.Printf("get %s head request failed,err:%v\n",k,err)
			continue
		}

		//打印结果（result.Status：head请求的状态码）
		fmt.Printf("get %s head request success,status code: %v\n",k,result.Status)

		}
}