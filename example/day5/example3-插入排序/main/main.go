package main

import "fmt"

//插入排序:选择排序（Selection sort）是一种简单直观的排序算法。它的工作原理是每一次从
// 待排序的数据元素中选出最小（或最大）的一个元素，存放在序列的起始位置，直到全部待排序的数据元素排完

//插入排序（将要排序的数据类型，先分成两个序列，一边存排序后的有序序列，一边存待排序的无序序列，循环无序序列，依次插入到有序序列中）
func iSort() {
	b := []int{8,7,5,3,4,10,15}
	num := len(b)

	//循环无序序列，随着有序序列元素的增多，无序序列最多需要比较len(num)-1次
	for i := 1;i < num;i++ {
		//循环有序序列，准备比较
		for j :=i;j>0;j-- {
			//比较并排序（后面的数比前面的大，说明是排序好的，退出）
			if b[j] > b[j-1] {
				break  //有序就退出
			}
			//交换
			b[j],b[j-1] = b[j-1],b[j]
		}
	}
	fmt.Println(b)
}

func main() {
	iSort()
}