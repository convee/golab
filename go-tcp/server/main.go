package main

import (
	"net"
	"fmt"
	"time"
	"strings"
	"strconv"
	"log"
)

func main()  {

	address := net.TCPAddr{
		IP:net.ParseIP("127.0.0.1"),
		Port:6666,
	}

	listener, err := net.ListenTCP("tcp4", &address) //创建tcp4服务器监听器

	if err != nil {
		log.Fatal(err)
	}
	for {
		conn, err := listener.AcceptTCP()
		if err != nil {
			continue
		}
		go handleClient(conn)
	}
}

func handleClient(conn *net.TCPConn)  {
	conn.SetReadDeadline(time.Now().Add(2 * time.Minute)) //设置2分钟超时
	request := make([]byte, 128) //创建request字节数组，设置最大请求长度为128b，防止dos攻击
	defer conn.Close()
	for {
		readLen, err := conn.Read(request)
		if err != nil {
			fmt.Println(err)
			break
		}
		if readLen == 0 {
			break //客户端已经关闭连接
		} else if strings.TrimSpace(string(request[:readLen])) == "timestamp" {
			daytime := strconv.FormatInt(time.Now().Unix(), 10)
			conn.Write([]byte(daytime))
		} else {
			daytime := time.Now().String()
			conn.Write([]byte(daytime))
		}
		request = make([]byte, 128) //清除最后读取的内容
	}
}