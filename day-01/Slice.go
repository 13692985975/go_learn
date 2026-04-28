// 切片
package main

func main() {
	nums := [5]int{1, 2, 3, 4, 5}

	println(nums[:])   // 子切片范围[0,5) -> [1 2 3 4 5]
	println(nums[1:])  // 子切片范围[1,5) -> [2 3 4 5]
	println(nums[:5])  // 子切片范围[0,5) -> [1 2 3 4 5]
	println(nums[2:3]) // 子切片范围[2,3) -> [3]
	println(nums[1:3]) // 子切片范围[1,3) -> [2 3]
}

func main1() {
	var nums [5]int

	println(len(nums))
	println(cap(nums))
}
