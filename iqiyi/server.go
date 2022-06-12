package main

import (
	"context"
	"log"
	"net/http"
	"strconv"

	"github.com/go-redis/redis/v8"
)

func pv(w http.ResponseWriter, r *http.Request) {
	pv, _ := redisClient.HGet(context.Background(), "PV_COUNT_200000", "211111").Int()
	log.Println(100)
	w.Write([]byte(strconv.Itoa(pv)))
}

// RedisClient redis 客户端
var redisClient *redis.Client

func init() {
	redisClient = redis.NewClient(&redis.Options{
		Addr:         "127.0.0.1:6379",
		Password:     "",
		DB:           0,
		MinIdleConns: 200,
		PoolSize:     100,
	})

	_, err := redisClient.Ping(context.Background()).Result()
	if err != nil {
		panic(err)
	}
}

func main() {
	http.HandleFunc("/pv.go", pv)
	log.Println("http listen 9091")
	err := http.ListenAndServe("0.0.0.0:9091", nil)
	if err != nil {
		log.Fatalln("http listen failed")
	}
}
