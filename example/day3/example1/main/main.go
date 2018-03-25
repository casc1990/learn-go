package main

import "fmt"

func IsPrime(n int) bool {
	if n % 2 == 0 {
		return false
	}
	return true
}

func main() {
	var min,max,num int
	fmt.Scanf("%d%d",&min,&max)  //接收控制台输入，并把值赋值给相应变量（如果使用值类型，那么min，max就只能等于默认值0）
	for i :=min;i < max;i++ {
		if IsPrime(i) == true {
			fmt.Println("素数:",i)
			num += 1
		}
	}
	fmt.Printf("%d-%d之间的素数有:%d",min,max,num)
}