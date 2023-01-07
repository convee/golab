package main

import (
	"fmt"
	"reflect"
)

func main() {
	var a interface{}
	var b interface{}
	fmt.Println(reflect.TypeOf(a).Kind() == reflect.TypeOf(b).Kind())
	reflect.DeepEqual(a, b)
}

type test struct {
}
