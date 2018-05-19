package main

import (
	"testing"
	"flag"
	"google.golang.org/grpc"
	"log"
	"go-grpc/router_guide/pb"
	"golang.org/x/net/context"
	"io"
	"fmt"
	"os"
	"os/signal"
	"syscall"
)


var serverAddr = flag.String("server_addr", "127.0.0.1:11001", "The server address in the format of host:port")

func Test_RouteGuide(t *testing.T)  {
	flag.Parse()
	conn, err := grpc.Dial(*serverAddr, grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	client := pb.NewRouteGuideClient(conn)
	stream, err := client.RouteChat(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	log.Println("client start")
	go func() {
		for {
			err := stream.Send(&pb.RouteNote{Message:"666", Uid:40824})
			log.Println(err)
			if err != nil {
				log.Println(err)
			}
			fmt.Println("server start2")
			in, err := stream.Recv()
			//io.EOF 输入结束
			if err == io.EOF {
				log.Println(err)
			}
			if err != nil {
				log.Println(err)
			}
			fmt.Println("server: ", in.Uid, in.Message)
		}
	}()

	go func() {
		c := make(chan os.Signal)
		signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
		log.Fatal("client over")
	}()

}