package main

import "fmt"

/**
费波那契数列由0和1开始，之后的费波那契系数就是由之前的两数相加而得出。首几个费波那契系数是：
0, 1, 1, 2, 3, 5, 8, 13, 21, 34, 55, 89, 144, 233……
*
*/
func fibonacci(n int) int {
	if n < 2 {
		return n
	}
	return fibonacci(n-2) + fibonacci(n-1)
}

func main() {
	for i := 0; i < 10; i++ {
		fmt.Printf("%d\t", fibonacci(i))
	}
}
