package main

import (
	"bytes"
	"fmt"
	"strings"
)

func StringPlus() string {
	var s string
	s += "昵称" + "convee"
	return s
}

func StringFmt() string {
	return fmt.Sprint("昵称", "convee")
}

func StringJoin() string {
	s := []string{"昵称", "convee"}
	return strings.Join(s, "")
}

func StringBuffer() string {
	var b bytes.Buffer
	b.WriteString("昵称")
	b.WriteString("convee")
	return b.String()
}

func StringBuilder() string {
	var b strings.Builder
	b.WriteString("昵称")
	b.WriteString("convee")
	return b.String()
}

func main() {
	fmt.Println("字符串拼接性能比较")
}
