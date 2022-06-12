package main

import "fmt"

//slice，map, chan 是引用类型
//1分配内存
//2初始化
func main() {
	slice := []int{1, 2, 3, 4, 5}

	newSlice := slice[1:3]
	newSlice[0] = 10
	fmt.Println(slice)
	fmt.Println(newSlice)
	fmt.Println(len(slice))
	fmt.Println(cap(slice))
	slice = append(slice, 20)
	fmt.Println(slice)
	fmt.Println(cap(slice)) //20

	for _, v := range slice {
		fmt.Println(v)
	}
}

/**
打印结果：

[1 10 3 4 5]
[10 3]
*/


func RemoveCopy(slice []int, i int) []int {
    copy(slice[i:], slice[i+1:])
    return slice[:len(slice)-1]
}