package main

import "fmt"

func main() {
	slice1 := []int{1, 2, 3}
	for i := 0; i < 16; i++ {
		slice1 = append(slice1, i)
		fmt.Printf("addr:%p,len:%v,cap:%v\n", slice1, len(slice1), cap(slice1))
	}
}
