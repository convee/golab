package main

import "fmt"

// race 帮助监测代码中的数据竞争
// go run -race goroutine/vie/main.go
func main() {
	i := 0

	go func() {
		i++
	}()
	fmt.Println(i)
}
