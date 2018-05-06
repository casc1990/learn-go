package main

import (
	"time"
	"fmt"
)

func main() {
	ch := make(chan int)
	ch2 := make(chan int)

	go func() {
		var i int
		for {
			ch <- i
			time.Sleep(time.Second)
			ch2 <- i*i
			time.Sleep(time.Second)
			i++
		}
	}()

	//select用于多channel切换（那个channel不阻塞，执行那个分支，如果都不阻塞，按顺序执行）
	for {
		select {
		case v:= <- ch:  //这个channel阻塞，执行下面的case分支
			fmt.Println(v)
		case v:= <- ch2:
			fmt.Println(v)
		default:  //都阻塞，执行默认分支
			fmt.Println("get data timeout!")
			time.Sleep(time.Second)


		}
	}
}
