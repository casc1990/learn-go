package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	var num int
	rand.Seed(time.Now().UnixNano()) //利用当前时间的UNIX时间戳初始化rand包
	randnum := rand.Intn(100)
	fmt.Println("生成的随机数:",randnum)
	for i :=0;i <3;i++ {
		fmt.Scanf("%d",&num)
		switch {
		case num == randnum:
			fmt.Println("恭喜猜对了..")
		case num > randnum:
			fmt.Println("往小的猜..")
		case num < randnum:
			fmt.Println("往大的猜..")
		default:
			fmt.Println("超出边界..")
		}
	}
}
