package main

import "fmt"

// 4. type 定义新类型
type NewInt int

type IntAlias = int // type 定义别名

func main() {
	var a NewInt
	fmt.Printf("a type: %T\n", a)

	var b IntAlias
	fmt.Printf("b type: %T\n", b)

	var c int
	fmt.Printf("c type: %T\n", c)
}
