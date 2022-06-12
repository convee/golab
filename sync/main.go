package main

import (
	"fmt"
	"math/rand"
	"sync"
	"sync/atomic" //atomic包提供了底层的原子级内存操作
	"time"
)

var rwLock sync.RWMutex //读写锁适用于读多写少，用错，性能相差100倍
var lock sync.Mutex     //互斥锁适用于读少写多
func testRWLock() {
	a := make(map[int]int, 5)
	var count int32
	a[1] = 10
	a[2] = 10
	a[3] = 10
	a[4] = 10
	a[5] = 10
	for i := 0; i < 2; i++ {
		go func(b map[int]int) {
			rwLock.Lock()
			b[4] = rand.Intn(100)
			time.Sleep(10 * time.Millisecond)
			rwLock.Unlock()
		}(a)
	}

	for i := 0; i < 100; i++ {
		go func(b map[int]int) {
			for {
				rwLock.RLock()
				time.Sleep(time.Millisecond)
				rwLock.RUnlock()
				atomic.AddInt32(&count, 1)
			}
		}(a)
	}
	time.Sleep(time.Second * 3)
	fmt.Println(atomic.LoadInt32(&count))
}

func main() {
	var wg sync.WaitGroup
	startT := time.Now()
	for i := 0; i < 10000; i++ {
		go func(i int) {
			defer wg.Done()
			fmt.Println(i)
			time.Sleep(time.Second)
		}(i)
	}
	wg.Wait()
	tc := time.Since(startT) //计算耗时
	fmt.Printf("time cost = %v\n", tc)
}
