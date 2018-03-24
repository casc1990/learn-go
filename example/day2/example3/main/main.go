package main

import "fmt"

var (
	a string = "lisi"
	b bool
	c byte = 'a'
	d int
)
func main() {
	e := 345
	fmt.Printf("%v\n", a) //%v 默认类型
	fmt.Printf("%v\n", c)
	fmt.Printf("值的类型:%T\n", b)    //%T 值的类型
	fmt.Printf("值的类型:%T\n", e)    //输出:值的类型:int
	fmt.Printf("90%%\n")          //90%  %%转义
	fmt.Printf("打印布尔类型:%t\n", b)  // %t 返回布尔类型的true和false
	fmt.Printf("打印值的二进制:%b\n", 300) //打印值的二进制:100101100
	fmt.Printf("打印值的二进制:%d\n", 300)
	fmt.Printf("打印字符串类型:%s\n",a)
	fmt.Printf("打印浮点型:%f\n",3.14444)
	fmt.Printf("%q\n","this is a book..") //加双引号；"this is a book.."
	fmt.Printf("%p\n",&a) //打印a的地址

	str := fmt.Sprintf("a=%d",100) //将打印的结果报存到str
	fmt.Print(str)


}