package main

import "fmt"

//第归方式完成二分查找
func binSearch(nums []int, start int, end int, target int) int {
	mid := int((start + end) / 2)
	if start <= end {
		if nums[mid] == target {
			return mid
		} else if target < nums[mid] {
			return binSearch(nums, 0, mid-1, target)
		} else {
			return binSearch(nums, mid+1, end, target)
		}
	}
	return -1
}

//非递归方式完成二分查找
func binSearch2(nums []int, start int, end int, target int) int {
	for start <= end {
		mid := int((start + end) / 2)
		if target == nums[mid] {
			return mid
		} else if target > nums[mid] {
			start = mid + 1
		} else {
			end = mid - 1
		}
	}
	return -1
}
func main() {
	nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	target := 9
	fmt.Println(binSearch2(nums, 0, len(nums)-1, target))
}
