package main

import "fmt"

type student struct {
	Name string
	Age int
}

//定义构建函数（构建结构体）
func NewStudent(name string,age int) *student {
	return  &student{
		Name: name,
		Age: age,
	}
}


func main() {
	s := new(student)
	s.Name = "Hey"

	//自定义构建函数创建结构体
	s1 := NewStudent("Tony",20)
	fmt.Println("hello! ",s1.Name)


}