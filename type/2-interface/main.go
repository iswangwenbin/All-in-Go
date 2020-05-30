package main

import "fmt"

// 2. type 定义接口
type Personer interface {
	Name() string
}

type person struct {
	name string
	age  int
}

func (p person) Name() string {
	return p.name
}

func main() {
	var p Personer
	p = person{name: "George", age: 5}
	fmt.Printf("Name:%s\n", p.Name())
}
