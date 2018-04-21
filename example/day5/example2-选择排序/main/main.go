package main

import "fmt"

//选择排序的原理是，对给定的数组进行多次遍历，每次均找出最大的一个值的索引。

//选择排序
func ssort() {
	ar := []int{9,2,4,8,0,5,3,7}
	num := len(ar)

	for i :=0;i <num;i ++ { //循环数组,比较len(num)次
		var maxIndex = 0
		for j := 1;j <num-i;j++ { //从第二个数开始循环，每次都可以找到一个数，所有循环次数为num-1
			if ar[j] > ar[maxIndex] {  //找出最大数的下标
				maxIndex = j
			}
		}
		ar[num-i-1],ar[maxIndex] = ar[maxIndex],ar[num-i-1] //交换交换
	}
	fmt.Println(ar)
}

func main() {
	ssort()
}

