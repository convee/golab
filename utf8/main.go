package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	l := utf8.RuneCountInString("昂科威Plus")
	li := len("昂科威Plus")
	fmt.Println(l)
	fmt.Println(li)

	res := l - (l-(li-l)/2)/2
	fmt.Println(res)
}
