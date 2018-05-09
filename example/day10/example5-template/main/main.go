package main

import (
	 "fmt"
	"html/template"
	"os"
)

type Person struct {
	Name string
	Age int
	score float32
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	t,err := template.ParseFiles("D:/python-projects/learn-go/example/day10/example5-template/main/index.html")
	if err != nil {
		fmt.Println("parse html file err:",err)
		return
	}

	p := Person{
		Name: "stu01",
		Age: 18,
		score: 90,
	}
	if err := t.Execute(os.Stdout,p); err != nil {
		fmt.Println("there was an err:",err.Error())
	}
}