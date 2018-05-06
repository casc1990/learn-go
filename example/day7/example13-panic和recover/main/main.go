package main

import "fmt"

func badCall() {
	panic("calling failed!")
	//fmt.Println("hello")
}

func test() {
	defer func() {
		if e := recover();e != nil {
			fmt.Printf("panicking %s\r\n",e)
		}
	}()
	badCall()
	fmt.Println("After bad call")
}

func main() {
	fmt.Printf("Calling test\r\n")
	test()
}