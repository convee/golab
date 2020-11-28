package main

import (
	"fmt"
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().Unix())
}
func main() {
	n := rand.Intn(999)
	fmt.Println(n)
}
