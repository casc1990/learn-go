package main

import (
	"fmt"
	"os"
	"flag"
)

func main() {
	//如果不加参数，默认长度为1，参数为命令行本身
	fmt.Printf("len of args:%d",len(os.Args))
	//循环参数列表
	for i,v := range os.Args {
		fmt.Printf("agrs[%d]=%s\n",i,v)
	}
	fmt.Println(os.Args[1]) //打印第一个参数


	//命令行参数分析
	var configPath string
	var logLevel int
	//例如：example.exe -c c:/test.log -d 10 匹配-c -d这样选项的值，用户未输入，使用默认，后面是说明
	//在将选项的值，保存到变量（&configPath）
	flag.StringVar(&configPath,"c"," ","please input config path!")
	flag.IntVar(&logLevel,"d",0,"please input log level!")

	//开始解析
	flag.Parse()
	//打印命令行参数的值
	fmt.Println("path:",configPath)
	fmt.Println("logLevel:",logLevel)
}
