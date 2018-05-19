package main

import (
	"flag"
	"google.golang.org/grpc"
	"testing"
	"log"
	"go-grpc/stream/pb"
	"golang.org/x/net/context"
	"time"
	"fmt"
)

var(
	serverAddr = flag.String("server_addr", "127.0.0.1:11002", "The Server Addr.")
)
func Test_singleStream(t *testing.T)  {
	conn, err := grpc.Dial(*serverAddr, grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}
	log.Println("grpc dial done")
	client := pb.NewGameClient(conn)
	stream, err := client.Stream(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	log.Println("stream done")
	go func() {
		log.Println("send start")
		defer func(begin time.Time) {
			log.Println("all done, took: ", time.Since(begin))
		}(time.Now())
		for i := 0; i < 10; i++  {
			err := stream.Send(&pb.Message{Body:fmt.Sprintf("body_%d", i)})
			if err != nil {
				log.Fatal(err)
			}
			log.Println("send out")
		}
	}()
	select {}
}