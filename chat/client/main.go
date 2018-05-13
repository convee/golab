package main

import (
	"flag"
	"fmt"
	"golab/chat/client/handler/tcp"
	"golab/chat/proto"
	"net"
)

var (
	msgChan chan proto.UserRecvMessageReq
	tcpAddr = flag.String("tcp_addr", "127.0.0.1:5004", "tcp server addr")
)

func init() {
	msgChan = make(chan proto.UserRecvMessageReq, 1000)
}

func main() {
	var userId int
	var passwd string
	fmt.Scanf("%d %s\n", &userId, &passwd)
	conn, err := net.Dial("tcp", *tcpAddr)
	if err != nil {
		fmt.Println("Dialing err,", err)
		return
	}
	tcpHandler := &tcp.TcpHandler{}
	err = tcpHandler.Login(conn, userId, passwd)
	if err != nil {
		fmt.Println("login failed,err:", err)
		return
	}

}
