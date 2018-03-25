package main

import (
	"fmt"
	"strings"
)

func main() {
	for i :=0;i<5;i++ {
		fmt.Println(strings.Repeat("A",i+1))
	}
}
