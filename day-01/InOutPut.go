// 输入输出
package main

import (
	"fmt"
	"os"
)

func main() {
	a := 1
	if a == 1 {
		goto A
	} else {
		fmt.Println("b")
	}
A:
	fmt.Println("a")
}

func main8() {
	num := 2
	switch {
	case num >= 0 && num <= 1:
		num++
	case num > 1:
		num--
		fallthrough // 执行完该分支后，会继续执行下一个分支
	case num < 0:
		num += num
	}
	fmt.Println(num)
}

func main7() {
	if x := 1 + 1; x >= 2 {
		fmt.Println(x)
	}
}

func main6() {
	a := 1
	var b = a << 2
	println(b)
}

func main5() {
	str := "abcdefg"
	fmt.Printf("%x\n", str)
	fmt.Printf("% x\n", str)
}

func main4() {
	fmt.Printf("hello world, %s!", "jack")

	fmt.Printf("%%%s\n", "hello world")

	fmt.Printf("%s\n", "hello world")
	fmt.Printf("%q\n", "hello world")
}

func main3() {
	fmt.Println("chagnxinyue")
}

func main2() {
	println("nihao")
}

func main1() {
	os.Stdout.WriteString("hello world!")
}
