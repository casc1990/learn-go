package main

import "fmt"

//输入一个字符串，判断是否是回文
//思路：从中间截断，第一个字符和最后一个比较，依次类推。符合条件，返回true
func Process(str string ) bool {
	//Rune 是int32 的别名。用UTF-8 进行编码.rune表示的是字符，byte表示的是字节。（中文字符就3字节）
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