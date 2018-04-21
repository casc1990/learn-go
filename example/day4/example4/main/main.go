package main

import "fmt"

func fab(num int) {
	var a []uint64
	a = make([]uint64,num)
	a[0] = 1
	a[1] = 1
	for i :=2; i < num;i ++ {
		a[i] = a[i-1] + a[i-2]
	}

	for _,v := range a {
		fmt.Println(v)
	}
}

func main() {

	fab(20)

	//数组的初始化
	var age0 [5]int = [5]int{1,2,3,4,5}
	var age1 = [5]int{1,2,3,4,5}  //声明及赋值
	var age3 = [...]int{1,2,3,4,5}
	var age4 = [5]string{3:"hello",4:"world"} //下标赋值
	fmt.Println(age0,age1,age3,age4)

	//多维数组及遍历
	var age5 = [...][4]int{{1,2,3,4},{5,6,7,8}}
	for row,v1 := range age5 {
		for col,v2 := range v1 {
			fmt.Printf("row:%d,col:%d,value:%d\n",row,col,v2)
		}
	}
}
