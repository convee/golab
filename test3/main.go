package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

func main() {
	per := Person{
		Name: "test",
		Age:  8,
	}
	fmt.Printf("原始struct的内存地址是：%p\n", &per)
	modifyStruct(per)
	fmt.Printf("改动后的值是：%v\n", per)
	defer
}
func modifyStruct(per Person) {
	fmt.Printf("函数里接收到struct的内存地址是：%p\n", &per)
	per.Age = 10
}
