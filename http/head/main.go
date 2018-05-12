package main

import (
	"fmt"
	"net"
	"net/http"
	"time"
)

var url = []string{
	"http://www.baidu.com",
	"http://www.google.com",
	"http://www.douyu.com",
}

func main() {
	//设置超时的http客户端
	c := http.Client{
		Transport: &http.Transport{
			Dial: func(network, addr string) (net.Conn, error) {
				return net.DialTimeout(network, addr, time.Second)
			},
		},
	}
	for _, v := range url {
		resp, err := c.Head(v)
		if err != nil {
			fmt.Printf("head %s is failed,err:%v\n", v, err)
			continue
		}
		fmt.Printf("head succ,status:%v\n", resp.Status)
	}
}
