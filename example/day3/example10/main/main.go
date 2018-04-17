package main

import (
	"bufio"
	"os"
	"fmt"
	"strings"
)

func process(str1,str2 string) (result string) {
	str1_index := len(str1) - 1
	str2_index := len(str2) - 1
	left := 0

	if str1_index == 0 && str2_index == 0 {
		result = "0"
		return
	}

	for str1_index >= 0 && str2_index >= 0 {
		c1 := str1[str1_index] - '0'
		c2 := str2[str2_index] - '0'
		fmt.Printf("c1 value %d,c2 value %d\n",c1,c2)
		sum := int(c1) - int(c2) - left
		if int(c1) > int(c2) {
			left = 0
		}else {
			sum = 10 + int(c1) - int(c2) - left
			left = 1
		}
		c3 := sum + '0'
		fmt.Printf("c3 is %c\n",c3)
		result = fmt.Sprintf("%c%s",c3,result)
		str1_index--
		str2_index--
	}

	for str1_index >= 0 {
		c1 :=  str1[str1_index] - '0'
		sum := int(c1) - left
		if sum < 0 {
			left = 1
			} else {
				left = 0
			}
			c3 := sum + '0'
			result = fmt.Sprintf("%c%s",c3,result)
		str1_index--
	}

	return
}


func main() {

	reader := bufio.NewReader(os.Stdin)
	result,_,err := reader.ReadLine()
	if err != nil {
		fmt.Println("read from console err!!")
		return
	}

	str_slice := strings.Split(string(result),"-")
	if len(str_slice) != 2  {
		fmt.Println("please input a-b type")
		return
	}

	str1 := str_slice[0]
	str2 := str_slice[1]

	fmt.Println(process(str1,str2))
}