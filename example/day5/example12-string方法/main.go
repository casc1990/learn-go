package main

import "fmt"

//string()方法：如果某个对象实现了string方法，在打印这个变量的时候就会自动调用string里定义的内容，如果没有实现，就默认输出


type Car struct {
	weight int
	name string
}

type Bike struct {
	Car  //继承了Car结构体
	lunzi int
}

//结构体方法
func (p Car) Run() {
	fmt.Printf("%s is running!!\n",p.name)
}

//string方法
func (p Bike) String() string {
	str := fmt.Sprintf("name=[%s],weight=[%d]",p.name,p.weight)
	return str
}

func main() {
	var a Bike
	a.weight = 300
	a.lunzi = 2
	a.name = "bike"
	a.Run() //调用run方法

	//打印指定的类型时就会自动的调用String()方法
	fmt.Println(a)   //name=[bike],weight=[300]

	}





