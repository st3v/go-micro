// Code generated by protoc-gen-go.
// source: go-micro/examples/server/proto/example/example.proto
// DO NOT EDIT!

/*
Package go_micro_srv_example is a generated protocol buffer package.

It is generated from these files:
	go-micro/examples/server/proto/example/example.proto

It has these top-level messages:
	Message
	Request
	Response
	StreamingRequest
	StreamingResponse
	Ping
	Pong
*/
package go_micro_srv_example

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	client "github.com/micro/go-micro/client"
	server "github.com/micro/go-micro/server"
	context "golang.org/x/net/context"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type Message struct {
	Say string `protobuf:"bytes,1,opt,name=say" json:"say,omitempty"`
}

func (m *Message) Reset()                    { *m = Message{} }
func (m *Message) String() string            { return proto.CompactTextString(m) }
func (*Message) ProtoMessage()               {}
func (*Message) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type Request struct {
	Name string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
}

func (m *Request) Reset()                    { *m = Request{} }
func (m *Request) String() string            { return proto.CompactTextString(m) }
func (*Request) ProtoMessage()               {}
func (*Request) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

type Response struct {
	Msg string `protobuf:"bytes,1,opt,name=msg" json:"msg,omitempty"`
}

func (m *Response) Reset()                    { *m = Response{} }
func (m *Response) String() string            { return proto.CompactTextString(m) }
func (*Response) ProtoMessage()               {}
func (*Response) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

type StreamingRequest struct {
	Count int64 `protobuf:"varint,1,opt,name=count" json:"count,omitempty"`
}

