package main

import (
	"fmt"
	"time"
)

func test() {
	time.Sleep(time.Millisecond * 300)
}

func main() {
	now := time.Now()
	fmt.Println(time.Now()) //打印当前时间
	fmt.Println(time.Now().Day()) //打印天，输出：25
	fmt.Printf("%d-%d-%d\n",now.Year(),now.Month(),now.Day()) //2018-3-25
	//时间格式化，固定写法--切记
	fmt.Println(now.Format("2006/01/02 15:04")) //2018/03/25 17:41
	fmt.Println(now.Format("01/02/2006 15:04")) //03/25/2018 17:41

	start := time.Now().UnixNano() //距今经过的纳秒数
	test()
	end := time.Now().UnixNano()
	fmt.Printf("执行程序花费了:%dms\n",(end-start)/1000) //执行程序花费了:300017ms

}


