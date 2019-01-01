package main

import "fmt"

func romanToInt(s string) int {
	a := map[rune]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}
	runes := []rune(s)
	ret := a[runes[0]]
	if len(runes) == 1 {
		return ret
	}
	for i := 1; i < len(runes); i++ {
		if a[runes[i-1]] < a[runes[i]] {
			ret += a[runes[i]] - 2*a[runes[i-1]]
		} else {
			ret += a[runes[i]]
		}
	}
	return ret
}
func main() {
	fmt.Println(romanToInt("IV"))
}
