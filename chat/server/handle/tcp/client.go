package tcp

import (
	"encoding/binary"
	"encoding/json"
	"errors"
	"fmt"
	"golab/chat/proto"
	"net"
)

type Client struct {
	conn   net.Conn
	userId int
	buf    [8192]byte
}

func (p *Client) readPackage() (msg proto.Message, err error) {
	n, err := p.conn.Read(p.buf[0:4])
	if n != 4 {
		err = errors.New("read header failed")
		return
	}
	fmt.Println("read package:", p.buf[0:4])
	var packLen uint32
	packLen = binary.BigEndian.Uint32(p.buf[0:4])
	fmt.Printf("receive len:%d", packLen)
	n, err = p.conn.Read(p.buf[0:packLen])
	if n != int(packLen) {
		err = errors.New("read body failed")
		return
	}
	fmt.Printf("reveive data:%s\n", string(p.buf[0:packLen]))
	err = json.Unmarshal(p.buf[0:packLen], &msg)
	if err != nil {
		fmt.Println("unmashal failed,err:", err)

	}
	return

}
func (p *Client) login(msg proto.Message) (err error) {
	return
}

func (p *Client) loginResp(err error) {
	return
}
func (p *Client) register(msg proto.Message) (err error) {
	return
}
func (p *Client) processUserSendMessage(msg proto.Message) (err error) {
	return
}
func (p *Client) Process() (err error) {
	for {
		var msg proto.Message
		msg, err = p.readPackage()
		if err != nil {
			return
		}
		err = p.processMsg(msg)
		if err != nil {
			fmt.Println("process msg failed,err:", err)
			continue
		}
	}
}
func (p *Client) processMsg(msg proto.Message) (err error) {
	switch msg.Cmd {
	case proto.UserLogin:
		err = p.login(msg)
	case proto.UserRegister:
		err = p.register(msg)
	case proto.UserSendMessageCmd:
		err = p.processUserSendMessage(msg)
	default:
		err = errors.New("unsupport messgae")
		return
	}
	return
}
