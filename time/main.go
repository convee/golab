package main

import (
	"fmt"
	"math/rand"
	"time"
)

func runtime() {
	now := time.Now()
	afterTime := time.Unix((time.Now().Unix() - 5), 0)
	beforeTime := time.Unix((time.Now().Unix() + 5), 0)
	after := now.After(afterTime)
	before := now.Before(beforeTime)
	time.Sleep(5 * time.Second)
	fmt.Printf("after => %v, before => %v\n", after, before)
}

func runtimes() {

}
func main() {
	times := rand.Intn(10000)
	fmt.Println(times)
	time.Sleep(time.Duration(times) * time.Microsecond)
	fmt.Println("Hello world")
	runtime()
	fmt.Println('a')
}
