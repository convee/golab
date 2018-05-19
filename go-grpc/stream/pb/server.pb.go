// Code generated by protoc-gen-go. DO NOT EDIT.
// source: server.proto

/*
Package pb is a generated protocol buffer package.

It is generated from these files:
	server.proto

It has these top-level messages:
	Message
*/
package pb

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type Message struct {
	Body string `protobuf:"bytes,1,opt,name=body" json:"body,omitempty"`
}

func (m *Message) Reset()                    { *m = Message{} }
func (m *Message) String() string            { return proto.CompactTextString(m) }
func (*Message) ProtoMessage()               {}
func (*Message) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Message) GetBody() string {
	if m != nil {
		return m.Body
	}
	return ""
}

func init() {
	proto.RegisterType((*Message)(nil), "pb.Message")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for Game service

type GameClient interface {
	Stream(ctx context.Context, opts ...grpc.CallOption) (Game_StreamClient, error)
}

type gameClient struct {
	cc *grpc.ClientConn
}

func NewGameClient(cc *grpc.ClientConn) GameClient {
	return &gameClient{cc}
}

func (c *gameClient) Stream(ctx context.Context, opts ...grpc.CallOption) (Game_StreamClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_Game_serviceDesc.Streams[0], c.cc, "/pb.Game/Stream", opts...)
	if err != nil {
		return nil, err
	}
	x := &gameStreamClient{stream}
	return x, nil
}

type Game_StreamClient interface {
	Send(*Message) error
	Recv() (*Message, error)
	grpc.ClientStream
}

type gameStreamClient struct {
	grpc.ClientStream
}

func (x *gameStreamClient) Send(m *Message) error {
	return x.ClientStream.SendMsg(m)
}

func (x *gameStreamClient) Recv() (*Message, error) {
	m := new(Message)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// Server API for Game service

type GameServer interface {
	Stream(Game_StreamServer) error
}

func RegisterGameServer(s *grpc.Server, srv GameServer) {
	s.RegisterService(&_Game_serviceDesc, srv)
}

func _Game_Stream_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(GameServer).Stream(&gameStreamServer{stream})
}

type Game_StreamServer interface {
	Send(*Message) error
	Recv() (*Message, error)
	grpc.ServerStream
}

type gameStreamServer struct {
	grpc.ServerStream
}

func (x *gameStreamServer) Send(m *Message) error {
	return x.ServerStream.SendMsg(m)
}

func (x *gameStreamServer) Recv() (*Message, error) {
	m := new(Message)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _Game_serviceDesc = grpc.ServiceDesc{
	ServiceName: "pb.Game",
	HandlerType: (*GameServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Stream",
			Handler:       _Game_Stream_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "server.proto",
}

func init() { proto.RegisterFile("server.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 110 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x29, 0x4e, 0x2d, 0x2a,
	0x4b, 0x2d, 0xd2, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x2a, 0x48, 0x52, 0x92, 0xe5, 0x62,
	0xf7, 0x4d, 0x2d, 0x2e, 0x4e, 0x4c, 0x4f, 0x15, 0x12, 0xe2, 0x62, 0x49, 0xca, 0x4f, 0xa9, 0x94,
	0x60, 0x54, 0x60, 0xd4, 0xe0, 0x0c, 0x02, 0xb3, 0x8d, 0x0c, 0xb8, 0x58, 0xdc, 0x13, 0x73, 0x53,
	0x85, 0x34, 0xb8, 0xd8, 0x82, 0x4b, 0x8a, 0x52, 0x13, 0x73, 0x85, 0xb8, 0xf5, 0x0a, 0x92, 0xf4,
	0xa0, 0x5a, 0xa4, 0x90, 0x39, 0x4a, 0x0c, 0x1a, 0x8c, 0x06, 0x8c, 0x49, 0x6c, 0x60, 0xb3, 0x8d,
	0x01, 0x01, 0x00, 0x00, 0xff, 0xff, 0x19, 0x93, 0xaf, 0xca, 0x6b, 0x00, 0x00, 0x00,
}