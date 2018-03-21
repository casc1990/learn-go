package test

import "fmt"

var Name string = "lisi"
var Age int = 20

func init() {
	fmt.Println("this is test package..")
	fmt.Println("test func Name: ",Name)
	fmt.Println("test func Age:",Age)
	Age = 30
	fmt.Println("test func Age:",Age)
}
