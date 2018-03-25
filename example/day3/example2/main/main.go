package main

import (
	"fmt"
	"strconv"
)

func shuixianhua(n string) bool {
	var result int
	for i :=0;i <len(n);i++ {
		num := int(n[i] - '0')
		result += num*num*num

		num1,err := strconv.Atoi(n)
		if err != nil {
			fmt.Printf("can not convert %s to int\n",n)
			break
		}
		if num1 == result {
			return true

		}
	}
	return false
}

func main() {
	var min,max string
	var num int
	fmt.Scanf("%s%s",&min,&max)
	min1,_ := strconv.Atoi(min)
	max1,_ := strconv.Atoi(max)
	for i:=min1;i<max1;i++ {
		if  j := shuixianhua(strconv.Itoa(i));j {
			fmt.Printf("%d 是水仙花数.\n",i)
			num += 1
		}
	}
	fmt.Printf("%d-%d的水仙花数有:%d",min1,max1,num)
	}



