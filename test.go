package main

import "fmt"

func main() {
	fmt.Println(len("Go语言"))         //8
	fmt.Println(len([]rune("Go语言"))) // 4
	fmt.Println(len([]byte("Go语言"))) // 8
}
