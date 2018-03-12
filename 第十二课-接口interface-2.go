package main
//判断interface变量存储的对象类型：第一：Comma-ok断言  第二：switch测试
//判断是否是该类型的变量： value, ok = element.(T)，这里value就是变量的值，ok是一个bool类型，element是interface变量，T是断言的类型。
import (
	"fmt"
	"strconv"
)


type Element interface {} //定义空接口
type List []Element  //定义了存储接口的slice
type Person struct {
	name string
	age int
}
//Person定义了String方法，实现fmt.Stringrt
func (p Person) String() string {
	return "name:" + p.name + " age: "+ strconv.Itoa(p.age) + " years"
}

func main() {
	list := make(List,3)
	list[0] = 1 // an int
      	list[1] = "Hello" // a string
      	list[2] = Person{"Dennis", 70} //a struct
	//Comma-ok断言
	for index,element := range list {
		if value,ok := element.(int); ok {
			fmt.Printf("list[%d] is an int and its value is %d\n", index, value)
		}else if value,ok := element.(string); ok  {
			fmt.Printf("list[%d] is an string and its value is %s\n", index, value)
		}else if value,ok := element.(Person); ok {
			fmt.Printf("list[%d] is an struct and its value is %s\n", index, value.name)
		}
	}

	//switch测试
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

	//interface嵌入

	}