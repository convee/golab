package main

import (
	"fmt"

	"github.com/garyburd/redigo/redis"
)

var pool *redis.Pool

func init() {
	pool = &redis.Pool{
		MaxIdle:     16,
		MaxActive:   0,
		IdleTimeout: 300,
		Dial: func() (redis.Conn, error) {
			return redis.Dial("tcp", "127.0.0.1:6379", redis.DialPassword("123456"))
		},
	}
}
func main() {
	c := pool.Get()
	defer c.Close()
	_, err := c.Do("Set", "name", "convee")
	if err != nil {
		fmt.Println(err)
		return
	}
	r, err := redis.String(c.Do("Get", "name"))
	if err != nil {
		fmt.Println("get name failed,", err)
		return
	}
	fmt.Println(r)
	pool.Close()
}
