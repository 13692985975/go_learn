// 循环控制

package main

import "fmt"

func main() {
	for i := range 10 {
		fmt.Println(i)
	}
}

func main2() {
	sequence := "hello world"
	for index, value := range sequence {
		fmt.Println(index, value)
	}
}

func main1() {
	for i, j := 1, 2; i < 100 && j < 1000; i, j = i+1, j+1 {
		fmt.Println(i, j)
	}
}
