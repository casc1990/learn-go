package main

import (
	"reflect"
	"fmt"
	_"testing"
)

//反射：可以在运行时动态获取变量的相关信息
// reflect.TypeOf，获取变量的类型，返回reflect.Type类型
// reflect.ValueOf，获取变量的值，返回reflect.Value类型
// reflect.ValueOf(x).Float() 获取指定类型变量的值
// reflect.Value.SetXX相关方法，比如:reflect.Value.SetInt





type Student struct {
	Name string
	Age int
	Score float32
}

func test(a interface{}) {
	//reflect.TypeOf查看对象的定义类别，返回reflect.Type类型
	t := reflect.TypeOf(a)
	fmt.Printf("%v type is: %s\n",a,t) //{stu01 32 90} type is: main.Student

	//reflect.ValueOf查看变量的值，返回reflect.Value类型
	v := reflect.ValueOf(a)
	fmt.Println(v) //{stu01 32 90}

	//reflect.ValueOf(元素).Kind() 获取值的底层数据类型
	k := v.Kind()
	fmt.Println(k) //查看值的底层数据类型，输出：struct

	//将reflect.ValueOf转换成接口类型
	iv := v.Interface()
	//接口断言：将iv接口转换成Student类型，如果转换成功，ok=true
	stu,ok := iv.(Student)
	if ok {
		fmt.Printf("%v %T\n",stu,stu) //{stu01 32 90} main.Student
	}

}


func TetsInt(a interface{}) {
	//reflect.ValueOf查看变量的值，返回reflect.Value类型
	val := reflect.ValueOf(a)
	//val.Int()获取int类型变量的值(如果传递的是指针的，要使用Elem()方法获取指针的值)
	c := val.Elem().Int()
	fmt.Printf("val value:%d\n", c)

	//改变变量的值
	val.Elem().SetInt(100)
	fmt.Printf("val value:%d\n",val.Elem().Int())
}

func main() {
	a := Student{
		Name: "stu01",
		Age: 32,
		Score: 90,
	}

	test(a)

	//SetInt和Elem的用法
	b := 1234
	TetsInt(&b)
}


