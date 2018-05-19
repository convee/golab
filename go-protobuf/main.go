package main

import (
	"github.com/golang/protobuf/proto"
	"go-protobuf/pb"
	"log"
)

func main()  {
	p := &pb.Person{
		Id: proto.Int32(1),
		Name: proto.String("wangkang"),
	}

	data, err := proto.Marshal(p)
	if err != nil {
		log.Fatal("marshaling error:", err)
	}

	newTest := &pb.Person{}
	err = proto.Unmarshal(data, newTest)

	if err != nil {
		log.Fatal("unmarshaling error:", err)
	}

	if p.GetId() != newTest.GetId() {
		log.Fatalf("data dismatcha %q != %q", p.GetId(), newTest.GetId())
	}

}