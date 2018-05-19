package main

import (
	"flag"
	"go-grpc/router_guide/pb"
	"io"
	"net"
	"log"
	"google.golang.org/grpc"
)

var (
	addr = flag.String("addr", ":11001", "The server address")
)

type routeGuideServer struct {}

func (s *routeGuideServer) RouteChat(stream pb.RouteGuide_RouteChatServer) error {
	for{
		recv, err := stream.Recv()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			return err
		}
		log.Println(recv.Message)
	}
}

func main() {
	flag.Parse()
	listen, err := net.Listen("tcp", *addr)
	if err !=nil {
		log.Fatal(err)
	}
	//运行一个grpc服务器，监听来自客户端的请求并返回服务的响应
	grpcServer := grpc.NewServer()
	grpcHandler := new(routeGuideServer)
	pb.RegisterRouteGuideServer(grpcServer, grpcHandler)
	grpcServer.Serve(listen)
}