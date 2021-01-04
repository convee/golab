package main

import (
	"flag"
	"go-grpc/normal/pb"
	"log"
	"net"

	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

var (
	addr = flag.String("port", ":11000", "The server addr.")
)

type userServer struct {
}

func (u userServer) GetUserInfo(ctx context.Context, req *pb.Req) (res *pb.Res, err error) {
	res = &pb.Res{Nick: "convee"}
	return
}
func main() {
	flag.Parse()
	listen, err := net.Listen("tcp", *addr)
	if err != nil {
		log.Fatal(err)
	}
	grpcServer := grpc.NewServer()
	grpcHandler := new(userServer)
	pb.RegisterUserServer(grpcServer, grpcHandler)
	grpcServer.Serve(listen)
}
