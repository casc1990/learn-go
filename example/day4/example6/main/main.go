package main

//sort函数只能对引用类型排序，因为值类型，你修改不了源数据
//sort.Ints对整数进行排序， sort.Strings对字符串进行排序, sort.Float64s和sort.Float32s对浮点型排序。
//sort.SearchInts:查询整数在数组中的下标（排序后的），sort.SearchInts和sort.SearchInts类似

import (
	"sort"
	"fmt"
)

func testIntSort() {
	var a = [...]int{1,3,8,2,99,66,486,321}
	sort.Ints(a[:]) //sort函数只能对引用类型排序，因为值类型，你修改不了源数据
	fmt.Println(a) //[1 2 3 8 66 99 321 486] 按数字的升序排序
}

func testStringSort() {
	var a = [...]string{"abc","a","A","ac","b"}
	sort.Strings(a[:]) //string排序
	fmt.Println(a) //[A a abc ac b] 按先大写，后小写，并按位排序
}


func testFloat64Sort() {
	var a= [...]float64{1.33,0.2892,99,77,77.0,23.10,23.1}
	sort.Float64s(a[:]) //float64排序
	for index,value := range a {
		fmt.Printf("下标%d的value = %f\n",index,value)  //按浮点数升序排序，数字相同取精度
	}
}

func TestSearchInt() {
	var a = [...]int{1,3,8,2,99,66,486,321}
	index := sort.SearchInts(a[:],2) //查询某个数字在数组中的下标（排序后的下标）
	fmt.Println(index)  //1
}


func main() {
	testIntSort()
	testStringSort()
	testFloat64Sort()
	TestSearchInt()
}
