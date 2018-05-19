package main

import (
	"testing"
	"flag"
	"google.golang.org/grpc"
	"log"
	"go-grpc/normal/pb"
	"golang.org/x/net/context"
	"fmt"
)

var (
	serverAddr = flag.String("server_addr", "127.0.0.1:11000", "The server address in the format of host:port")
	uid = flag.Int64("uid", 40824, "userId")
)

func Test_Normal(t *testing.T)  {
	flag.Parse()
	conn, err := grpc.Dial(*serverAddr, grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	client := pb.NewUserClient(conn)
	userInfo, err := client.GetUserInfo(context.Background(), &pb.Req{Uid:*uid})
	if err != nil {
		log.Fatal(err)
	}
	nick := userInfo.Nick
	fmt.Println(nick)
}