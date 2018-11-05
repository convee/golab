package main

import (
	"bytes"
	"fmt"
	"io"
	"runtime"
	"sync"
	"sync/atomic"
)

func interfaces() {
	var b bytes.Buffer
	fmt.Fprintf(&b, "Hello World")

	var w io.Writer
	w = &b
	fmt.Println(w)
	fmt.Println(b.String())
}

func goroutine() {
	fmt.Println("cpu num:", runtime.NumCPU())
	runtime.GOMAXPROCS(runtime.NumCPU())
	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		defer wg.Done()
		for i := 1; i < 100; i++ {
			fmt.Println("A:", i)
		}
	}()
	go func() {
		defer wg.Done()
		for i := 0; i < 100; i++ {
			fmt.Println("B:", i)
		}
	}()
	wg.Wait()
}

var (
	count int32
	wg    sync.WaitGroup
)

func incCount() {
	defer wg.Done()
	for i := 0; i < 2; i++ {
		value := atomic.LoadInt32(&count)
		runtime.Gosched() //暂停当前goroutine
		value++
		atomic.StoreInt32(&count, value)
	}
}
func safeGoroutine() {
	wg.Add(2)
	go incCount()
	go incCount()
	wg.Wait()
	fmt.Println(count)

}

func main() {
	ch := make(chan int)
	go func() {
		var sum int = 0
		for i := 0; i < 10; i++ {
			sum += i
		}
		ch <- sum
	}()
	fmt.Println(<-ch) //45
}
