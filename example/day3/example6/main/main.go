package main

import (
	"fmt"
)

//回文数判断函数（从左到右和从右到左是一样的）
func reverseString(s string)  {
	result := []int32(s)
	// 字符串翻转。通过下标，前后位置互换
	for from, to := 0, len(result)-1; from < to; from, to = from+1, to-1 {
		result[from], result[to] = result[to], result[from]
	}
	result_str := string(result)
	if result_str == s {
		fmt.Println("这个字符串是回文数:", s)
	} else {
		fmt.Println("抱歉，这个不是...")
	}

}

func main() {
	fmt.Println(string(99))  //string函数把数字转换成字符
	var str string
	fmt.Scanf("%s",&str)
	reverseString(str)

}
