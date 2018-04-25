package main

import "fmt"

//接口：简单的说，interface是一组method的组合，我们通过interface来定义对象的一组行为
//interface类型定义了一组方法，如果某个对象实现了某个接口的所有方法，则此对象就实现了此接口

//定义test接口（接口不需要实现某个方法，方法又具体的对象实现）
type Test interface {
	Print()
	Sleep()
}

type Student struct {
	name string
	age int
	score float32
}

//Student实现了Print方法
func (p Student) Print() {
	fmt.Println("name:",p.name)
	fmt.Println("age:",p.age)
	fmt.Println("score:",p.score)
}

func (p Student) Sleep() {
	fmt.Printf("%s is sleep!\n",p.name)
}

type People struct {
	name string
	age int
}
//People实现了Print方法
func (people People) Print() {
	fmt.Println("name:",people.name)
	fmt.Println("age:",people.age)
}

func main() {
	var t Test
	stu := Student{
		name: "stu1",
		age: 18,
		score: 90,
	}
	//stu实现了t接口的所有方法(Print(),Sleep())，所有stu就实现了t接口(t就可以等于stu)
	t= stu
	//调用t接口的Print方法（实际上调用的是stu的Print方法）
	t.Print()
	//调用Sleep方法
	t.Sleep()  //stu1 is sleep!

	//打印t，t是指向stu1
	fmt.Println("t:",t) //t: {stu1 18 90}

	people := People{
		name: "people",
		age : 29,
	}
	//t = people
	fmt.Println(people)
	fmt.Println("people未实现Test接口的Sleep方法，所有它未实现Test接口")
}