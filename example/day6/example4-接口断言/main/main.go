package main

import "fmt"

//接口断言：b,ok := a.(string) 判断接口是否是string类型

type Student struct {
	Name string
	Sex string
}

func Test(a interface{}) {
	//接口断言，如果接口是string类型，ok==true，b==具体的字符串
	b,ok := a.(string)
	if ok == false {
		fmt.Printf("%v not is string\n",a)
		return
	}
	//是string类型，则打印
	fmt.Println(b)
}

//类型断言
func Just(items ...interface{}) {
	for index,value := range items {
		switch value.(type) {
			case bool:
				fmt.Printf("%d params is bool! value is %v\n",index,value)
			case int,int64,int32:
				fmt.Printf("%d params is int! value is %v\n",index,value)
			case float32,float64:
				fmt.Printf("%d params is float! value is %v\n",index,value)
			case string:
				fmt.Printf("%d params is string! value is %v\n",index,value)
			case Student:  //判断是否是结构体
				fmt.Printf("%d params is student! value is %v\n",index,value)
			case *Student:  //判断是否是指针类型
				fmt.Printf("%d params is *student! value is %v\n",index,value)
		}
	}
}

func main() {
	var a int
	Test(a) //0 not is string

	var b string = "hello"
	Test(b) //hello

	//类型推倒
	c := Student{
		Name: "stu01",
		Sex: "female",
	}
	Just(28,2.28,"string",c,&c) //判断各种类型
}