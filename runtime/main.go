package main

import (
	"fmt"
	"runtime"
	"time"
)

func a() {
	for i := 1; i < 10; i++ {
		fmt.Println("A:", i)
	}
}

func b() {
	for i := 1; i < 10; i++ {
		fmt.Println("B:", i)
	}
}

func main() {

	runtime.GOMAXPROCS(1)
	go a()
	go b()
	fmt.Println(runtime.NumCPU())
	fmt.Println(runtime.NumGoroutine())
	time.Sleep(time.Second)
}
