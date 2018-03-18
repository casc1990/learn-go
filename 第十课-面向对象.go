package main

//method:是附属在一个给定的类型上的，他的语法和函数的声明语法几乎一样，只是在func后面增加了一个receiver(也就是method所依从的主体)
//  定义语法：func (r 作用于的那个主体) methodName(参数,...) (返回值类型) {}
//如果一个method的receiver是*T,你可以在一个T类型的实例变量V上面调用这个method，而不需要&V去调用这个method
//如果一个method的receiver是T，你可以在一个T类型的变量P上面调用这个method，而不需要*P去调用这个method
import (
	"fmt"
	"math"
)

type Rectangle struct {  //定义长方形结构体
	width, height float64
}
type Circle struct {  //定义圆形的结构体
	radius float64
}
//定义Rectangle的aera方法，调用方法不用传值
func (r Rectangle) aera() float64 {
	return r.height * r.width
}
//定义Circle的aera方法，调用方法不用传值（mothod）
func (c Circle) aera() float64 {
	return c.radius * c.radius * math.Pi
}

//复杂一点的mothod的例子
const (
	WHITE = iota
	BLACK
	BLUE
	RED
	YELLOW
)
type Color byte
type Box struct {
	width,height,depth float64
	color Color  //color是byte类型
}
type BoxList []Box  //结构体类型的数组
func (b Box) Volume() float64 {  //定义Box的Volume方法
	return b.width * b.height * b.depth
}
//定义Box的SetColor方法，这个方法需要传一个color类型的值，方法无返回值
func (b *Box) SetColor(c Color) {
	b.color = c  //相当于：*b.color = c (这里b是指针类型，在方法里使用，go会自动帮我们转化)
}
//比较体积最大的Box的color
func (b1 BoxList) BiggestColor() Color {
	v := 0.00
	k := Color(WHITE)  //初始化color并赋值给k（k=0）
	for _, b := range b1 {
		if bv := b.Volume(); bv > v {
			v = bv
			k = b.color  //如果体积最大，记录color的值
		}
	}
	return k
}
//循环每个Box结构体，调用结构体的SetColor方法，给实例设置颜色
func (b1 BoxList) PaintItBlack() {
	for i,_ := range b1 {
		b1[i].SetColor(BLACK)  //b1[i].SetColor(1)  //相当于: &b1[i].SetColor(BLACK)
	}
}
func (c Color) String() string {
	strings := []string {"WHITE", "BLACK", "BLUE", "RED", "YELLOW"}
	return strings[c]
}

//方法的继承（和结构体的继承一样）
type Human struct {
	name string
	age int
	phone string
}
type Student struct {
	Human  //匿名字段
	school string
}
type Employee struct {
	Human
	company string
}
//在human上面定义了一个method
func (h *Human) SayHi () {
	fmt.Printf("Hi, I am %s you can call me on %s\n",h.name,h.phone)
}

//方法的重写
func (e *Employee) SayHi() {  ////Employee的method重写Human的method
	 fmt.Printf("Hi, I am %s, I work at %s. Call me on %s\n", e.name,
        e.company, e.phone)
}

func main() {
	r1 := Rectangle{12,2}  //赋值并初始化
	c1 := Circle{10}
	//访问实例的方法--aera()
	r1_areas := r1.aera()
	fmt.Println("Area of r1 is:",r1_areas)  //Area of r1 is: 24
	fmt.Println("Area of c1 is:",c1.aera())

	// 复杂的方法套用的例子
	boxes := BoxList {  //实例化BoxList
        Box{4, 4, 4, RED},
        Box{10, 10, 1, YELLOW},
        Box{1, 1, 20, BLACK},
        Box{10, 10, 1, BLUE},
        Box{10, 30, 1, WHITE},
        Box{20, 20, 20, YELLOW},
	}
    	fmt.Printf("We have %d boxes in our set\n", len(boxes))
    	fmt.Println("The volume of the first one is", boxes[0].Volume(), "cm³")  //调用Box的Volume()
    	fmt.Println("The color of the last one is",boxes[len(boxes)-1].color.String())  //调用Color的String()方法
    	fmt.Println("The biggest one is", boxes.BiggestColor().String()) //调用BoxList的BiggestColor()方法
    	fmt.Println("Let's paint them all black")
    	boxes.PaintItBlack() //调用BoxList的PaintItBlack()方法（修改所有的实例为BLACK颜色）
    	fmt.Println("The color of the second one is", boxes[1].color.String()) //调用Color的String()方法
    	fmt.Println("Obviously, now, the biggest one is", boxes.BiggestColor().String()) //调用Color的String()方法

	//method的继承(Student继承了Human方法)
	mark := Student{Human:Human{"Mark",25,"123456"},school:"IT"}
	mark.SayHi() //继承了Human的方法；Hi, I am Mark you can call me on 123456

	////method的重写(Employee重写了Human方法)
	sam := Employee{Human{"Sam", 45, "111-888-XXXX"}, "Golang Inc"}
	sam.SayHi()  //重写了Human的方法；Hi, I am Sam, I work at Golang Inc. Call me on 111-888-XXXX

}
