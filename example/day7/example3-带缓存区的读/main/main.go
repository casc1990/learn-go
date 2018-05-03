package main

import (
	"bufio"
	"os"
	"fmt"
	"io"
)

//从标准输入读取
func OsStdin() {
	//创建读缓存区实例(bufio.NewReader)，接收来自标准输入的数据
	reader := bufio.NewReader(os.Stdin)
	// 从头读到末尾（'\n'末尾，字节类型要用单引号）
	str,err := reader.ReadString('\n')
	if err == io.EOF {
		return
	}
	if err != nil {
		fmt.Println("read file err:",err)
		return
	}
	fmt.Printf("read str success! result: %s",str)
}

//从文件读取
func ReadFile() {
	file,err := os.Open("c:/test.log")
	if err == io.EOF {
		fmt.Println("open file err:",err)
		return
	}
	defer file.Close()
	reader := bufio.NewReader(file)
	// 从头读到行末尾（'\n'末尾，字节类型要用单引号）
	str,err := reader.ReadString('\n')
	if err == io.EOF {
		fmt.Println("read string done!")
	}
	fmt.Printf("read str success! result: %s",str)
}

func main() {
	//OsStdin()
	ReadFile()
}
