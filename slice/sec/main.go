package main

import (
	"fmt"
	"sync"
)

// 切片非并发安全
// 多次执行，每次得到的结果都不一样
// 可以考虑使用 channel 本身的特性（阻塞）来实现安全的并发读写
func main() {
	a := make([]int, 0)
	var wg sync.WaitGroup
	for i := 0; i < 10000; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			a = append(a, i)

		}(i)
	}
	wg.Wait()
	fmt.Println(len(a))
	// not equal 10000
}
