package main

import (
	"fmt"
	"net"
	"io/ioutil"
	"log"
)

func main()  {

	tcpAddr, err := net.ResolveTCPAddr("tcp4", "127.0.0.1:6666")
	if err != nil {
		log.Fatal(err)
	}
	conn, err := net.DialTCP("tcp", nil, tcpAddr)
	if err != nil {
		log.Fatal(err)
	}
	_, err = conn.Write([]byte("HEAD / HTTP/1.0\n\n"))
	if err != nil {
		log.Fatal(err)
	}
	result, err := ioutil.ReadAll(conn)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(result))
}