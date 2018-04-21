package main

import "fmt"

//切片：切片是数组的一个引用，因此切片是引用类型(一般是对数组做切片或者自定义切片类型)
//切片的长度可以改变，是一个可变长的数组
// cap()函数可以求出slice的最大容量，len()函数可以求出切片的长度
//切片的定义: var 变量名 []类型 ：var str []string
//初始化: var slice []int = arr[start:end]

func testslice() {
	var a [5]int = [...]int{6,7,8,9,0}
	e := a[1:]
	fmt.Printf("e addr=%p,a[1] addr=%p",e,&a[1]) //切片指向了同一个底层数组，所以内存地址一致
	e[1] = 100
	fmt.Printf("before a=%d",a) //before a=[6 7 100 9 0]  因为是指针类型，其实是修改了a的值

	e = append(e,10)
	e = append(e,11)
	e = append(e,12)
	e = append(e,13)
	e = append(e,14)
	fmt.Printf("e addr=%p,a[1] addr=%p",e,&a[1]) //再次打印，内存不一致了，因为slice的容量用完了，重写分配了一块内存地址
	e[1] = 1000
	fmt.Printf("after a=%d",a) //after a=[6 7 100 9 0] 不同了内存地址了，所以就不会影响之前的数组了
	fmt.Println(e) //[7 1000 9 0 10 11 12 13 14] //新数组的值被修改了
}

//字符串本来是不能修改的，通过slice就可以修改了
func testModifyString() {
	s := "hell0,中国"
	s1 := []rune(s) //将字符串转换成rune或者bytes类型的切片
	s1[6] = '祖'
	str := string(s1)
	fmt.Println(str) //hell0,祖国

}

func main() {
	var s []int   //定义切片类型
	var arr = [5]int{1,2,3,4}  //定义数组类型
	fmt.Println(arr)
	s = arr[1:3]  //切片初始化
	s1 := arr[:]  //取数组所有元素
	s3 := arr[:2]  //取下标0-1的元素
	s4 := arr[:len(arr)-1]  //取第一个到倒数第二个
	fmt.Printf("%p,%p",s,&arr[1])  //切片指向的是底层的数组，所以他们内存地址都是一样的
	fmt.Println(s1,s3,s4)

	testslice()

	//切片拷贝
	w1 := []int{1,2,3,4}
	w2 := make([]int,10)
	copy(w1,w2) //将w2拷贝到w1中。
	fmt.Println(w1) //[0 0 0 0] 因为容量只有4，拷贝是从头到尾进行，又因为w2有10个元素，所以w1原有元素被覆盖，w2多余数据被丢失
	w3 := []int{1,2,3}
	w3 = append(w3,w2...)
	fmt.Println("w3=",w3) //w3= [1 2 3 0 0 0 0 0 0 0 0 0 0] 追加和copy的区别是，append只会往后加
	w3 = append(w3,4,5,6) //追加多个元素

	testModifyString()



}
