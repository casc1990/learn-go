package go基础语法
// map也是python中字典的概念，格式: map[keytype]valuetype
/* make函数和new函数
	make函数：用于内建类型（map、slice 和channel）的内存分配。make返回初始化后的（非零）值。
	new函数：new返回指针
  */
import (
	"fmt"
)

func main() {
	// 声明一个key是字符串，值为int类型的字典，这种方式的声明使用前需要使用make初始化
	var number map[string]int

	// 另一种方式声明：make(map[keytype]valuetype)
	numbers := make(map[string]int)
	numbers["one"] = 1  //赋值
	numbers["two"] = 2
	numbers["three"] = 3
	fmt.Println("number：",number) /* 输出：number： map[];空字典 */
	//读取数据
	fmt.Printf("numbers字典里key是three的值是：%d\n",numbers["three"]) //numbers字典里key是three的值是：3
	//修改数据或者新增（key存在修改，不存在新增）
	numbers["three"] = 33
	fmt.Printf("numbers字典里key是three的值是：%d\n",numbers["three"])

	//map可以通过key:value初始化
	learnning := map[string]float32 {
		"C" : 5,
		"Go" : 4.5,
		"Python" : 4.5,
		"C++" : 2,
	}
	fmt.Println(learnning)
	//map有两个返回值，第二个返回值，如果不存在key，那么ok为false，如果存在ok为true
	results,ok := learnning["C#"]
	if ok {
		fmt.Println("C# is in the map and its value is ",results)
	}else {
		fmt.Println("C# not in the map...")
	}

	//删除key
	delete(learnning,"ee")  /* 删除不存在的key也不会保错 */

	//map也是一种引用类型
	m := make(map[string]string)
	m["Hello"] = "xiaoming"
	m1 := m
	m1["Hello"] = "xiaomei"
	fmt.Print(m["Hello"])  //现在m["Hello"]的值是xiaomei


}
