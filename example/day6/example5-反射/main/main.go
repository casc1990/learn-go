package main

import (
	"reflect"
	"fmt"
)

//反射：可以在运行时动态获取变量的相关信息

type Student struct {
	Name string
	Age int
	Score float32
}

func test(a interface{}) {
	//reflect.TypeOf查看对象的类型
	t := reflect.TypeOf(a)
	fmt.Printf("%v type is: %s\n",a,t) //{stu01 32 90} type is: main.Student

	//reflect.ValueOf查看变量的值
	v := reflect.ValueOf(a)
	fmt.Println(v) //{stu01 32 90}

	//
	k := v.Kind()
	fmt.Println(k) //查看值的底层数据类型，输出：struct

	//将reflect.ValueOf转换成接口类型
	iv := v.Interface()
	//判断接口是否是Student类型
	stu,ok := iv.(Student)
	if ok {
		fmt.Printf("%v %T\n",stu,stu) //{stu01 32 90} main.Student
	}




}


func main() {
	a := Student{
		Name: "stu01",
		Age: 32,
		Score: 90,
	}

	test(a)

}


