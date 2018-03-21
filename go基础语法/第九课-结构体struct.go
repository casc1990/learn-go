package go基础语法

//struct：用来描述一组类型或者字段的集合,集合里的类型或者字段叫做属性（struct也是传值的）
//struct匿名字段：只提供类型，而不写字段名的struct，匿名字段可以包括所有的类型。如果匿名字段是struct类型，那么这个struct所拥有的全部字段都
// 	被隐式地引入了当前定义的这个struct
//struct匿名字段：能够实现字段的继承。如果被继承的字段里也有外层的同名字段是，可以通过匿名字段名区分访问那个字段

import (
	"fmt"
)

/* 定义struct，这个struct有2个属性：name和age  */
type person struct {
	name string
	age  int
}

// P使用了这个结构体类型
var P person  //P现在就是person类型的变量了

//比较两个人年龄，返回年龄大的那个人，并返回年龄差
func Older(p1,p2 person) ( person, int ) {
	if p1.age > p2.age {  // 比较p1和p2这两个人的年龄
		return p1, p1.age - p2.age
	}
	return p2, p2.age - p1.age
}

//匿名字段--继承
type Skills  []string  //自定义类型
type Human struct {  //定义人类这个结构体
	name string
	age int
	weight  int
	phone int
}
type Student struct {
	Human  //匿名字段，默认Student就包含了Human的所有字段
	Skills //匿名字段，自定义的类型string slice
	int // 内置类型作为匿名字段
	speciality string  //定义新字段：专业
	phone int  //匿名字段的属性和外层的属性同名了（通过匿名字段来区分访问的属性）
}

func main()  {
	P.name = "zhangsan"  //给结构体P的name属性赋值为zhangsan
	P.age = 26  //同上
	fmt.Printf("The person is name is %s", P.name) /* 访问name的值  */

	// struct赋值方式
	var tom person
	tom.name,tom.age = "Tom",18  //初始化赋值
	bob := person{age:19,name:"Bob"} /* 两个字段都写清楚的初始化赋值 */
	paul := person{"Paul",20} //按struct定义顺序赋值
	tb_Older,tb_diff := Older(tom,bob)  //调用Older函数比较tom和bob谁大，返回谁大（结构体），大多少（int）
	bp_Older,bp_diff := Older(bob,paul)
	fmt.Printf("of %s and %s, %s is older by %d years\n",tom.name,bob.name,tb_Older.name,tb_diff)  //tb_Older.name：取结构体的name字段
	fmt.Printf("of %s and %s, %s is older by %d years\n",bob.name,paul.name,bp_Older.name,bp_diff)

	//匿名字段
	mark := Student{Human:Human{"Mark",25,120,123456789},speciality:"computer Science"}
	fmt.Println("His name is",mark.name) //直接访问
	fmt.Println("His speciality is ", mark.speciality)
	mark.speciality = "AI" //修改字段
	fmt.Println("Mark changed his speciality")
	fmt.Println("His speciality is ", mark.speciality)
	fmt.Println("Mark become old")
	mark.age = 46 // 修改他的年龄信息
	fmt.Println("His age is", mark.age)
	//通过匿名字段访问或者修改
	mark.Human = Human{"AAA",55,221,123456789}  // 通过匿名字段修改
	mark.Human.age -= 1  //年龄减一岁
	fmt.Println("His age is", mark.age)

	//其他类型的匿名字段
	mark.Skills = []string{"anatomy"}
	mark.Skills = append(mark.Skills,"python","golang")
	fmt.Println("Her skills now is ",mark.Skills)  //Her skills now is  [anatomy python golang]

	//匿名字段的属性和外层的属性同名了
	mark.phone = 12345678 //操作的是外层的phone字段
	mark.Human.phone = 222222222  //操作的是内层的phone字段
	fmt.Println("mark's work phone is:",mark.phone) //mark's work phone is: 12345678
	fmt.Println("mark's personal phone is:",mark.Human.phone)  //mark's personal phone is: 222222222



}