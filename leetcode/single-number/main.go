package main

import "fmt"

//给定一个非空整数数组，除了某个元素只出现一次以外，其余每个元素均出现两次。找出那个只出现了一次的元素。
func singleNumber(nums []int) int {
	lenth := len(nums)
	result := 0
	for i := 0; i < lenth; i++ {
		result ^= nums[i]
	}
	return result
}
func main() {
	nums := []int{2, 2, 3}
	fmt.Println(singleNumber(nums))
}
