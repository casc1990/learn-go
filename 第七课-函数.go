package main
/* 函数是go里面的核心设计，使用func关键字声明
  func funcname(input1 type,input2 type2..) (output1 type1,output2,type2..) {
  	//这里是处理逻辑代码
    	//返回多个值
    	return value1, value2
  }
 */

import (
	"fmt"
	"os"
)
/*  求最大值函数，接收两个参数，返回值是int类型 */
func max(a,b int) int {
	if a > b {
		return a
	}
	return b
}
/* 返回A+B 和 A*B */
func SumAndProduct(A,B int)(int,int) {
	return A+B,A*B
	/* fmt.Print("会不会执行呢?") */
}

/* 在初始化函数里，命名返回值变量.官方不太建议这样做  */
func SumAndProducts(A,B int) (Add int, Multiplied int) {  /* 定义了返回值变量名称  */
	Add = A + B
	Multiplied = A * B
	return  /* 相当于 return Add  Multiplied*/
}

/*  不固定长度参数。表达式：func myfunc(arg ...int)  */
func ArgFunc(args ...int) {  /* 如果函数没返回值，可以省略返回值定义字段  */
	fmt.Println(args)  /* 输出：[1 2 3 4 5 6]  */
	for _,n := range args {  /*  循环参数列表并打印  */
		fmt.Printf("And the number is %d\n",n)
	}
}

/* 值传递和指针传递
 值传递：传参数值到函数里时，只是对这个值做了一份copy，在函数里修改参数值，实际参数不会发生变化
 指针传递：因为指针指向了实参的内存地址，修改指针指向的内存地址，实参就改变了
 */
//简单的一个函数，实现了参数+1的操作
func add1(a *int) int {
	*a = *a +1  /* 修改指针指向的内存地址  */
	return *a /* *<指针变量>：表示取指针指向的变量的值 */
}

/* defer: 延迟语句，在函数返回或者退出之前的一些收尾工作，会按照逆序执行
示例：（程序退出之前会关闭文件）
 func ReadWrite() bool {
    file.Open("file")
    defer file.Close()
    if failureX {
        return false
    }
    if failureY {
        return false
    }
    return true
}
 */

/* 函数作为值、类型: 函数也是一种变量，可以通过type来定义
定义格式： type typeName func(input1 inputType1 , input2 inputType2 [, ...]) (result1 resultType1 [, ...])
 */
type testInt func(int) bool  // 声明一个函数类型（这个函数接收int类型参数，返回值为bool类型）
func isOdd(integer int) bool {
	if integer %2 == 0 { /* 判断是否是奇数 */
		return false
	}
	return true
}

func isEven(integer int) bool {
	if integer %2 ==0 {  /* 判断是否是偶数 */
		return true
	}
	return false
}
// 声明的函数类型在这个地方当做了一个参数,这个参数要符合：1.是一个函数，2，参数为int，3，返回值为bool类型
func filter(slice []int,f testInt) []int {
	var result []int
	for _,value := range slice {
		if f(value) {  /* value传入函数，如果返回值为bool类型，追加到result的slice里 */
			result = append(result,value)
		}
	}
	return result
}

/*  Panic和Recover
 panic：由程序运行出错或者直接调用panic产生，panic会中断原有控制流程，如果有defer函数，会跳到执行defer函数，否则异常退出
 recover：恢复异常，仅在defer函数中有效；正常执行过程中，调用recover返回nil，panic模式下，调用recover可以捕获到panic的输入值，并且恢复正常的执行
 */
func myPainc() {
	fmt.Println("panic开始之前...\n")
	panic("程序panic啦...\n")   /*  程序panic之后调用了defer函数，panic之后的程序不会再执行了  */
	fmt.Println("panic之后不会执行..\n")
}


/* main函数和init函数（不能有任何的参数和返回值）
main：main函数只能存在于package main里，而且只能有一个
init: init可以存在所有的package里，而且可以都多个，init函数先于main函数执行
import: 包只会导入一次，如果要导入的包里还导入了其他包，那么会先将其它包导入进来，然后再对这些包中的包级常量和变量进行初始化，接着执行init函数（如果有的话），依次类推
*/
var user = os.Getenv("USER")  /*  获取当前用户 */
 func init() {   //init函数会被首先执行。。
	 if user == "" {
		fmt.Println("no value for user")
	 }
 }


func main() {
	x := 3
	y := 4
	z := -5

	max_xy := max(x,y) /*  调用函数max，传递x，y参数 */
	max_xz := max(x,z)

	fmt.Printf("max(%d,%d) = %d\n",x,y,max_xy) /* max(3,4) = 4 */
	fmt.Printf("max(%d,%d) = %d\n",x,z,max_xz) /* max(3,-5) = 3 */
	fmt.Printf("max(%d,%d) = %d\n",y,z,max(y,z)) /* max(4,-5) = 4; 也可以直接调用 */

	xPLUSY, xTIMESY := SumAndProduct(x,y)  /* 函数多返回值  */
	fmt.Printf("%d + %d = %d\n", x, y, xPLUSY)
	fmt.Printf("%d * %d = %d\n", x, y, xTIMESY)

	ArgFunc(1,2,3,4,5,6)  /* 没返回值的函数直接调用  */
	ArgFunc(x,y,z) /* 同上 */

	fmt.Println("x = ",x)  // 应该输出 "x = 3"
	x1 := add1(&x) /*  调用add1函数，传x的地址 */
	fmt.Println("x + 1 = ",x1)  // 应该输出 "x+1 = 4"
	fmt.Println("x = ",x)  // 应该输出 "x = 4"

	/* defer 执行顺序 （采用后进先出模式） */
	for i := 0; i < 5; i++ {
		defer fmt.Printf("%d\n", i) /* 输出：4 3 2 1 0 */
	}

	/*  函数作为值、类型 */
	slice := []int {1,2,3,4,5,7}
	fmt.Println("slice =",slice)
	odd := filter(slice,isOdd)  // 函数当做值来传递了
	fmt.Println("Odd elements of slice are:",odd)  //输出：Odd elements of slice are: [1 3 5 7]
	even := filter(slice, isEven)  // 函数当做值来传递了
    	fmt.Println("Even elements of slice are: ", even) //输出：Even elements of slice are:  [2 4]

	/* defer函数 */
	defer func() {
		fmt.Print("defer函数用来做收尾工作，函数退出之前执行，一般配合recover捕捉异常")
		if err := recover(); err != nil {  /* recover函数保存了panic的出错信息，如果正常，panic为空，否则，为错误信息  */
			fmt.Print(err)  /*  打印出错信息 */
		}
		fmt.Print("defer执行完成\n")
	}()

	//调用自定义的painc函数，触发异常，程序会走到defer函数
	myPainc()   /* 调用自定义的panic函数  */


}

