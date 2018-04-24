package main

import "fmt"

//方法:方法是针对某个类型的
//方法和函数的区别： 函数调用： function(variable，参数列表)   方法： variable.function(参数列表)
//方法的继承也是通过匿名字段来实现.如果一个struct嵌套了多个匿名结构体，那么这个结构可以直接访问多个匿名结构体的方法，从而实现了多重继承。


type integer int

//int类型的方法（任何类型都可以有方法）
func (p integer) print() {
	fmt.Println(p)
}

//int类型的set方法
func (p *integer) set(b integer) {
	*p = b
}

type Student struct {
	Name string
	Age int
	Score float32
	sex int
}

//结构体Student的init方法（结构体是指针类型，要修改，就要传指针）
func (p *Student) init (name string,age int,score float32) {
	p.Name = name
	p.Age = age
	p.Score = score
	fmt.Println(p)
}

//结构体Student的get方法
func (p Student) get() Student {
	return p
}


func main() {
	var stu Student
	//调用结构体init方法
	stu.init("stu",18,80)

	//调用结构体get方法
	stu1 := stu.get()
	fmt.Println(stu1)  //{stu 18 80 0}

	var a integer
	a = 100
	a.print() //100
	//调用set方法
	a.set(1000)
	a.print() //1000

}