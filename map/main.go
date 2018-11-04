package main

import "fmt"

//map的key是无序的
func main() {
	dist := make(map[int]string)
	dist[1] = "666"
	dist[2] = "777"
	fmt.Println(dist)
}
