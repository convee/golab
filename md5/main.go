package main

import (
	"crypto/md5"
	"fmt"
)

func main() {
	pwd := "123456"
	md5 := fmt.Sprintf("%x", md5.Sum([]byte(pwd)))
	fmt.Println(md5)
	fmt.Println(len(md5))
}
