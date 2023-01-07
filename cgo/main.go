package main

import (
	"fmt"
	"reflect"
)

func MyAdd(a, b int) int { return a + b }

func CallAdd(f func(a int, b int) int) {
	v := reflect.ValueOf(f)
	if v.Kind() != reflect.Func {
		return
	}
	argv := make([]reflect.Value, 2)
	argv[0] = reflect.ValueOf(1)
	argv[1] = reflect.ValueOf(1)

	result := v.Call(argv)

	fmt.Println(result[0].Int())
}

func main() {

	CallAdd(MyAdd)
}
