package main

import (
	"fmt"
)

func main() {
	str1 := "hello!world!\n"  //双引号会解释特殊字符
	str2 := `
			id:
			name:
			age: `           //``符号会原样输出,合适多行输出
	b := 'c'   //字符类型(bytes)使用单引号定义
	fmt.Println(str1) //hello!world!
	fmt.Println(str2)
	fmt.Println(b) //输出：99 ；默认会输出字符对应的ascii编号
	fmt.Printf("字符的字符串形式：%c",b) //字符的字符串形式：c
}