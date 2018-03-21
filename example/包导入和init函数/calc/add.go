package calc

import (
	"fmt"
)


var Name string = "zhangsan"
var Age int = 200

func init() {
	fmt.Println("this is calc package...")
	fmt.Println("calc func Name:",Name)
	fmt.Println("calc func Age:",Age)
}

func Add(n1,n2 int)  {
	fmt.Printf("%d + %d = %d",n1,n2,n1+n2)
}