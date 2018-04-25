package main

import "fmt"


type Reader interface {
	Read()
}
type Writer interface {
	Writer()
}
//接口嵌套
type ReadWrite interface {
	Read()
	Writer()
}

//File实现了ReadWrite、Reader、Writer接口
type File struct {

}
func (f *File) Read() {
	fmt.Println("read data..")
}
func (f *File) Writer() {
	fmt.Println("write data..")
}

//实现了ReadWrite接口才可以调用Test函数（实现了ReadWrite接口的对象包含有Read和Writer方法）
func Test(rw ReadWrite) {
	rw.Read()
	rw.Writer()
}

func main() {
	var f File
	//调用Test方法,因为是*File指针实现了ReadWrite接口，所有要传指针
	Test(&f)
}