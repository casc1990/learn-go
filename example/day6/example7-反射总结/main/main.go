package main

import (
	"fmt"
	"reflect"
)

//reflect中有两个重要的类型：reflect.Type，reflect.Value
//reflect中还有2个重要的函数：reflect.TypeOf(i interface{});reflect.ValueOf(i interface{})

//reflect.TypeOf(i interface{})；返回reflect.Type类型，获取对象的类别，并可以根据这个函数获取底层的数据类型
//
type Foo struct {
	X string
	Y int
}

func (p Foo) Do() {
	fmt.Printf("X is %s,Y is %d\n",p.X,p.Y)
}

func ReflectTypeOf() {
	//获取结构体时，显示的就是定义的类别了
	var foo Foo
	fmt.Println(reflect.TypeOf(foo)) //main.Foo
	fmt.Println(reflect.TypeOf(foo).Name()) //Foo 返回结构体的名字

	fmt.Println("reflect.Type()的方法")
	/*
	type Type interface {
  	Align() int
  	FieldAlign() int
  	Method(int) Method  //
  	MethodByName(string) (Method, bool) //
  	NumMethod() int  //
  	Name() string  //
  	String() string
  	Elem() Type //
  	Field(i int) StructField  //
  	FieldByName(name string) (StructField, bool)
  	Len() int
  	.....
}
	 */

	//NumField()有几个字段；typ.Field(i)指定的某个字段
	typ := reflect.TypeOf(foo)
	for i :=0;i< typ.NumField();i++ {
		field := typ.Field(i)
		fmt.Println(field,"---")
		fmt.Printf("%s type is :%s\n",field.Name,field.Type) //字段名，字段底层数据类型
	}
	field2,_ := typ.FieldByName("X")  //等价于typ.Field(0)
	fmt.Println(field2.Name) //X

	//method相关操作
	fmt.Println(typ.NumMethod())  //1  Foo方法的个数
	m := typ.Method(0)
	fmt.Println(m.Name) //Do  方法的名字
	fmt.Println(m.Type) //func(main.Foo)

	//Kind相关(Kind方法Type和Value都有，它返回的是对象的基本类型，例如int,bool,slice等)
	fmt.Println(typ.Kind()) //struct 结构体类型
	var f2  = &Foo{}
	typ2 := reflect.TypeOf(f2)
	fmt.Println(typ2.Kind())  //ptr 指针类型
}

func ReflectValueOf() {
	//reflect.ValueOf()的返回值类型为reflect.Value,它实现了interface{}参数到reflect.Value的反射
	var i int = 123
	var f = Foo{"abc", 123}
	var s = "abc"
	fmt.Println(reflect.ValueOf(i))  //123
	fmt.Println(reflect.ValueOf(f))  //{abc 123}
	fmt.Println(reflect.ValueOf(s))  //abc

	//reflact.Value对象可以通过调用Interface()方法，再反射回interface{}对象
	//				reflect.ValueOf()                        Interface()
	//interface{} ---------------------> reflect.Value -------------------> interface{}
	var x int = 123
	fmt.Println(reflect.ValueOf(x).Interface()) //123

	var f2 = Foo{"abc", 123}
	fmt.Println(f2) //{abc 123}
	fmt.Println(reflect.ValueOf(f2).Interface() == f)  //true
	fmt.Println(reflect.ValueOf(f2).Interface())  //{abc 123}

	//value的Field方法(和Type的Filed方法不一样，Type.Field()返回的是StructFiled对象，有Name,Type等属性，Value.Field()返回的还是一个Value对象)
	val := reflect.ValueOf(f2)
	fmt.Println(val.FieldByName("X")) //abc 返回字段的值
	typ := reflect.TypeOf(f2)
	fmt.Println(typ.FieldByName("X")) //{X  string  0 [0] false} true 返回结构体对象

	//混合使用
	rv := reflect.ValueOf(f2)
	rt := reflect.TypeOf(f2)
	for i:=0;i < rv.NumField();i++ {
		fv := rv.Field(i)
		ft := rt.Field(i)
		fmt.Printf("%s type is :%s,value is %v\n",ft.Name,fv.Type(),fv.Interface())
	}

	//设置value的值
	var k int = 123
	fv := reflect.ValueOf(k)
	fmt.Println(fv.CanSet()) //false  值拷贝是不能直接修改的
	fe := reflect.ValueOf(&k).Elem() //必须是指针的Value才能调用Elem
	fmt.Println(fe.CanSet()) //true
	fe.SetInt(456)
	fmt.Println(k) //456

	//method方法
	var foo2 = Foo{"abc",123}
	reflect.ValueOf(foo2).MethodByName("Do").Call([]reflect.Value{})




}

func main() {
	var i int = 123
	var f float32 = 1.23
	var l []string = []string{"a","b","c"}
	var s string = "abc"
	//获取对象的类型（一般类型确实可以获取到）
	fmt.Println(reflect.TypeOf(i))  //int
	fmt.Println(reflect.TypeOf(f))  //float32
	fmt.Println(reflect.TypeOf(l))  //[]string
	fmt.Println(reflect.TypeOf(s).Name()) //string
	fmt.Println(reflect.TypeOf(s).String())  //string

	//TypeOf类型方法
	ReflectTypeOf()

	//ValueOf类型方法
	ReflectValueOf()



}