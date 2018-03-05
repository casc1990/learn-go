package main
/*
数组定义格式：var varName [n]<type> n>0
数组长度也是类型的一部分，因此具有不同长度的数组为不同类型
注意区分指向数组的指针和指针数组
数组在Go中为值类型
数组之间可以使用==或!=进行比较，但不可以使用<或>
可以使用new来创建数组，此方法返回一个指向数组的指针
Go支持多维数组
 */

import (
	"fmt"
)

func main()  {
	var arrays [4]int  /* 数组定义格式：var varName [n]<type> n>0 */
	var a [2]string   /* string类型的数组 */
	a[0] = "ab"
	a[1] = "bc"  /* 数组赋值  */
	b := [2]string{"a","b"} /* 定义并赋值，最简形式 */
	c := [3]int{1,2}   /* 元素不够使用默认值补齐 */
	d := []int{1,2,3,4,5,6}  /* 根据元素个数确定数组长度  */
	e := [20]int{9:2} /* 数组共20个元素，第10个元素（下标为9）的值为2，其他元素使用默认值 */
	f := [10]int{0:1,4:5,9:10}  /* 索引赋值，第一个元素值为1，第五个元素值为5  */

	var p *[10]int = &f  /* 指向数组的指针  定义指针：*<type>；&<变量>：取变量地址 */
	x,y := 11,22
	/* p1 := [2]*int{&x,&y} */
	var p1 [2]*int  /* 指针数组；这个指针数组有2个元素  */
	p1[0] = &x
	p1[1] = &y  /* 等同于：p1 := [2]*int{&x,&y} */

	/* 数组之间可以使用==或!=进行比较，但不可以使用<或> */
	q := [2]int{1,2}  /* [2]int和[3]int是不同类型的数组（长度不同，类型不同）  */
	w := [2]int{1,3}
	fmt.Println(w != q)  /* 输出：true  */

	t := new([2]int)
	fmt.Println(t) /* 输出: &[0 0] ; new(type)创建指向数组的指针 */

	r := [3][2]int{
		{1,2},
		{3,4},
		{5,6},
	}
	fmt.Println(r)  /* 多维数组；有3个元素，每个元素又是一个两个元素的数组 */
	fmt.Println(r[1][1]) /* 输出：4；访问多维数组 */

	/* 冒泡排序 */
	k := [...]int{3,5,2,6,9,1,0}
	fmt.Println(k)
	num := len(k)
	for i :=0;i < num;i++ {
		for j := i+1;j < num;j++ {
			if k[i] < k[j] {
				temp := k[i]
				k[i] = k[j]
				k[j] = temp
			}
		}
	}
	fmt.Println(k) /* 输出：[9 6 5 3 2 1 0] */

	fmt.Println(arrays)   /* 输出:[0 0 0 0] 数组不赋值，默认使用默认值填充 */
	fmt.Println(a)  /* 输出：[  ] */
	fmt.Println(b)  /* 输出: [a b] */
	fmt.Println(c)  /* 输出: [1 2 0]  元素不够使用默认值补齐 */
	fmt.Println(len(d))  /* 输出：6； len(元素) 统计元素的长度  */
	fmt.Println(e) /* 输出:[0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 2]  */
	fmt.Println(f) /* 输出：[1 0 0 0 5 0 0 0 0 10]  */

	fmt.Println(*p)  /*  输出：[1 0 0 0 5 0 0 0 0 10]，*指针变量: 取指针变量值 */
	fmt.Println(p1)

}

