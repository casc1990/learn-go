package main

/* 练习
  定义两个常量Man=1和Female=2,获取当前时间的秒数，如果能被Female整除,则在终端打印female，否则打印man。*/

import (
	"fmt"
	"time"
)

const (
	Man int = 1
	Female int =2
)


func main() {
	for i :=0;i < 10;i++ {
		Second := time.Now().Unix()
		fmt.Println("now time:", Second)
		if Second % 2 == 0 {
			fmt.Println("female")
		}else {
			fmt.Println("man")
		}
		time.Sleep(1000 * time.Millisecond)  //time.Millisecond毫秒
	}
}


