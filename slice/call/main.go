package main

import "fmt"

func main() {
	slice := make([]int, 0)
	slice = append(slice, 1)
	fmt.Println(slice, len(slice), cap(slice))
}
