package main

//任何类型都实现了一个空接口
import (
	"fmt"
)

type student struct {
	Name string
}

//类型断言
func just(items ...interface{}) {
	for index, v := range items {
		switch v.(type) {
		case bool:
			fmt.Printf("param %d is bool, value is %v\n", index, v)
		case int, int32, int64:
			fmt.Printf("param %d is int, value is %v\n", index, v)
		case float32, float64:
			fmt.Printf("param %d is float, value is %v\n", index, v)
		case string:
			fmt.Printf("param %d is string, value is %v\n", index, v)
		case student:
			fmt.Printf("param %d is student, value is %v\n", index, v)
		}
	}
}

type MyStruct struct {
}

func (m MyStruct) String() string {
	return ""
}

type Stringer interface {
	String() string
}

//断言，判断一个变量是否实现了指定接口
func check() {
	var v MyStruct
	if sv, ok := v.(Stringer); ok {
		fmt.Printf("v implements String():%s\n", sv.String())
	}
}
func main() {
	stu := student{
		Name: "Wangkang",
	}
	just(100, 1.00, "this is a string", stu)
}
