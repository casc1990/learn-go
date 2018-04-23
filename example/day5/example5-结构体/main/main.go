package main

import "fmt"

type Student struct {
	Name string
	Age int
	score float32
}

func main() {
	//初始化1
	var str Student
	str.Name = "zhangsan"
	str.Age = 18
	fmt.Printf("Name:%s,Age:%d,score:%f\n",str.Name,str.Age,str.score) //Name:zhangsan,Age:18,score:0.000000

	//初始化2
	var str1 = Student{
		Name: "zhangsan",
		Age: 20,
		score: 80,
	}
	fmt.Printf("Name:%s,Age:%d,score:%f\n",str1.Name,str1.Age,str1.score)

	//初始化3
	var str3 *Student = &Student{
		Age: 28,
		Name: "hua",
	}
	fmt.Printf("Name:%s,Age:%d,score:%f",str3.Name,str3.Age,str3.score)


}
