package main

import (
	"fmt"
	"math"
)

func order(n int) int {
	result := 0
	if n > 0 {
		result = n * order(n-1)  //递归函数
		return result
	}
	return 1
}

func main() {
	num := 0
	for i :=101;i <= 200;i++ {
		if (i % 2 != 0) {
			fmt.Print(i,"\n")
			num += 1
		}
	}
	fmt.Print("素数的个数:",num,"\n")

	num1 := 0
	for i :=100;i< 999;i++ {
		for j :=0;j<= 9;j++ {
			for k :=0;k <=9;k++ {
				for n :=0;n <=9;n++ {
					result1 := math.Pow(float64(n),3) + math.Pow(float64(k),3) + math.Pow(float64(j),3)

					if result1 == float64(i) {
						fmt.Println("水仙花数:",j*100+k*10+n)
						num1 += 1
					}

					}

					}
				}
		}
	fmt.Print("水仙花数的个数:",num1,"\n")

	a := math.Sqrt(9)
	b := math.Pow(2,2)
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(math.Pow(2,2))

	i := 6
	res := order(i)
	fmt.Printf("%d的乘阶是:%d",i,res)

}
