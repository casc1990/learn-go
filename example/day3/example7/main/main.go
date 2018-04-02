package main

import "fmt"

//输入一个字符串，判断是否是回文

func Process(str string ) bool {
	//Rune 是int32 的别名。用UTF-8 进行编码.为了获得实际的字符，需要使用rune 类型
	result := []rune(str)
	length := len(result)
	for i,_ := range result {
		if i == length/2 {
			break
		}
		last := length -i -1
		if result[i] != result[last] {
			return false
		}
	}
	return true
}

func main() {
	var str string
	fmt.Scanf("%s",&str)
	if Process(str) {
		fmt.Printf("%s 是回文",str)
	}else {
		fmt.Printf("%s 不是回文",str)
	}
}