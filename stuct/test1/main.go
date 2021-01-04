package main

import (
	"fmt"
	"runtime"
	"sync"
)

type student struct {
	Name string
	Age  int
}

func parseStudent() {
	m := make(map[string]*student)
	stus := []student{
		student{Name: "convee1", Age: 13},
		student{Name: "convee2", Age: 14},
		student{Name: "convee3", Age: 15},
	}
	for k, stu := range stus {
		m[stu.Name] = &stus[k]
	}
	fmt.Println(m)
}

func syncc() {
	runtime.GOMAXPROCS(1)
	wg := sync.WaitGroup{}
	wg.Add(20)
	for i := 0; i < 10; i++ {
		go func() {
			fmt.Println("i: ", i)
			wg.Done()
		}()
	}
	for i := 0; i < 10; i++ {
		go func(i int) {
			fmt.Println("i: ", i)
			wg.Done()
		}(i)
	}
	wg.Wait()
}

func chanc() {
	fmt.Println(runtime.NumCPU())
	runtime.GOMAXPROCS(runtime.NumCPU())
	int_chan := make(chan int, 1)
	string_chan := make(chan string, 1)
	int_chan <- 1
	string_chan <- "hello"
	select {}
}

func twoSum(nums []int, target int) []int {
	maps := make(map[int]int)
	for index, value := range nums {
		member := target - value
		if _, ok := maps[member]; ok {
			return []int{member, value}
		} else {
			maps[value] = index
		}
	}
	return []int{}

}

func main() {
	// parseStudent()
	// syncc()
	nums := []int{2, 7, 11, 15}
	a := twoSum(nums, 9)
	fmt.Println(a)

}
