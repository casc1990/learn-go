package go基础语法
//判断interface变量存储的对象类型：第一：Comma-ok断言  第二：switch测试
//判断是否是该类型的变量： value, ok = element.(T)，这里value就是变量的值，ok是一个bool类型，element是interface变量，T是断言的类型。
import (
	"fmt"
	"strconv"  //strconv.Itoa(int):将int转换成string；strconv.Atoi(string)：将string转换成int
	"math"
)


type Element interface {} //定义空接口
type List []Element  //定义了存储接口的slice
type Person struct {
	name string
	age int
}
//Person定义了String方法，实现fmt.String
func (p Person) String() string {
	return "name:" + p.name + " age: "+ strconv.Itoa(p.age) + " years"
}

//interface嵌入：如果一个interface1作为interface2的一个嵌入字段，那么interface2隐式的包含了interface1里面的method
type Rectangle struct {
	width float64
	height float64
}
type Circle struct {
	radius float64
}
func (r Rectangle) area() float64 {
	return r.width * r.height
}
func (c Circle) area() float64 {
	return math.Pi * math.Pow(c.radius,2)   //math.Pow(int,2)：求int的2次方
}
//定义Shape接口
type Shape interface {
	area() float64
}
type MultiShape interface {
	Shape   //嵌入Shape接口，继承Shape的所有method
}
//定义GetArea函数，接收实现了MultiShape接口的对象实例，然后调用对象实例的area方法
func GetArea(shape MultiShape) float64 {
	return shape.area()   //MultiShape继承了Shape接口的area()
}


func main() {
	list := make(List,3)  //定义一个能存储3个元素的空接口类型的slice
	list[0] = 1 // an int
	list[1] = "Hello" // a string
	list[2] = Person{"Dennis", 70} //a struct
	//Comma-ok断言
	for index,element := range list {
		if value,ok := element.(int); ok {  //判断对象是否是int类型
			fmt.Printf("list[%d] is an int and its value is %d\n", index, value)
		}else if value,ok := element.(string); ok  {
			fmt.Printf("list[%d] is an string and its value is %s\n", index, value)
		}else if value,ok := element.(Person); ok {
			fmt.Printf("list[%d] is an struct and its value is %s\n", index, value.name)
		}
	}

	//switch测试(仅适用测试interface类型)
	for index,element := range list {
		switch value := element.(type) {
			case int :
				fmt.Printf("list[%d] is an int and its value is %d\n", index, value)
			case string :
				fmt.Printf("list[%d] is an string and its value is %s\n", index, value)
			case Person:
				fmt.Printf("list[%d] is an struct and its value is %s\n", index, value.name)
			case nil :
				fmt.Println("类型 :%T",value)
			default:
         			fmt.Printf("未知型")
			}
		}

	//interface嵌入：如果一个interface1作为interface2的一个嵌入字段，那么interface2隐式的包含了interface1里面的method
	r := Rectangle{20,10}
	c := Circle{4}
	//调用GetArea函数，计算面积（调用对象实例的area方法）
	fmt.Println("Rectangle Area = ",GetArea(r))  //Rectangle Area =  200
	fmt.Println("Circle Area =",GetArea(c))
	}