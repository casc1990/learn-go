package main

import (
	"bufio"
	"os"
	"fmt"
	"io"
)

type CharCount struct {
	wordcount int
	numcount int
	spacecount int
	others int
}

func main() {
	//os.Open：只读方式打开文件。os.OpenFile()读写方式打开文件
	file,err := os.Open("c:/test.log")
	if err != nil {
		fmt.Println("open file err:",err)
		return
	}

	//defer函数先于return函数之前执行，关闭文件
	defer file.Close()
	//实例化一个带缓存的读文件对象
	reader := bufio.NewReader(file)
	var count CharCount
	//循环读每一行
	for {
		str, err := reader.ReadString('\n') //读到行尾
		if err == io.EOF {
			//fmt.Println("read file done")
			break
		}
		if err != nil {
			fmt.Println("read file err:", err)
			break
		}
		//将每行字符串转换成rune类型的切片，循环每个字符，判断类型
		runeArr := []rune(str)
		for _,v := range runeArr {
		switch {
			case v >= 'a' && v <= 'z':
				fallthrough
			case v >= 'A' && v <= 'Z':
				count.wordcount++
			case v >= '0' && v <= '9':
				count.numcount++
			case v == ' ' || v == '\t':
				count.spacecount++
			default:
				count.others++

		}

		}
	}
	fmt.Printf("char count: %d\n",count.wordcount)
	fmt.Printf("num count: %d\n",count.numcount)
	fmt.Printf("space count: %d\n",count.spacecount)
	fmt.Printf("other count: %d\n",count.others)
}
