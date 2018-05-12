package main

import (
	"bufio"
	"flag"
	"fmt"
	"net"
	"os"
	"strings"
)

var (
	port = flag.String("port", "0.0.0.0:5000", "server port")
)

func main() {
	conn, err := net.Dial("tcp", *port)
	if err != nil {
		fmt.Println("dialing error", err.Error())
		return
	}
	defer conn.Close()
	inputReader := bufio.NewReader(os.Stdin)
	for {
		input, _ := inputReader.ReadString('\n')
		trimmedInput := strings.Trim(input, "\r\n")
		if trimmedInput == "Q" {
			return
		}
		_, err := conn.Write([]byte(trimmedInput))
		if err != nil {
			return
		}
	}
}
