package main

import (
	"net"
	"fmt"
	"os"
	"tcp-proto/pb"
	"github.com/golang/protobuf/proto"
	"log"
)
/**
客户端发送线程
发送连接 conn
 */
func chatSend(conn net.Conn)  {
	var input string
	username := conn.LocalAddr().String()
	for {
		fmt.Scanln(&input)
		if input == "/quit" {
			fmt.Println("ByeBye..")
			conn.Close()
			os.Exit(0)
		}
		msg := &pb.Chat{
			User:proto.String(username),
			Message:proto.String(input),
		}
		//进行编码
		data, err := proto.Marshal(msg)
		if err !=nil {
			conn.Close()
			os.Exit(1)
		}
		//发送服务端
		length, err := conn.Write(data)
		if length > 0 {

		}
		if err != nil {
			fmt.Println(err.Error())
			conn.Close()
			break
		}

	}
}

func main()  {
	tcpAddr, err := net.ResolveTCPAddr("tcp4", "127.0.0.1:7777")
	if err != nil {
		log.Fatal(err)
	}
	conn, err := net.DialTCP("tcp", nil, tcpAddr)
	if err != nil {
		log.Fatal(err)
	}
	//启动客户端发送线程
	go chatSend(conn)

	//接收数据
	buf := make([]byte, 1024)
	for {
		length, err := conn.Read(buf)
		if err != nil {
			conn.Close()
			log.Fatal(err)
		}
		msg := &pb.Chat{}
		err = proto.Unmarshal(buf[0:length], msg)
		if err !=nil {
			conn.Close()
			log.Fatal(err)
		}
		fmt.Println(msg.GetUser() + " " + msg.GetMessage())
	}

}