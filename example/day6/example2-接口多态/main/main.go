package main

import (
	"fmt"
	"math/rand"

	"sort"
)

//多态：一种事物的多种形态，都可以按照统一的接口进行操作
//多态一般是用来定义一些规范，只要某个类型实现了所定义的规范就可以使用这个接口
//sort.Sort函数的实现如下：(只要我们的数据类型实现了Len()、Less(i, j int)、Swap(i, j int)方法，我们就可以使用Sort排序)
/*
func Sort(data Interface) {
	n := data.Len()
	quickSort(data, 0, n, maxDepth(n))

type Interface interface {
	// Len is the number of elements in the collection.
	Len() int
	// Less reports whether the element with
	// index i should sort before the element with index j.
	Less(i, j int) bool
	// Swap swaps the elements with indexes i and j.
	Swap(i, j int)
}
 */

type  Student struct {
	Name string
	Id string
	Age int
}

//定义一个Student类型的切片存student
type StudentArray []Student

//实现Sort函数中的Len方法
func (p StudentArray) Len() int {
	return len(p)
}

//从小到大排序
func (p StudentArray) Less(i,j int) bool {
	return p[i].Name < p[j].Name
}

//交换
func (p StudentArray) Swap(i,j int) {
	p[i],p[j] = p[j],p[i]
}


func main() {
	var stus StudentArray
	//生成slice
	for i:=0;i<10;i++ {
		stu := Student{
			Name: fmt.Sprintf("stu%d",rand.Intn(100)),
			Id: fmt.Sprintf("110%d",rand.Int()),
			Age : rand.Intn(50),
		}
		stus = append(stus,stu)
	}

	for _,v := range stus {
		fmt.Println(v)
	}
	fmt.Println("------华丽的分隔符-----")

	//stus是StudentArray实例.StudentArray实现了Sort函数参数Interface的所有方法，所以他也实现了这个接口
	sort.Sort(stus)
	//再次查看
	for _,v := range stus {
		fmt.Println(v)
	}
}
