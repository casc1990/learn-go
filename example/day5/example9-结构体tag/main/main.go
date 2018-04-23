package main

import (
	"encoding/json"
	"fmt"
	"reflect"
)

//结构体tag: 用来描述结构体字段.经常用于转存json,bson等类型
//在golang中，命名都是推荐都是用驼峰方式，并且在首字母大小写有特殊的语法含义：包外无法引用。
//标签（Tag），在转换成其它数据格式的时候，会使用其中特定的字段作为键值

type Student struct {
	Name string   `json:"student_name" bson:"student_name"`
	Age int  `json:"student_age" bson:"student_age"`
	Score float32  `json:"student_score" bson:"student_score"`
}

func main() {
	stu := Student{
		Name: "hey",
		Age: 18,
		Score: 80,
	}

	//转换成json格式
	data,err := json.Marshal(stu)
	if err != nil {
		fmt.Println("json encode stu failed,err:",err)
		return
	}
	//使用json格式的定义的字段输出
	fmt.Println(string(data)) //{"student_name":"hey","student_age":18,"student_score":80}


	//获取tag(使用反射包（reflect）中的方法来获取)
	t := reflect.TypeOf(&stu)
	field := t.Elem().Field(0)
	fmt.Println(field.Tag.Get("json")) //获取json类型  student_name
	fmt.Println(field.Tag.Get("bson")) //获取bson类型  student_name
}





