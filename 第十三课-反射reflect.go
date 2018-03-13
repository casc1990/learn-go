package main
//所谓反射就是能检查程序在运行时的状态
import (
	"fmt"
	"reflect"
)

//通过传递的对象，获取方法，字段等等
type User struct {
	Id int
	Name string
	Age int
}
func (u User) Hello() {
	fmt.Println("hello!",u.Name)
}
func Info(o interface{}) {
	types := reflect.TypeOf(o)
	//获取传递对象的名字
	fmt.Println("This is type:",types.Name()) //This is type: User
	//获取字段内容
	v := reflect.ValueOf(o)
	fmt.Println(v)  //{12 ok 25}
	fmt.Println("Fields:")
	for i :=0;i < types.NumField(); i++ { //types.NumField()取字段的数量
		field := types.Field(i)  //取每个字段
		value := v.Field(i).Interface()  //取每个字段的值
		fmt.Printf("%6s: %v = %v\n",field.Name,field.Type,value)
	}
	//获取方法
	for i :=0;i< types.NumMethod(); i++ {
		m := types.Method(i) //取每个方法
		fmt.Println("Method:")
		fmt.Printf("smethod list: %s",m.Name)
	}

}

func main()  {
	var x float64 = 3.4
	fmt.Println("type:",reflect.TypeOf(x))  //type: float32 打印对象的类型
	fmt.Println("value:",reflect.ValueOf(x)) //value: 3.4 打印对象的值
	v := reflect.ValueOf(x) //被reflect处理过的对象有许多方法
	fmt.Println("value type:",v.Type())  //value type: float64 (Value有一个Type方法返回reflect.Value的Type)
	fmt.Println("kind is float64:",v.Kind() == reflect.Float64)  //kind is float64: true; (Type和Value都有Kind方法返回一个常量来表示类型：Uint、Float64、Slice等等)
	fmt.Println("value:", v.Float()) //value: 3.4; (Value有Int和Float的方法可以获取存储在内部的值)

	//通过传递的对象，获取方法，字段等等
	u := User{12,"ok",25}
	Info(u)
}
