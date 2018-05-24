package main

import (
	"fmt"
	"runtime"
	"time"
)

func test() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("panic:", r)
		}
	}()
	var m map[string]int
	m["stu"] = 100
}
func main() {
	num := runtime.NumCPU()
	runtime.GOMAXPROCS(num)
	for i := 0; i < 100; i++ {
		go test()
	}
	time.Sleep(10000 * time.Second)
}
