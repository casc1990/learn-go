package main

import (
	"fmt"
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())  //添加随机种子(以毫秒为种子)。保证每次随机数不重复
}

func main()  {
	//rand.Seed(time.Now().Unix())  //添加随机种子。保证每次随机数不重复
	for i :=0;i < 10;i++ {
		a := rand.Int()
		fmt.Println("生成的随机int类型整数:",a)
	}

	for i :=0;i<10;i++ {
		a := rand.Intn(100)
		fmt.Println("生成100以内的随机整数:",a)
	}

	for i :=0;i<10;i++ {
		a := rand.Float32()
		fmt.Println("生成随机float32类型的浮点数:",a)
	}
}
