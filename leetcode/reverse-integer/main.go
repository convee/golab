package main

import "fmt"

//给出一个 32 位的有符号整数，你需要将这个整数中每位上的数字进行反转。
//注意:
//假设我们的环境只能存储得下 32 位的有符号整数，则其数值范围为 [−2^31,  2^31 − 1]。请根据这个假设，如果反转后整数溢出那么就返回 0。
func reverse(x int) int {
	var nums, newnums int
	for x != 0 {
		a := x % 10
		newnums = nums*10 + a
		nums = newnums
		y := x / 10
		x = x / 10
		fmt.Println(y)

		maxInt32 := 1<<31 - 1
		minInt32 := -1 << 31
		if nums > maxInt32 || nums < minInt32 {
			return 0
		}
	}
	return nums
}
func main() {
	fmt.Println(reverse(123))
}
