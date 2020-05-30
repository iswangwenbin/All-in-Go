package main

import "fmt"

// 3. type 定义别名(就像孩子小时候的小名)
// 语法 type TypeAlias = Type ，TypeAlias 与 Type 是同一个类型
type IntAlias = int

func main() {
	var a int
	var b IntAlias
	a = 5
	b = a // 可以直接相互赋值
	fmt.Printf("a type : %T , a value : %d\n", a, a)
	fmt.Printf("b type : %T , b value : %d\n", b, b)
}
