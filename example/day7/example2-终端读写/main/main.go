package main

import (
	"fmt"
	"os"
)

var (
	firstName,lastName,s string
	i int
	f float32
	input = "56.12 / 5212 / Go"
	format = "%f / %d / %s"
)

func main() {
	fmt.Println("please enter your full name: ")
	//fmt.Scanln和fmt.Scanf从终端读取数据，如果牵扯到修改数据的值，一定要传地址
	fmt.Scanln(&firstName,&lastName) //从终端读取数据，只能读空格分割的字符串，和fmt.Scanf效果一样
	//fmt.Scanf("%s %s",&firstName,&lastName) //同上,fmt.Scanf可以将读到的数据格式化
	fmt.Printf("Hi %s %s!\n",firstName,lastName)

	//fmt.Sscanf从字符串里格式化输出
	fmt.Sscanf(input,format,&f,&i,&s) //从input字符串里按照format定义的格式化方法，将格式化后的存入指定变量中(f,i,s)
	fmt.Println("From the strings we read:",f,i,s)

	//写入文件(如果文件不存在，创建，文件权限664)
	file,err := os.OpenFile("c:/test.log",os.O_CREATE|os.O_WRONLY,0664)
	if err != nil {
		fmt.Println("open file err:",err)
		return
	}
	fmt.Fprint(file,"this is")
}