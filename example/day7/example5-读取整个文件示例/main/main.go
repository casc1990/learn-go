package main

import (
	"os"
	"io/ioutil"
	"fmt"
)

func main() {
	inputfile := "D:/go-project/src/learn-go/example/day7/example5-读取整个文件示例/main/products.txt"
	outfile := "products_copy.txt"

	buf,err := ioutil.ReadFile(inputfile)
	if err != nil {
		fmt.Fprintf(os.Stderr,"File Error: %s",err)
	}
	return

	fmt.Printf("%s\n",string(buf))
	err = ioutil.WriteFile(outfile,buf,0644)
	if err != nil {
		panic(err.Error())
	}
}
