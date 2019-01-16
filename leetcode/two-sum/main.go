package main

//给定一个整数数组 nums 和一个目标值 target，请你在该数组中找出和为目标值的那 两个 整数，并返回他们的数组下标。

//你可以假设每种输入只会对应一个答案。但是，你不能重复利用这个数组中同样的元素。

func twoSum(nums []int, target int) []int {
	l := len(nums)
	for i := 0; i < l; i++ {
		for j := 0; j < l; j++ {
			if nums[j] == target-nums[i] {
				return []int{i, j}
			}
		}
	}
	return []int{}
}

func twoSum2(nums []int, target int) []int {
	m := make(map[int]int)
	for index, num := range nums {
		anotherNum := target - num
		if _, ok := m[anotherNum]; ok {
			return []int{m[anotherNum], index}
		}
		m[num] = index
	}
	return []int{}
}
func main() {

}
