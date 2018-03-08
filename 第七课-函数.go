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



}
