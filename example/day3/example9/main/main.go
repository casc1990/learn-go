package main

import (
	"fmt"
	"bufio"
	"os"
	"strings"
)

func multi(str1,str2 string) (result string) {
	str1_length := len(str1) - 1
	str2_length := len(str2) - 1
	left := 0  //存进位
	//如果字符串长度为空，返回0
	if str1_length == 0 && str2_length == 0 {
		result = "0"
		return
	}
	//从右到左，字符串转换成数字相加（计算到最短位）
	for str1_length >=0 && str2_length >= 0 {
		c1 := str1[str1_length] - '0'  //利用字符串的ascii码求值
		c2 := str2[str2_length] - '0'
		//fmt.Printf("c1 is %d,c2 is %d,c1 type %T,c2 type %T\n",c1,c2,c1,c2)
		sum := int(c1) + int(c2) + left
		//如果大于10，进位
		if sum >= 10 {
			left = 1
		}else {
			left = 0
		}
		c3 := (sum % 10) + '0'  //计算数字的ascii值，相当于 数字N加上0的ascii码就是某个数的ascii值
		result = fmt.Sprintf("%c%s",c3,result)
		fmt.Printf("-%c%s\n",c3,result)
		str1_length--
		str2_length--
	}

	//第一个数进位
	fmt.Println("第一个数进位",left)
	for str1_length >= 0 {
		c1 := str1[str1_length] - '0'
		sum := int(c1) + left
		if sum >= 10 {
			left = 1
		}else {
			left = 0
		}
		c3 := (sum % 10) + '0'
		result = fmt.Sprintf("%c%s",c3,result)
		fmt.Printf("--%c%s\n",c3,result)
		str1_length--
	}
	//第二个数进位
	fmt.Println("第二个数进位",left)
	for str2_length >= 0 {
		c2 := str2[str1_length] - '0'
		sum := int(c2) + left
		if sum >= 10 {
			left = 1
		}else {
			left = 0
		}
		c3 := (sum % 10) + '0'
		result = fmt.Sprintf("%c%s",c3,result)
		fmt.Printf("---%c%s\n",c3,result)
		str2_length--
	}
	//最后进位
	fmt.Println("最后进位",left)
	if left == 1 {
		result = fmt.Sprintf("1%s",result)
		fmt.Printf("----1%s\n",result)
	}

	return
}

func main() {
	//实例化对象，从终端读取内容，并做异常判断
	reader := bufio.NewReader(os.Stdin)
	result,_,err := reader.ReadLine()
	if err != nil {
		fmt.Print("read from console err",err)
		return
	}
	//将字符串以“+”分割，异常退出
    strSlice := strings.Split(string(result),"+")
    if len(strSlice) != 2 {
    	fmt.Println("please input a+b")
	}
	//调用函数，传递相加的字符串
	strNumber1 := strSlice[0]
	strNumber2 := strSlice[1]
	fmt.Println(multi(strNumber1,strNumber2))
}
