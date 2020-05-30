package main

import "fmt"

// 1. type 定义结构体
type Person struct {
	name string
	age  int
}

func main() {
	p := Person{
		name: "George",
		age:  5,
	}
	fmt.Printf("%+v", p)
}
