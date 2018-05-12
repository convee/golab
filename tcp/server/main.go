package main

import (
	"flag"
	"fmt"
	"net"
)

var (
	port = flag.String("port", "0.0.0.0:5000", "server port")
)

func process(conn net.Conn) {
	defer conn.Close()
	for {
		buf := make([]byte, 512)
		_, err := conn.Read(buf)
		if err != nil {
			fmt.Println("read err:", err)
			return
		}
		fmt.Println("read:", string(buf))
	}
}

func main() {
	flag.Parse()
	fmt.Println("start server on:", *port)
	listen, err := net.Listen("tcp", *port)
	if err != nil {
		fmt.Println("listen failed, err:", err)
		return
	}
	for {
		conn, err := listen.Accept()
		if err != nil {
			fmt.Println("accept failed, err:", err)
			return
		}
		go process(conn)
	}
}
