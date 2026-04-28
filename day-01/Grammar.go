package main

// 语法
func main() {
	println("hello print")
	test01("lisi")
	test03()
}

func test01(name string) {
	println(name)

}

func test02() {
	var temp = 2*3 + 6*9
	println(temp)
}

func test03() {
	for i := 0; i < 10; i++ {
		println(i)
	}
}
