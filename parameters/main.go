package main

import "fmt"

func modify(ip *int) {
	fmt.Printf("函数里接收到的指针的内存地址是：%p\n", &ip)
	*ip = 1
}

func modifyMap(p map[string]int) {
	fmt.Printf("函数里接收到的map内存地址是： %p", &p)
	p["张三"] = 20
}
func main() {
	i := 10
	ip := &i
	fmt.Printf("原始指针的内存地址是：%p\n", &ip)
	modify(ip)
	fmt.Println("int值被修改了，新值为：", i)

	persons := make(map[string]int)
	persons["张三"] = 19
	mp := &persons

	fmt.Printf("原始map的内存地址是：%p\n", mp)

	modifyMap(persons)

	fmt.Println("map值被修改了，新值为：", persons)

}
