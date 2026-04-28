package main

// 数据类型
const (
	Num  = iota * 2 // 0
	Num1            // 2
	Num2            // 4
	Num3            // 6
	Num4            // 8
)

const (
	Numb  = iota // 0
	Numb1        // 1
	Numb2        // 2
	Numb3        // 3
	Numb4        // 4
)

func main() {
	println(Num, Num1, Num2, Num3, Num4)
	println(Numb, Numb1, Numb2, Numb3, Numb4)
}
