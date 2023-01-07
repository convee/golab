package main

import (
	"fmt"
	"time"

	"github.com/robfig/cron/v3"
)

func main() {
	c := cron.New()
	fmt.Println("cron start")
	addFunc, err := c.AddFunc("* * * * *", func() {
		fmt.Println("cron")
		for i := 0; i < 4; i++ {
			go task(i)
		}
	})
	fmt.Println(addFunc)
	if err != nil {
		fmt.Println(err.Error())

	}
	c.Start()
	time.Sleep(2 * time.Minute)
	fmt.Println("cron end")

}

func task(i int) {
	fmt.Println(i)
}
