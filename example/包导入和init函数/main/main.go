package main

//导入包的执行顺序是：从后往前推，依次执行被导入包的全局变量，init函数，然后才会执行主函数main
import (
	"os"
	"fmt"
	"learn-go/example/包导入和init函数/calc"
	_ "learn-go/example/包导入和init函数/test"  //在导出的包前加上_，只会导出，不引用，但是会执行包里的全局变量和init函数
)

func init() {
	fmt.Println("init func会先于main函数执行,init函数主要做初始化....")
}

func main() {
	fmt.Println(os.Getenv("GOPATH"))
	fmt.Println("calc中Name:",calc.Name)
	fmt.Println("calc中Age:",calc.Age)  //访问外部包里的变量值
	calc.Add(4,5) //4 + 5 = 9;调用外部包里的函数
	fmt.Println("外部包里大写字母开头的变量或者函数才能被调用")

}
