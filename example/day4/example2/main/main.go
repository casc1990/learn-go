package main

import (
	"fmt"
	"time"
)

//函数递归：自己调用自己
//原则：问题规模会不断缩小
//	   要有退出条件


func recusive(n int) {
	fmt.Println(n)
	time.Sleep(time.Second)
	if n >= 3 {
		return
	}
	recusive(n+1) //自己调用自己
}

func factor(n int) int {
	if n == 1 {
		return 1
	}
	return factor(n-1) * n
}

func fab(n int) int {
	if n <= 1 {
		return 1
	}
	return fab(n-1) + fab(n-2)
}

func main() {
	recusive(0)
	fmt.Println(factor(5))
	fmt.Println(fab(10))
}