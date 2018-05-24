package main

import "fmt"

//插入排序
// 从第一个元素开始，该元素可以认为已经被排序；
// 取出下一个元素，在已经排序的元素序列中从后向前扫描；
// 如果该元素（已排序）大于新元素，将该元素移到下一位置；
// 重复步骤3，直到找到已排序的元素小于或者等于新元素的位置；
// 将新元素插入到该位置后；
// 重复步骤2~5。
func isort(a []int) {

	for i := 1; i < len(a); i++ {
		for j := i; j > 0; j-- {
			if a[j] > a[j-1] {
				break
			}
			a[j], a[j-1] = a[j-1], a[j]
		}
	}
}

func isort1(a []int) {
	lenA := len(a)
	for i := 1; i < lenA; i++ {
		for j := i; j > 0; j-- {
			if a[j] < a[j-1] {
				a[j], a[j-1] = a[j-1], a[j]
			}
		}
	}
}

func main() {
	b := [...]int{8, 7, 5, 4, 3, 10, 15}
	isort1(b[:])
	fmt.Println(b)
}
