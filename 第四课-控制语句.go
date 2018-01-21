package main

import "fmt"


/* 指针使用流程：
	1.定义指针变量。
	2.为指针变量赋值。
	3.访问指针变量中指向地址的值。   */

/*
 break 语句 : 经常用于中断当前 for 循环或跳出 switch 语句
 continue 语句: 跳过当前循环的剩余语句，然后继续进行下一轮循环。
 goto 语句 : 将控制转移到被标记的语句。
*/


func main() {
	var a int = 10
	var b int = 20
	fmt.Println("变量的地址: %x", &a)  /* 操作符”&”取变量地址  */
	/*  定义指针： var var_name *var-type */
	var ip *int        /* 定义了指向整型的指针  */
	ip = &a
	fmt.Println("ip变量储存的指针地址: %x", ip)
	 /* 使用指针访问值 */
   	fmt.Printf("*ip变量的值: %d\n", *ip )  /* 使用”*”通过指针间接访问目标对象 */
	/*
	判断语句if
		条件表达式没有括号
		支持一个初始化表达式（可以是并行方式）
		左大括号必须和条件语句或else在同一行
		支持单行模式
		初始化语句中的变量为block级别，同时隐藏外部同名变量
	 */
	if a > 10 {
		fmt.Println("a > 10")
	}else {
		fmt.Println("a < 10")
	}
	/* if嵌套 */
	if a == 10 {
		fmt.Println("a等于10")
		if b == 20 {
			fmt.Println("b等于20")
		}else {
			fmt.Println("b不等于20")
		}
	}else  {
		fmt.Println("a不等于10")
	}
	fmt.Printf("a值为:%d\n" ,a)
	fmt.Printf("b值为:%d\n", b)
	/* if支持初始化表达式 */
	if c,d := 22,44;d>c {
		fmt.Printf("d大于c")
	}

	/*  for循环语句*/
	/* 第一种形式，不加参数 */
	for {
		a ++
		if a > 15 {
			break
		}
		fmt.Println("a的值为:",a)
		}
	fmt.Println("Over")
	/* for的第二种形式  for 条件表达式 { }*/
	for a < b {
		a ++
		fmt.Println("a的值为:",a)
	}
	/* for的第三种形式 */
	for a := 0; a < 10; a++ {
      		fmt.Printf("a 的值为: %d\n", a)
   	}
	/*
	switch 语句用于基于不同条件执行不同动作，每一个 case 分支都是唯一的，从上直下逐一测试，直到匹配为止。。
	switch 语句执行的过程从上至下，直到找到匹配项，匹配项后面也不需要再加break

		switch var1 {
    			case val1:
        			...
    			case val2:
        			...
    			default:
        			...
		}

	  */
	/* 定义局部变量 */
   	var grade string = "B"
   	var marks int = 90
	/* switch第一种形式  */
   	switch marks {
      		case 90: grade = "A"
      		case 80: grade = "B"
      		case 50,60,70 : grade = "C"
      		default: grade = "D"
		}
	/* switch第二种形式  */
   	switch {
      		case grade == "A" :
         		fmt.Printf("优秀!\n" )
      		case grade == "B", grade == "C" :
         		fmt.Printf("良好\n" )
      		case grade == "D" :
         		fmt.Printf("及格\n" )
      		case grade == "F":
         		fmt.Printf("不及格\n" )
      		default:
         		fmt.Printf("差\n" );
		}
   	fmt.Printf("你的等级是 %s\n", grade );
	/* switch第三种形式(跟赋值语句)   switch marks a := 1; {  } */

	/* goto 跳转 */
	for i := 1; i < 10;i++ {
		if i > 3 {
			goto label1  /* 去到 label1处 */
		}
	}
	label1:
		fmt.Println("ok")
}

