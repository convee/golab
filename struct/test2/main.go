package main

import (
	"fmt"
	"reflect"
)

type Student struct {
	Name  string
	Age   int
	Score float32
}

func test(b interface{}) {
	t := reflect.TypeOf(b)
	fmt.Println(t)

	v := reflect.ValueOf(b)
	k := v.Kind()
	fmt.Println(k)

	iv := v.Interface()
	stu, ok := iv.(Student)
	if ok {
		fmt.Printf("%v %T\n", stu, stu)
	}
}

func testInt(b interface{}) {
	val := reflect.ValueOf(b)
	val.Elem().SetInt(100)
	c := val.Elem().Int()
	fmt.Printf("get value interface{} %d \n", c)
	fmt.Printf("string val:%d\n", val.Elem().Int())
}
func main() {
	a := Student{
		Name:  "convee",
		Age:   18,
		Score: 100,
	}
	test(a)
	testInt(a)
}
