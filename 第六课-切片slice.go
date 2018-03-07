package main
/*
 切片Slice

其本身并不是数组，它指向底层的数组
作为变长数组的替代方案，可以关联底层数组的局部或全部
为引用类型
可以直接创建或从底层数组获取生成
使用len()获取元素个数，cap()获取容量
一般使用make()创建
如果多个slice指向相同底层数组，其中一个的值改变会影响全部

make([]T, len, cap)
其中cap可以省略，则和len的值相同
len表示存数的元素个数，cap表示容量
 */
import (
	"fmt"

)

func main() {
	var arrays [3]int
	var slices []int  /* 直接创建slice，slice和array的区别就是array要指定长度 */
	var a []string  /* 直接创建的slice为空slice，长度为0 */
	a1 := [...]int{0,1,2,3,4,5,6,7,8,9}
	s1 := a1[5:10]  /*  从a1数组中获取；取5-10这个范围的元素（下标4-9） */
	s2 := a1[6:] /* 取6到最后 */
	s3 := a1[:8]
	s4 := a1[:]  /* 取所有元素 */

	/*  make([]TYPE, len, cap)，len和cap默认都可以省略，cap省略表示和len的值相同，如果超过cap容量，会重新分配内存空间 */
	s5 := make([]int,3,10)  /* make创建语法：make([]TYPE, len, cap)，len表示存放元素的个数，cap表示预分配内存空间长度  */
	fmt.Println(s5)

	/* Reslice(对slice在做slice)
	 Reslice时索引以被slice的切片为准
	索引不可以超过被slice的切片的容量cap()值(slice的cap等于array的cap值减去slice已经使用的cap值)
	索引越界不会导致底层数组的重新分配而是引发错误
	 */
	a2 := [...]string{"a","b","c","d","e","f","g","h","i","j"}
	s6 := a2[2:5]
	s7 := s6[1:3] /* 对slice在做slice，取reslice里的第二到第五个元素 */
	fmt.Println(s6,len(s6),cap(s6)) /* 输出:[c d e] 3 8,因为slice指向的是底层的array，所以cap=array_cap-slice_useCAP */
	fmt.Println(s7,len(s7),cap(s7)) /* 输出：[d e] 2 7 */

	/* Append
	可以在slice尾部追加元素
	可以将一个slice追加在另一个slice尾部
	如果最终长度未超过追加到slice的容量则返回原始slice
	如果超过追加到的slice的容量则将重新分配数组并拷贝原始数据
	*/
	s8 := make([]int,3,6)
	fmt.Printf("%p\n",s8) /* 打印内存地址：0xc042066030 */
	s8 = append(s8,1,2,3)
	fmt.Printf("%v %p\n",s8,s8) /*  输出：[0 0 0 1 2 3] 0xc04200a1e0；内存地址没变 */
	s8 = append(s8,4,5,6)
	fmt.Printf("%v %p\n",s8,s8) /*  输出：[0 0 0 1 2 3 4 5 6] 0xc04203c060；内存变了，因为cap容量到上限了，重新分配了*/

	a3 := [5]int{1,2,3,4,5}
	s9 := a3[2:5]
	s10 := a3[1:3]
	fmt.Print(s9,s10) /* 输出：[3 4 5] [2 3]，交集为3（下标为2） */
	s9[0] = 9 /* 如果多个slice指向相同底层数组，其中一个的值改变会影响全部 */
	fmt.Print(s9,s10) /* 输出：[9 4 5] [2 9] */

	/* copy方法  语法：copy(目标，源)：把源slice拷贝到目标slice */
	s11 := []int{1,2,3,4}
	s12 := []int{5,6,7,8,9,10}
	copy(s12,s11)  /*  把s11拷贝到s12；默认从头开始替换，默认替换len(目标)次 */
	fmt.Print(s12) /* 输出：[1 2 3 4 9 10] */

	fmt.Println("array:",arrays,"slice:",slices)  /* 输出：array: [0 0 0] slice: []  */
	fmt.Println(a)
	fmt.Println(a1)
	fmt.Println(s1)
	fmt.Println(s2,s3,s4)
}
