package main

import "fmt"

//判断是否是完数（一个数如果恰好等于它的约数之和）
func PerfectNum(n int)  {
	var result,num int
	for i :=1;i < n;i++ {  //找出所有的约数
		if n % i == 0 {
			result += i
		}
		}
	if result == n {  //判断约数之和是否等于这个数本身
		fmt.Println("找到完数:",n)
		num += 1
	}
}

func main() {
	//九九乘法表
	for i :=1;i < 10;i++ {  //共九行
		for j :=0;j < i; j++ {  //每行打印j个字符
			fmt.Printf("%d * %d = %d ",i,j+1,i*(j+1))
			}
			fmt.Println() //换行
		}

	var min,max int
	fmt.Scanf("%d%d",&min,&max)
	for i :=min;i < max;i++ {
		PerfectNum(i)
	}


	}
