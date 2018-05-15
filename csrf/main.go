package main

import (
	"crypto/md5"
	"fmt"
	"io"
	"strconv"
	"time"
)

func genToken() string {
	h := md5.New()
	now := strconv.FormatInt(time.Now().Unix(), 10)
	key := "1234567890"
	fmt.Println("now time unix:", now)
	io.WriteString(h, now)
	io.WriteString(h, key)
	token := fmt.Sprintf("%x", h.Sum(nil))
	return token
}

func main() {
	genToken()
}
