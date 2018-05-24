package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int, 10)
	ch2 := make(chan int, 10)
	for i := 0; i < 10; i++ {
		ch <- i
		ch2 <- i
	}

	for {
		select {
		case v := <-ch:
			fmt.Println(v)
		case v := <-ch2:
			fmt.Println(v)
		case <-time.After(time.Microsecond): //todo 存在内存泄漏
			fmt.Println("get time out")
		}
	}
}
