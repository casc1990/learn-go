package main

import (
	"fmt"
	"reflect"
)

// reflect.Value.NumField()获取结构体中字段的个数
// reflect.Value.Method(n).Call来调用结构体中的方法

type Student struct {
	Name string
	Age int
	Score float32
}

func (p Student) Print() {
	fmt.Println("----start----")
	fmt.Println(p)
	fmt.Println("----stop----")
}

//*Student是指针Student的方法，并不是Student的方法
func (p *Student) Set(name string,age int,score float32) {
	p.Name = name
	p.Age = age
	p.Score = score
}


func TestStruct(a interface{}) {
	//将a转换成reflect.ValueOf类型
	val := reflect.ValueOf(a)
	//获取a的底层数据类型
	kd := val.Kind()
	if kd != reflect.Struct {
		fmt.Println("传递的类型不是struct")
		return
	}

	//获取有多少个字段
	num := val.NumField()
	fmt.Println("struct有%d个字段",num)
	//打印某个字段
	field2 := val.Field(2)
	fmt.Println("strcut第2个字段的值是: ",field2) //strcut第2个字段的值是:  98.2
	//val.Field(2).SetInt(99)

	//获取有多少方法
	method := val.NumMethod()
	fmt.Println("struct有%d个方法",method)
	//通过下标调用方法(call可以传递参数)
	var params []reflect.Value
	val.Method(0).Call(params)

}

func main() {
	a := Student{
		Name: "stu1",
		Age: 28,
		Score: 98.2,
	}

	TestStruct(a)
}