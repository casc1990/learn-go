package go基础语法

import (
	"fmt"
	"math"   //科学计算的包
	"strconv"  //字符串转换包
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
	var w float32 = 32.45
	var ww = int(w)  /* 类型转换,转换只能发生在两种相互兼容的类型之间 */
	/* 类型转换的格式：<ValueA> [:]= <TypeOfValueA>(<ValueB>) */
	/* 浮点型和int型可以相互转换，int型可以转换为string类型，但是string类型不可用转换为int型 */
	var tt int = 65
	var ttt string = strconv.Itoa(tt) /* strconv.Itoa函数可以把int转换为字符串类型 */
	var tttt,_ = strconv.Atoi(ttt) /* strconv.Atoi函数把字符串类型转换为int类型，注意此函数返回2个值  */
	fmt.Println(string(tt)) /* 输出结果：A，因为A的ascii值是65，计算机中存储的任何东西本质上都是数字  */
	fmt.Println(ttt)  /* 输出结果："65". */
	fmt.Println(tttt)  /* 输出结果：65. */
	fmt.Println(result)  /* 变量被声明后都会有默认值；值类型(int,float)的默认值是0，布尔值类型为false，字符串类型为空  */
	fmt.Println(str)  //输出：空
	fmt.Println(num)  //输出：0
	fmt.Println(int64max) /* 查看int64类型的最大值 */
	fmt.Println(a,b,c,d,e,g)
	fmt.Println(ww)  //输出：32
}

