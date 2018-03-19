package main

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
func (u User) Hello(name string) {
	fmt.Println("hello!",u.Name)
}

func (u User) Eat() {
	fmt.Printf("%s is eatting",u.Name)
}


func Info(o interface{}) {

	types := reflect.TypeOf(o)
	fmt.Println("传递过来的参数类似是:",types.Name())
	values := reflect.ValueOf(o)
	fmt.Println("传递过来的参数第一个字段的值是:",values.Field(0))
	fields := types.Field(1)  //获取字段名
	fmt.Printf("第一个字段名称是:%s,字段类型是:%v,字段值是:%s\n",fields.Name,fields.Type,values.Field(1))


	for i := 0;i < types.NumField(); i++ {
		field := types.Field(i)
		value := values.Field(i)
		fmt.Println(field.Name,field.Type,value)
	}


	//获取方法(被获取的对象方法一定要是可导入的，就是大写字母开头的)
	for i :=0;i< types.NumMethod(); i++ {
		m := types.Method(i) //取每个方法
		fmt.Println("Method:")
		fmt.Printf("smethod list: %s\n",m.Name)
	}

	//通过反射来进行方法的调用（Value的Call方法的参数是一个Value的slice，返回值也是一个Value的slice）
	mv := values.MethodByName("Hello")
	args := []reflect.Value{reflect.ValueOf("lisi")}
	mv.Call(args) //调用hello方法
	mv2 := values.MethodByName("Eat")
	mv2.Call(nil) //Eat方法不需要参数。直接传递nil即可


}

func main() {
	u := User{Id:20,Name:"lisi",Age:21}
	u.Hello("sss")
	Info(u)
}