func (m *StreamingRequest) Reset()                    { *m = StreamingRequest{} }
func (m *StreamingRequest) String() string            { return proto.CompactTextString(m) }
func (*StreamingRequest) ProtoMessage()               {}
func (*StreamingRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

type StreamingResponse struct {
	Count int64 `protobuf:"varint,1,opt,name=count" json:"count,omitempty"`
}

func (m *StreamingResponse) Reset()                    { *m = StreamingResponse{} }
func (m *StreamingResponse) String() string            { return proto.CompactTextString(m) }
func (*StreamingResponse) ProtoMessage()               {}
func (*StreamingResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

type Ping struct {
	Stroke int64 `protobuf:"varint,1,opt,name=stroke" json:"stroke,omitempty"`
}

func (m *Ping) Reset()                    { *m = Ping{} }
func (m *Ping) String() string            { return proto.CompactTextString(m) }
func (*Ping) ProtoMessage()               {}
func (*Ping) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

type Pong struct {
	Stroke int64 `protobuf:"varint,1,opt,name=stroke" json:"stroke,omitempty"`
}

func (m *Pong) Reset()                    { *m = Pong{} }
func (m *Pong) String() string            { return proto.CompactTextString(m) }
func (*Pong) ProtoMessage()               {}
func (*Pong) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

func init() {
	proto.RegisterType((*Message)(nil), "go.micro.srv.example.Message")
	proto.RegisterType((*Request)(nil), "go.micro.srv.example.Request")
	proto.RegisterType((*Response)(nil), "go.micro.srv.example.Response")
	proto.RegisterType((*StreamingRequest)(nil), "go.micro.srv.example.StreamingRequest")
	proto.RegisterType((*StreamingResponse)(nil), "go.micro.srv.example.StreamingResponse")
	proto.RegisterType((*Ping)(nil), "go.micro.srv.example.Ping")
	proto.RegisterType((*Pong)(nil), "go.micro.srv.example.Pong")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ client.Option
var _ server.Option

// Client API for Example service

type ExampleClient interface {
	Call(ctx context.Context, in *Request, opts ...client.CallOption) (*Response, error)
	Stream(ctx context.Context, in *StreamingRequest, opts ...client.CallOption) (Example_StreamClient, error)
	PingPong(ctx context.Context, opts ...client.CallOption) (Example_PingPongClient, error)
}

type exampleClient struct {
	c           client.Client
	serviceName string
}

func NewExampleClient(serviceName string, c client.Client) ExampleClient {
	if c == nil {
		c = client.NewClient()
	}
	if len(serviceName) == 0 {
		serviceName = "go.micro.srv.example"
	}
	return &exampleClient{
		c:           c,
		serviceName: serviceName,
	}
}

func (c *exampleClient) Call(ctx context.Context, in *Request, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.serviceName, "Example.Call", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *exampleClient) Stream(ctx context.Context, in *StreamingRequest, opts ...client.CallOption) (Example_StreamClient, error) {
	req := c.c.NewRequest(c.serviceName, "Example.Stream", &StreamingRequest{})
	stream, err := c.c.Stream(ctx, req, opts...)
	if err != nil {
		return nil, err
	}
	if err := stream.Send(in); err != nil {
		return nil, err
	}
	return &exampleStreamClient{stream}, nil
}

type Example_StreamClient interface {
	SendMsg(interface{}) error
	RecvMsg(interface{}) error
	Close() error
	Recv() (*StreamingResponse, error)
}

type exampleStreamClient struct {
	stream client.Streamer
}

func (x *exampleStreamClient) Close() error {
	return x.stream.Close()
}

func (x *exampleStreamClient) SendMsg(m interface{}) error {
	return x.stream.Send(m)
}

func (x *exampleStreamClient) RecvMsg(m interface{}) error {
	return x.stream.Recv(m)
}

func (x *exampleStreamClient) Recv() (*StreamingResponse, error) {
	m := new(StreamingResponse)
	err := x.stream.Recv(m)
	if err != nil {
		return nil, err
	}
	return m, nil
}

func (c *exampleClient) PingPong(ctx context.Context, opts ...client.CallOption) (Example_PingPongClient, error) {
	req := c.c.NewRequest(c.serviceName, "Example.PingPong", &Ping{})
	stream, err := c.c.Stream(ctx, req, opts...)
	if err != nil {
		return nil, err
	}
	return &examplePingPongClient{stream}, nil
}

type Example_PingPongClient interface {
	SendMsg(interface{}) error
	RecvMsg(interface{}) error
	Close() error
	Send(*Ping) error
	Recv() (*Pong, error)
}

type examplePingPongClient struct {
	stream client.Streamer
}

func (x *examplePingPongClient) Close() error {
	return x.stream.Close()
}

func (x *examplePingPongClient) SendMsg(m interface{}) error {
	return x.stream.Send(m)
}

func (x *examplePingPongClient) RecvMsg(m interface{}) error {
	return x.stream.Recv(m)
}

func (x *examplePingPongClient) Send(m *Ping) error {
	return x.stream.Send(m)
}

func (x *examplePingPongClient) Recv() (*Pong, error) {
	m := new(Pong)
	err := x.stream.Recv(m)
	if err != nil {
		return nil, err
	}
	return m, nil
}

// Server API for Example service

type ExampleHandler interface {
	Call(context.Context, *Request, *Response) error
	Stream(context.Context, *StreamingRequest, Example_StreamStream) error
	PingPong(context.Context, Example_PingPongStream) error
}

func RegisterExampleHandler(s server.Server, hdlr ExampleHandler) {
	s.Handle(s.NewHandler(&Example{hdlr}))
}

type Example struct {
	ExampleHandler
}

func (h *Example) Call(ctx context.Context, in *Request, out *Response) error {
	return h.ExampleHandler.Call(ctx, in, out)
}

func (h *Example) Stream(ctx context.Context, stream server.Streamer) error {
	m := new(StreamingRequest)
	if err := stream.Recv(m); err != nil {
		return err
	}
	return h.ExampleHandler.Stream(ctx, m, &exampleStreamStream{stream})
}

type Example_StreamStream interface {
	SendMsg(interface{}) error
	RecvMsg(interface{}) error
	Close() error
	Send(*StreamingResponse) error
}

type exampleStreamStream struct {
	stream server.Streamer
}

func (x *exampleStreamStream) Close() error {
	return x.stream.Close()
}

func (x *exampleStreamStream) SendMsg(m interface{}) error {
	return x.stream.Send(m)
}

func (x *exampleStreamStream) RecvMsg(m interface{}) error {
	return x.stream.Recv(m)
}

func (x *exampleStreamStream) Send(m *StreamingResponse) error {
	return x.stream.Send(m)
}

func (h *Example) PingPong(ctx context.Context, stream server.Streamer) error {
	return h.ExampleHandler.PingPong(ctx, &examplePingPongStream{stream})
}

type Example_PingPongStream interface {
	SendMsg(interface{}) error
	RecvMsg(interface{}) error
	Close() error
	Send(*Pong) error
	Recv() (*Ping, error)
}

type examplePingPongStream struct {
	stream server.Streamer
}

func (x *examplePingPongStream) Close() error {
	return x.stream.Close()
}

func (x *examplePingPongStream) SendMsg(m interface{}) error {
	return x.stream.Send(m)
}

func (x *examplePingPongStream) RecvMsg(m interface{}) error {
	return x.stream.Recv(m)
}

func (x *examplePingPongStream) Send(m *Pong) error {
	return x.stream.Send(m)
}

func (x *examplePingPongStream) Recv() (*Ping, error) {
	m := new(Ping)
	if err := x.stream.Recv(m); err != nil {
		return nil, err
	}
	return m, nil
}

var fileDescriptor0 = []byte{
	// 270 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x84, 0x91, 0x5f, 0x4b, 0xc3, 0x30,
	0x14, 0xc5, 0x17, 0x56, 0xdb, 0x79, 0xfd, 0x83, 0x06, 0x99, 0x52, 0x50, 0x34, 0x0f, 0xba, 0x17,
	0xd3, 0xa1, 0x7e, 0x03, 0x11, 0x7d, 0x11, 0x64, 0x3e, 0xfb, 0x10, 0xc7, 0x25, 0x0c, 0x9b, 0xa6,
	0xe6, 0x66, 0x43, 0x3f, 0xbb, 0x2f, 0x6e, 0x69, 0x3b, 0xc6, 0xec, 0xf0, 0x29, 0x70, 0x7e, 0xe7,
	0x5c, 0xce, 0x21, 0x70, 0xa7, 0xed, 0xb5, 0x99, 0x8c, 0x9d, 0xcd, 0xf0, 0x4b, 0x99, 0x32, 0x47,
	0xca, 0x08, 0xdd, 0x0c, 0x5d, 0x56, 0x3a, 0xeb, 0x97, 0x6a, 0xf3, 0xca, 0xa0, 0xf2, 0x23, 0x6d,
	0x65, 0x48, 0x49, 0x72, 0x33, 0x59, 0x33, 0xd1, 0x87, 0xe4, 0x19, 0x89, 0x94, 0x46, 0xbe, 0x03,
	0x5d, 0x52, 0xdf, 0x27, 0xec, 0x9c, 0x0d, 0xb6, 0xc5, 0x31, 0x24, 0x23, 0xfc, 0x9c, 0x22, 0x79,
	0xbe, 0x0b, 0x51, 0xa1, 0x0c, 0x2e, 0x41, 0x6f, 0x84, 0x54, 0xda, 0x82, 0x42, 0xc2, 0x90, 0xae,
	0xc1, 0x05, 0x1c, 0xbc, 0x7a, 0x87, 0xca, 0x4c, 0x0a, 0xdd, 0x44, 0xf7, 0x60, 0x6b, 0x6c, 0xa7,
	0x85, 0x0f, 0x96, 0xae, 0x10, 0x70, 0xb8, 0x62, 0xa9, 0x8f, 0xac, 0x79, 0xfa, 0x10, 0xbd, 0xcc,
	0x31, 0xdf, 0x87, 0x98, 0xbc, 0xb3, 0x1f, 0xb8, 0xa2, 0xdb, 0xbf, 0xfa, 0xcd, 0x0f, 0x83, 0xe4,
	0xa1, 0x1a, 0xc3, 0x1f, 0x21, 0xba, 0x57, 0x79, 0xce, 0x4f, 0x65, 0xdb, 0x56, 0x59, 0xb7, 0x4a,
	0xcf, 0x36, 0xe1, 0xaa, 0x91, 0xe8, 0xf0, 0x37, 0x88, 0xab, 0xa2, 0xfc, 0xb2, 0xdd, 0xbb, 0xbe,
	0x34, 0xbd, 0xfa, 0xd7, 0xd7, 0x1c, 0x1f, 0x32, 0xfe, 0x04, 0xbd, 0xc5, 0xc6, 0xb0, 0x27, 0x6d,
	0x0f, 0x2e, 0x78, 0xba, 0x89, 0xcd, 0x73, 0xa2, 0x33, 0x60, 0x43, 0xf6, 0x1e, 0x87, 0xbf, 0xbd,
	0xfd, 0x0d, 0x00, 0x00, 0xff, 0xff, 0x53, 0xb5, 0xeb, 0x31, 0x13, 0x02, 0x00, 0x00,
}
