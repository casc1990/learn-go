package main

import "fmt"

//快速排序的原理是，首先找到一个数pivot把数组‘平均'分成两组，使其中一组的所有数字均大于另一组中的数字，
// 此时pivot在数组中的位置就是它正确的位置。然后，对这两组数组再次进行这种操作。


//快速排序
func qsort(a []int,left,right int) {
	if left >= right {
		return
	}

	//确定val所在的位置
	val := a[left]
	k := left

	for i:=left+1;i<=right;i++ {
		if a[i] < val {
			a[k] = a[i]
			a[i] = a[k+1]
			k++
		}
	}

	a[k] = val
	qsort(a,left,k-1)
	qsort(a,k+1,right)
}

func main() {
	b := []int{8,7,5,2,6,10,1,44,21}
	qsort(b,0,len(b)-1)
	fmt.Println(b)
}