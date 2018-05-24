package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	num := runtime.NumCPU()
	runtime.GOMAXPROCS(num - 1)
	for i := 0; i < 10240; i++ {
		go func() {
			for {
				t := time.NewTicker(time.Second)
				select {
				case <-t.C:
					fmt.Println("time out")
				}
				t.Stop()
			}
		}()
	}
	time.Sleep(100 * time.Second)
}
