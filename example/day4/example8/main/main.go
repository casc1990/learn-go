package main

import (
	"sort"
	"fmt"
)

//map排序（需要自己实现：把key追加到slice里，对slice排序，并打印map[排序后的key]）
func testMapSort() {
	var a map[int]int
	a = make(map[int]int,5)
	a[8] = 8
	a[3] = 3
	a[2] = 2
	a[5] = 5

	var b []int
	for k,_ := range a {
		b = append(b,k)  //将key追加到slice里
	}
	// 对key排序，并循环字典，打印排序后key所对应的key和value
	sort.Ints(b)
	for _,v := range b {
		fmt.Println(v,a[v])
	}
}

//map反转：初始化另外一个map，把key、value互换即可

func testMapreverse() {
	a := make(map[string]int)
	b := make(map[int]string)

	a["abc"] = 10
	a["efg"] = 20

	for k,v := range a {
		b[v] = k  //key、value互换
	}
	fmt.Println("字典a:",a)
	fmt.Println("字典b:",b)
}


func main() {
	testMapSort()
	testMapreverse()
}
