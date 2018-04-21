package main

//map的声明：var map1 map[keytype]valuetype


import "fmt"

func testMap() {
	//定义及赋值
	var a  map[string]string = map[string]string {
		"key" : "value",  //记得加逗号，不然就报错了
	}

	a["abc"] = "def"
	fmt.Println(a)  //map[key:value abc:def]

	// 使用make声明
	b := make(map[string]string)
	b["test"] = "test1"
	fmt.Println(b)
/*
	//错误使用方法(这样只是定义了map类型，未分配内存地址)
	var t map[string]string
	t["hello"] = "world"
	fmt.Println(t)
*/

}

//map的例子：判断某个key是否存在，存在就修改里面的内容，不存在就新建
func modify(a map[string]map[string]string) {  //map嵌套
	val,ok := a["zhangsan"]  //查找key是否存在
	if ok {
		val["passwd"] = "123456"
		val["phone"] = "1100000000000"
	}else {
		//初始化字典并赋值
		a["zhangsan"] = make(map[string]string)
		a["zhangsan"]["passwd"] = "123456"
		a["zhangsan"]["phone"] = "1100000000000"
	}
}

//map遍历
func traverse(a map[string]map[string]string) {
	for k,v := range a {
		fmt.Println(k)
		for k1,v1 := range v {
			fmt.Println("\t",k1,v1)
		}
	}
}

func testMap2() {
	a := make(map[string]map[string]string)
	modify(a)
	fmt.Println(a)

}

func testMap3() {
	//多维map
	a := make(map[string]map[string]string)
	a["key1"] = make(map[string]string)
	a["key1"]["key2"] = "abc"
	a["key1"]["key3"] = "abc"
	a["key1"]["key4"] = "abc"
	a["key1"]["key5"] = "abc"
	a["key2"] = make(map[string]string)
	a["key2"]["key2"] = "abc"
	a["key2"]["key3"] = "abc"
	a["key2"]["key4"] = "abc"

	traverse(a)
	//删除某个key,go没清空map的方法，要想清空可以重新make或者遍历key，逐个清除
	delete(a,"key1")
	fmt.Println()
	fmt.Println(a)
	//map长度是key的个数
	fmt.Println(len(a))
}

func testMap4() {
	var a []map[int]int
	//引用类型在使用前一定要初始化,如 map，slice，channel
	a = make([]map[int]int,10)

	// map、slice、数组 都可以用某个元素是否为nil来判定是否有值
	if a[0] == nil {
		a[0] = make(map[int]int) //map一定要初始化
	}
	a[0][10] = 10
	fmt.Println(a) //[map[10:10] map[] map[] map[] map[] map[] map[] map[] map[] map[]]
}

func main() {
	testMap()
	testMap2()
	testMap3()
	testMap4()
}