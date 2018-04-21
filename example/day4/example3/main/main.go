package main

import (
	"fmt"
	"strings"
)

//闭包： 一个函数和与其相关的引用环境组合而成的实体（函数和外部变量绑定）
//		可以理解为类方法引用类变量
//闭包形式：定义函数的返回值也是函数类型，返回值函数可以访问外部变量，并修改外部变量

func Adder() func(i int) int {  //函数的返回值也是函数
	 var x int
	 return func(i int) int {
	 	x += i    //返回值函数访问了外部变量，相应的外部的x的值也会改变
	 	return x
	 }
}

//闭包的例子。判断字符串是否已指定字符结尾。不是的话加上后缀
func makeSuffix(suffix string) func(string) string {
	return func(name string) string {
		if strings.HasSuffix(name,suffix) == false {  //返回函数中使用外部变量的值
			return name + suffix
		}
		return name
	}
}


func main() {
	f := Adder()
	fmt.Println(f(1))  // 1
	fmt.Println(f(100))  //101
	fmt.Println(f(1000))  //1101

	f2 := makeSuffix(".txt") //实例化对象。
	fmt.Println(f2("test"))  //test.txt
	fmt.Println(f2("log"))   //log.txt

	f3 := makeSuffix(".jpg") //实例化对象。
	fmt.Println(f3("test"))  //test.jpg
	fmt.Println(f3("log"))   //log.jpg
}