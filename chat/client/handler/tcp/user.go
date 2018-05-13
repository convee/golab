package tcp

import (
	"net"
)

type TcpHandler struct {
}

func (t *TcpHandler) Login(conn net.Conn, userId int, passwd string) (err error) {
	return
}
