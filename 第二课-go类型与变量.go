package main

import (
	"fmt"
	"math"
)


/* 以块的形式声明变量  */
var (
	result bool  /* bool类型，取值范围：true和false */
	str string   /* 字符串类型  */
	num int   /* 整数类型，根据系统平台识别是32还是64位  */
	 /* int/uint, int8/uint8, int16/uint16等等类似，只是位数不同，取值范围不同  */
	num2 int8   /*  int8类型，占一字节。取值范围：-128-127 */
	k byte   /* uint8类型的别名；byte = uint8  */
	runes rune   /* int32类型的别名；rune = int32  */
	floats float32   /* float32类型，精确到7位小数；float64精确到15位小数  */
)
/* 可以自定义一些类型别名  */
type (
	byte  int8
	rune int32
	文本  string  /* 定义了string类型别名是 文本 */
)

func main()  {
	var a 文本  /*  引用了上面定义的“文本”的类型，其实真正的类型是string类型  */
	a = "中文类型名"  /* 先声明，后赋值 */
	var b int = 123  /* 变量声明及赋值  */
	var c,d float32 = 3.14,-9.2222  /* 并行赋值 */
	var e,_,g string= "q","e","e" /* _的变量被忽略 */
	int64max := math.MaxInt64 /* 变量声明及赋值 --最简形式 */
	w := 32
	var ww = float32(w)  /* 类型转换 */
	fmt.Println(result)  /* 变量被声明后都会有默认值；值类型(int,float)的默认值是0，布尔值类型为false，字符串类型为空  */
	fmt.Println(str)
	fmt.Println(num)
	fmt.Println(int64max) /* 查看int64类型的最大值 */
	fmt.Println(a,b,c,d,e,g)
	fmt.Println(ww)


}

