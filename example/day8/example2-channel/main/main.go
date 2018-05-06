package main

import "fmt"

type student struct {
	name string
}

func main() {
	//定义channel
	var interChan chan interface{}
	//channel的初始化,可以存10个元素
	interChan = make(chan interface{},10)
	//因为是空接口类型，所以可以向channel里发送各种数据类型
	str := student{name:"stu01"}
	interChan <- &str

	//存进去是空接口类型，取也要是空接口类型
	var stu01 interface{}
	stu01 = <- interChan
	fmt.Printf("%s type is %T\n",stu01,stu01) //{stu01} type is main.student

	//将接口类型在转换为结构体类型
	stu02,ok := stu01.(*student)
	if !ok {
		fmt.Println("can not convert")
		return
	}
	fmt.Println(stu02)

}
