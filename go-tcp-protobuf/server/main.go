package main

import (
	"net"
	"log"
	"fmt"
)


/**
服务端接收数据线程
客户端 conn
消息队列 message
 */
func Handler(conn net.Conn, messages chan string)  {
	buf := make([]byte, 1024)//设置最大长度为1024,防止攻击
	for {
		//读取数据
		length, err := conn.Read(buf)
		if err != nil {
			conn.Close()
			log.Fatal(err)
		}
		if length > 0 {
			buf[length] = 0
		}
		receiveStr := string(buf[0:length])
		//数据转发给所有客户端
		messages <- receiveStr

	}
}
/**
服务器发送数据的线程
客户端集合 conns
消息队列 messages
 */
func echoHandler(conns *map[string]net.Conn, messages chan string)  {
	//遍历数据
	for {
		msg := <- messages
		//遍历客户端
		for key, value := range *conns {
			_, err := value.Write([]byte(msg))
			if err != nil {
				delete(*conns, key)
				fmt.Println(err.Error())
			}
		}
	}
}

func main()  {
	service := "127.0.0.1:7777"
	tcpAddr, err := net.ResolveTCPAddr("tcp4", service)
	if err != nil {
		log.Fatal(err)
	}
	listener, err := net.ListenTCP("tcp", tcpAddr)
	if err != nil {
		log.Fatal(err)
	}
	//客户端集合
	conns := make(map[string]net.Conn)
	//发送消息队列
	messages := make(chan string, 10)
	//启动发送消息线程
	go echoHandler(&conns, messages)

	for {
		//新客户端
		conn, err := listener.Accept()
		if err != nil {
			log.Fatal(err)
		}
		conns[conn.RemoteAddr().String()] = conn
		//启动接收消息线程
		go Handler(conn, messages)
	}
}