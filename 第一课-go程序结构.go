/* 当前代码的包名称,package必须放在行首  */
package main

/*  导入其他的包,d导入未调用，也会报编译异常  */
import std "fmt"   /*  设置导入包别名 */
/* import . "fmt"     省略调用 */

/*  常量的定义  */
const PI  = 3.14
/*  全局变量的声明与赋值    */
var name string = "zhangsan"  /*  根据规定，首字母大写或者全大写的变量、常量、函数方法等可以别外部调用，小写的不可别外部调用  */
/*  一般类型声明    */
type newtype string

/*  结构体的声明    */
type gopher struct {}

/*  接口的声明    */
type golang interface {}
/*  由 main函数作为程序入口点启动   */
func main()  {
	std.Println("你好，世界！！")
}