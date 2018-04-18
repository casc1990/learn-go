package main

import (
	"fmt"
	"errors"
)

func initConfig() (err error) {
	return errors.New("init config failed")
}

func recovers() {
	// defer函数调用recover可以捕获异常
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()

	a := 0
	b := 100/ a  //除数为0会触发painc
	fmt.Println(b)
	return

}

func main() {
	var i int
	fmt.Println(i)

	//new用来分配值类型的内存地址，生成指针对象
	j := new(int)
	fmt.Println(j) // 输出: 0xc0420080d0
	*j = 100  //指针赋值
	fmt.Println(*j) //打印指针指向的值，输出:100

	//make：用来分配引用类型的内存地址，比如chan、map、slice，生成的是数据类型
	k := make([]int,5)
	fmt.Println(k) //输出：[0 0 0 0 0]

	//append用来追加元素到数组、slice中
	k = append(k,1,2,3,4,5)
	fmt.Println(k) //  [0 0 0 0 0 1 2 3 4 5]
	s := []int{6,7,8,9}

	//将slice追加到slice中
	k = append(k,s...)
	fmt.Println(k) //[0 0 0 0 0 1 2 3 4 5 6 7 8 9]

	recovers() //调用recovers函数

	err := initConfig()
	if err != nil {
		fmt.Println(err)
		panic(err)  //调用panic函数，直接panic
	}
}