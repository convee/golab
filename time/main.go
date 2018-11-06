package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	times := rand.Intn(10000)
	fmt.Println(times)
	time.Sleep(time.Duration(times) * time.Microsecond)
	fmt.Println("Hello world")
}
