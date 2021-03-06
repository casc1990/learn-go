package main

import "fmt"

//冒泡排序的原理是，对给定的数组进行多次遍历，每次均比较相邻的两个数，如果前一个比后一个大，则交换这两个数。经过第一次遍历之后，
//最大的数就在最右侧了；第二次遍历之后，第二大的数就在右数第二个位置了；以此类推。

//冒泡排序
func bsort() {
	ar := []int{9,2,4,8,0,5,3,7}
	num := len(ar)

	//从小到大排序
	for i := 0;i<num;i++ { //循环数组,比较len(num)次
		for j := 1;j<num-i;j++ {  //从第二个数开始循环，因为每次都可以冒出最大或者最小的数，所有循环num-i次
			if ar[j] < ar[j-1] {  //比较相邻两个元素，如果第一个数小于第二个数
				ar[j],ar[j-1] = ar[j-1],ar[j]  //交换位置，大的数就会冒出来
			}
		}
	}
	fmt.Println(ar)
}

func main() {
	bsort()
}
