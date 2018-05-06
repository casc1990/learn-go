package main

import (
	"errors"
	"fmt"
	"os"
	"bufio"
	"io"
	"time"
)

var NotFoundErr error = errors.New("Not found error")

type errs struct {
	Op string
	Path string
	CreateTime string
	message string

}

//自定义错误，实现Error方法（实现了type error interface接口）,错误就可以按我们定义的打印了
func (e *errs)  Error() string {
	return fmt.Sprintf("path=%s,op=%s,createtime=%s,message=%s",e.Path,
		e.Op,e.CreateTime,e.message)
}

func FromFileReader(f string) (err error) {
	file,err := os.Open(f)
	if err != nil {
		return &errs{
			Op: "open",
			Path: f,
			message: err.Error(),
			CreateTime: fmt.Sprintf("%v",time.Now()),
		}
	}
	defer file.Close()

	reader := bufio.NewReader(file)
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
		fmt.Println(str)

	}
	return nil
}

func main() {
	fmt.Printf("error:%s\n",NotFoundErr) //error:Not found error

	//自定义错误
	err := FromFileReader("c:/texaasvasfg.log")
	//判断是不是我们自定义的错误
	switch v := err.(type) {
	case *errs:
		fmt.Println(v)
	default:
	}

}
