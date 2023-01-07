package main

import (
	"fmt"
	"math/rand"
	"time"
)

//1、如果不使用rand.Seed(seed int64)，每次运行，得到的随机数会一样，程序不停止，一直获取的随机数是不一样的；

//2、每次运行时rand.Seed(seed int64)，seed的值要不一样，这样生成的随机数才会和上次运行时生成的随机数不一样；

//3、rand.Intn(n int)得到的随机数int i，0 <= i < n。

func init() {
	rand.Seed(time.Now().Unix())
}
func main() {
	n := rand.Intn(999)
	fmt.Println(n)
}
