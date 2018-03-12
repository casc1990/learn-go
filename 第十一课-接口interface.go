package main
//interface: 表示一组method的组合，我们通过interface来定义对象的一组行为
//interface和struct一样，也是一种类型，使用type定义
//interface类型定义了一组方法，如果某个对象实现了某个interface的所有方法，则此对象就实现了此接口
//任意的类型都实现了空interface(我们这样定义：interface{})，也就是包含0个method的interface
//interface可以被任意的对象实现.同理，一个对象可以实现任意多个interface
import (
	"fmt"
	"strconv"
)

type Human struct {
	name string
	age int
	phone string
}
type Student struct {
	Human  //匿名字段
	school string
	loan float32
}
type Employee struct {
	Human  //匿名字段
	company string
	money float32
}
//Human实现SayHi方法
func (h Human) Sayhi() {
	fmt.Printf("Hi,I am %s you can call me on %s\n",h.name,h.phone)
}
//Human实现Sing方法
func (h Human) Sing(lyrics string) {
	fmt.Println("La la la...",lyrics)
}
//Employee重载Human的Sayhi方法
func (e Employee) Sayhi() {
	fmt.Printf("Hi,I am %s,I work at %s. Call me on %s\n",
		e.name,e.company,e.phone)
}
//Student实现BorrowMoney方法
func (s *Student) BorrowMoney(amount float32) {
	s.loan += amount  //借钱等于贷款加花费
}
//Employee实现SpendSalary方法
func (e *Employee) SpendSalary(amount float32) {
	e.money -= amount //花钱等于工资减花费
}

//定义interface
//Interface Men被Human,Student和Employee实现,因为这三个类型都实现了这两个方法
type Mem interface {
	Sayhi()
	Sing(lyrics string)
}
//Interface YoungChap被Student实现
type YoungChap interface {
	Sayhi()
	Sing(lyrics string)
	BorrowMoney(amount float32)
}
//Interface ElderlyGent被Employee实现
type ElderlyGent interface {
	Sayhi()
	Sing(lyrics string)
	SpendSalary(amount float32)
}

//定义Human的String方法（自定义对象的输出方式）
func (h Human) String() string {  //实现String方法
	return "❰"+h.name+" - "+strconv.Itoa(h.age)+" years -  ✆ " +h.phone+"❱"
}


func main() {
	mike := Student{Human{"Mike",25,"1234"},"beida",0.00}
	paul := Student{Human{"Paul",26,"5678"},"qinghua",100}
	sam := Employee{Human{"Sam",36,"8901"},"JD",1000}
	tom := Employee{Human{"Tom",37,"5555"},"baidu",5000}
	//定义Men类型的变量
	var i Mem

	//i能存储Student，也能存Employee
	i = mike
	fmt.Println("This is mike,a Student:")
	i.Sayhi()
	i.Sing("我爱你中国..")
	fmt.Println("定义的interface变量里可以存储实现这个interface的任意类型的对象")

	//将interface作为类型存放在slice里
	x := make([]Mem,3)
	x[0],x[1],x[2] = paul,sam,tom
	for _,value := range x {
		value.Sayhi() //每个对象调用Sayhi方法
	}

	//空interface--可以存储任意类型的数值
	var a interface{}
	var r int = 5
	s := "hello world!"
	a = r
	a = s  //空接口可以存储任意类型的数值
	fmt.Println(a)


	//任何实现了String方法的类型都能作为参数被fmt.Println调用，如果没实现String方法，fmt将以默认方式输出
	abc := Human{"Abc",39,"11111"}
	fmt.Println("This Human is: ", abc) //This Human is:  ❰Abc - 39 years -  ✆ 11111❱
}