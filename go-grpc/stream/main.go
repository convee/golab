package main

import (
	"flag"
	"golab/go-grpc/stream/pb"
	"log"
	"net"
	"time"

	"google.golang.org/grpc"
)

var (
	addr = flag.String("addr", ":11002", "The Server Addr")
)

type server struct{}

func (s *server) Stream(stream pb.Game_StreamServer) error {
	log.Println("new stream")
	defer func(begin time.Time) {
		log.Println("all done, took: ", time.Since(begin))
	}(time.Now())
	count := 0
	for {
		recv, err := stream.Recv()
		log.Println("stream recv")
		if err != nil {
			log.Println(err)
			return err
		}
		stream.Send(recv)
		log.Println("send out")
		count++
		if count >= 10 {
			break
		}
	}
	return nil
}

func main() {
	listen, err := net.Listen("tcp", *addr)
	if err != nil {
		log.Fatal(err)
	}
	grpcServer := grpc.NewServer()
	grpcHandler := new(server)
	pb.RegisterGameServer(grpcServer, grpcHandler)
	grpcServer.Serve(listen)
}
