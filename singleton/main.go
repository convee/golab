package main

import (
	"fmt"
	"sync"
)

type singleton struct {
}

func (s *singleton) test() {
	fmt.Println("test")
}

var instance *singleton
var once sync.Once
var mu sync.Mutex

func GetInstance() *singleton {
	once.Do(func() {
		instance = &singleton{}
	})
	return instance
}

//线程安全单例
func GetInstance2() *singleton {
	if instance == nil {
		mu.Lock()
		defer mu.Unlock()
		instance = &singleton{}
	}
	return instance
}
func main() {
	GetInstance().test()
	GetInstance2().test()
}
