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
	left := 0

	if str1_length == 0 && str2_length == 0 {
		result = "0"
		return
	}

	for str1_length >=0 && str2_length >= 0 {
		c1 := str1[str1_length] - '0'  //利用字符串的ascii码求值
		c2 := str2[str2_length] - '0'
		//fmt.Printf("c1 is %d,c2 is %d,c1 type %T,c2 type %T\n",c1,c2,c1,c2)
		sum := int(c1) + int(c2) + left

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

	if left == 1 {
		result = fmt.Sprintf("1%s",result)
		fmt.Printf("----1%s\n",result)
	}

	return
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	result,_,err := reader.ReadLine()
	if err != nil {
		fmt.Print("read from console err",err)
		return
	}

    strSlice := strings.Split(string(result),"+")
    if len(strSlice) != 2 {
    	fmt.Println("please input a+b")

	}

	strNumber1 := strSlice[0]
	strNumber2 := strSlice[1]
	fmt.Println(multi(strNumber1,strNumber2))
}
