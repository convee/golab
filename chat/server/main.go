package main

import (
	"flag"
	"golab/chat/server/cache"
	"golab/chat/server/handle/tcp"
)

var (
	tcpAddr = flag.String("tcp_addr", "127.0.0.1:5004", "tcp addr")
)

func main() {
	cache.InitRedis()
	tcp.RunServer(*tcpAddr)
}
