package main

// 变量

import (
	"cmp"
	"fmt"
)

func main() {
	var rde = cmp.Compare(1, 2)
	var rda = cmp.Less(1, 2)
	fmt.Println(rde, rda)
}

func testVar() {
	var a int = 1
	var b int = 2
	println(a, b)
}

func testVar1() {
	var (
		a int = 11
		b int = 22
	)

	println(a, b)
}

func testVar2() {
	name := "张三"
	age := 19
	println(name, age)
}
