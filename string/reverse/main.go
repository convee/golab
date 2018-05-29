package main

import (
	"fmt"
)

func reverseString(s string) string {

	r := []rune(s)
	rLen := len(r)
	for i, j := 0, (rLen - 1); i < rLen/2; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
	return string(r)
}

func main() {
	s := "hello golang语言"
	fmt.Println(reverseString(s))
}
