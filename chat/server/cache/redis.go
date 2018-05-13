package cache

import (
	"github.com/garyburd/redigo/redis"
)

var pool *redis.Pool

func InitRedis() {
	pool = &redis.Pool{
		MaxIdle:     16,
		MaxActive:   1024,
		IdleTimeout: 300,
		Dial: func() (redis.Conn, error) {
			return redis.Dial("tcp", "106.15.205.6:6379", redis.DialPassword("wang1019"))
		},
	}
	return
}

func GetConn() redis.Conn {
	return pool.Get()
}

func PutConn(conn redis.Conn) {
	conn.Close()
}
