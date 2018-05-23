package main

import "fmt"

func bsort(a []int) {

	for i := 0; i < len(a); i++ {
		for j := 1; j < len(a)-i; j++ {
			if a[j] < a[j-1] {
				a[j], a[j-1] = a[j-1], a[j]
			}
		}
	}
}

// 比较相邻的元素。如果第一个比第二个大，就交换它们两个；
// 对每一对相邻元素作同样的工作，从开始第一对到结尾的最后一对，这样在最后的元素应该会是最大的数；
// 针对所有的元素重复以上的步骤，除了最后一个；
// 重复步骤1~3，直到排序完成。
func bsort1(a []int) {
	lenA := len(a)
	for i := 0; i < lenA; i++ {
		for j := 0; j < lenA-i-1; j++ {
			if a[j] > a[j+1] {
				a[j], a[j+1] = a[j+1], a[j]
			}
		}
	}
}
func main() {
	b := [...]int{8, 7, 5, 4, 3, 10, 15}
	bsort1(b[:])
	fmt.Println(b)
}
