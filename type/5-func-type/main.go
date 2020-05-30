package main

import "fmt"

// 5 type 定义新的函数类型
type NewFunc func(name string) string

func main() {
	var a NewFunc = func(name string) string {
		return "My name is " + name
	}
	fmt.Printf("%s\n", a("George"))
}
