package main

import (
	"fmt"
	"os"
	"strings"
)

//os.Args Args保管了命令行参数，第一个是程序名。
func init() {
	fmt.Println(os.Args[0])
	fmt.Println(os.Args[1:])
}
func fun1() {
	var s, sep string
	for i := 1; i < len(os.Args); i++ {
		s += sep + os.Args[i]
		sep = " "
	}
	fmt.Println(s)
}

func fun2() {
	var s, sep string
	for _, arg := range os.Args[1:] {
		s += sep + arg
		sep = " "
	}
	fmt.Println(s)
}

func fun3() {
	sep := " "
	s := strings.Join(os.Args[1:len(os.Args)], sep)
	fmt.Println(s)
}
func main() {
	fun1()
	fun2()
	fun3()
}
