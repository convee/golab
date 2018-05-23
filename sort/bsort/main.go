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
