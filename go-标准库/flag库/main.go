package main

import (
	"fmt"
	"flag"
)

var (
	id int
	name string
	male bool
)

func main() {
	//定义命令行中需要使用的参数
	flag.IntVar(&id, "id",123,"id msg.")
	flag.StringVar(&name,"name","liming","user name msg.")
	flag.BoolVar(&male,"male",false,"male msg.")
	//也可以用以下带返回值的方法代替，不过他们返回的是指针，比较麻烦点
	married := flag.Bool("married", false, "Are you married?")

	//是否解析
	if flag.Parsed() != true {
		flag.Parse()
	}
	// 获取非flag参数（不是命令行指定的参数）
	fmt.Println(flag.NArg()) //获取非flag参数的个数
	fmt.Println("------ Args start ------")
	fmt.Println(flag.Args()) //获取非flag参数的参数列表
	for i, v := range flag.Args() {  //循环非flag参数列表
		fmt.Printf("arg[%d] = (%s).\n", i, v)
	}
	fmt.Println("------ Args end ------")
	//
	// visit只包含已经设置了的flag（命令行指定的参数）
	fmt.Println("------ visit flag start ------")
	flag.Visit(func(f *flag.Flag) {  //函数必须传递指针
		fmt.Println(f.Name, f.Value, f.Usage, f.DefValue)  //打印指定参数的名字，值，用法，默认值
	})
	fmt.Println("------ visit flag end ------")
	//
	// visitAll只包含所有的flag(包括未设置的)
	fmt.Println("------ visitAll flag start ------")
	flag.VisitAll(func(f *flag.Flag) {
		fmt.Println(f.Name, f.Value, f.Usage, f.DefValue)

	})
	fmt.Println("------ visitAll flag end ------")
	//
	// flag参数
	fmt.Printf("id = %d\n", id)
	fmt.Printf("name = %s\n", name)
	fmt.Printf("male = %t\n", male)
	fmt.Println("married = %t\n",married)

	// flag参数默认值
	fmt.Println("------ PrintDefaults start ------")
	flag.PrintDefaults()
	fmt.Println("------ PrintDefaults end ------")

	//非flag参数个数
	fmt.Printf("NArg = %d\n", flag.NArg())
	// 已设置的flag参数个数
	fmt.Printf("NFlag = %d\n", flag.NFlag())

	}


