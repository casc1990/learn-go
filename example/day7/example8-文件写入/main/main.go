package main

import (
	"os"
	"fmt"
	"bufio"
	"io"
)

//读文件
func Writefile() {
	file,err := os.OpenFile("output.log",os.O_WRONLY|os.O_CREATE,0644)
	if err != nil {
		fmt.Println("an error occurred with file create.",err)
		return
	}

	defer file.Close()
	writer := bufio.NewWriter(file)
	for i:=0;i<10;i++ {
		writer.WriteString(fmt.Sprintf("hello!stu%d\n",i))
	}
	writer.Flush()

}


//copy文件（io.Copy(dst,src)）
func CopyFile(dstName,srcName string) (written int64,err error) {
	src,err := os.Open(srcName)
	if err != nil {
		return
	}
	defer src.Close()
	dst,err := os.OpenFile(dstName,os.O_WRONLY|os.O_CREATE,0644)
	if err != nil {
		return
	}
	defer dst.Close()
	return io.Copy(dst,src)
}

func main() {
	Writefile()
}
