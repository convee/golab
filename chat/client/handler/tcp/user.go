package tcp

import (
	"encoding/binary"
	"encoding/json"
	"errors"
	"fmt"
	"golab/chat/common"
	"golab/chat/proto"
	"net"
	"os"
)

type TcpHandler struct {
	Conn net.Conn
}

func (t *TcpHandler) Login(userId int, passwd string) (err error) {

	loginCmd := proto.LoginCmd{
		Id:     userId,
		Passwd: passwd,
	}
	data, err := json.Marshal(loginCmd)
	if err != nil {
		return
	}
	msg := proto.Message{
		Cmd:  proto.UserLogin,
		Data: string(data),
	}
	data, err = json.Marshal(msg)
	if err != nil {
		return
	}
	var buf [4]byte
	packLen := uint32(len(data))
	//将包长度packLen转成大端序 保存到buf字节数组中
	binary.BigEndian.PutUint32(buf[0:4], packLen)
	n, err := t.Conn.Write(buf[:])
	if err != nil || n != 4 {
		fmt.Println("write data failed")
		return
	}
	msg, err = t.readPackage()
	if err != nil {
		fmt.Println("read package failed,err:", err)
		return
	}
	var loginResp proto.LoginCmdRes
	json.Unmarshal([]byte(msg.Data), &loginResp)
	if loginResp.Code == 500 {
		fmt.Println("user not register, start register")
		//todo register
		os.Exit(0)
	}

	return
}

func (t *TcpHandler) Register(userId int, passwd string) (err error) {
	registerCmd := proto.RegisterCmd{
		User: common.User{
			UserId: userId,
			Passwd: passwd,
			Nick:   fmt.Sprintf("stu%d", userId),
			Sex:    "man",
			Header: "",
		},
	}
	data, err := json.Marshal(registerCmd)
	if err != nil {
		return
	}
	var buf [4]byte
	packLen := uint32(len(data))
	binary.BigEndian.PutUint32(buf[0:4], packLen)
	//长度字节流写入
	n, err := t.Conn.Write(buf[:])
	if err != nil || n != 4 {
		fmt.Println("write data failed")
		return
	}
	//内容字节流写入
	_, err = t.Conn.Write([]byte(data))
	if err != nil {
		return
	}
	msg, err := t.readPackage()
	if err != nil {
		fmt.Println("read package failed,err:", err)
	}
	fmt.Println(msg)

	return
}

func (t *TcpHandler) readPackage() (msg proto.Message, err error) {
	var buf [8192]byte
	n, err := t.Conn.Read(buf[0:4])
	if n != 4 {
		err = errors.New("read head failed")
		return
	}
	packLen := binary.BigEndian.Uint32(buf[0:4])
	n, err = t.Conn.Read(buf[0:packLen])
	if n != int(packLen) {
		err = errors.New("read body failed")
		return
	}
	err = json.Unmarshal(buf[0:packLen], &msg)
	if err != nil {
		return
	}
	return
}
