package main

import "fmt"

//选择排序
func ssort(a []int) {

	for i := 0; i < len(a); i++ {
		var min int = i
		for j := i + 1; j < len(a); j++ {
			if a[min] > a[j] {
				min = j
			}
		}
		a[i], a[min] = a[min], a[i]
	}
}

func ssort1(a []int) {
	lenA := len(a)
	for i := 0; i < lenA; i++ {
		var min int = i
		for j := i + 1; j < lenA; j++ {
			if a[j] < a[min] {
				min = j
			}
		}
		a[i], a[min] = a[min], a[i]

	}
}
func main() {
	b := [...]int{8, 7, 5, 4, 3, 10, 15}
	ssort1(b[:])
	fmt.Println(b)
}
