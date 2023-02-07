// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.12
// source: test.proto

package pb

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// MessageSenderClient is the client API for MessageSender service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MessageSenderClient interface {
	Send(ctx context.Context, in *MessageRequest, opts ...grpc.CallOption) (*MessageResponse, error)
}

type messageSenderClient struct {
	cc grpc.ClientConnInterface
}

func NewMessageSenderClient(cc grpc.ClientConnInterface) MessageSenderClient {
	return &messageSenderClient{cc}
}

func (c *messageSenderClient) Send(ctx context.Context, in *MessageRequest, opts ...grpc.CallOption) (*MessageResponse, error) {
	out := new(MessageResponse)
	err := c.cc.Invoke(ctx, "/MessageSender/Send", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MessageSenderServer is the server API for MessageSender service.
// All implementations must embed UnimplementedMessageSenderServer
// for forward compatibility
type MessageSenderServer interface {
	Send(context.Context, *MessageRequest) (*MessageResponse, error)
	mustEmbedUnimplementedMessageSenderServer()
}

// UnimplementedMessageSenderServer must be embedded to have forward compatible implementations.
type UnimplementedMessageSenderServer struct {
}

func (UnimplementedMessageSenderServer) Send(context.Context, *MessageRequest) (*MessageResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Send not implemented")
}
func (UnimplementedMessageSenderServer) mustEmbedUnimplementedMessageSenderServer() {}

// UnsafeMessageSenderServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to MessageSenderServer will
// result in compilation errors.
type UnsafeMessageSenderServer interface {
	mustEmbedUnimplementedMessageSenderServer()
}

func RegisterMessageSenderServer(s grpc.ServiceRegistrar, srv MessageSenderServer) {
	s.RegisterService(&MessageSender_ServiceDesc, srv)
}

func _MessageSender_Send_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MessageRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MessageSenderServer).Send(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/MessageSender/Send",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MessageSenderServer).Send(ctx, req.(*MessageRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// MessageSender_ServiceDesc is the grpc.ServiceDesc for MessageSender service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var MessageSender_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "MessageSender",
	HandlerType: (*MessageSenderServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Send",
			Handler:    _MessageSender_Send_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "test.proto",
}

// 并发处理server注册
func RegisterConcurrencyMessageSenderServer(s grpc.ServiceRegistrar, srvs *ServerPool) {
	s.RegisterService(&MessageSender_Concurrency_ServiceDesc, srvs)
}

func _MessageSender_Concurrency_Send_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MessageRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MessageSenderServer).Send(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/MessageSender/Send",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MessageSenderServer).Send(ctx, req.(*MessageRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// MessageSender_ServiceDesc is the grpc.ServiceDesc for MessageSender service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var MessageSender_Concurrency_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "MessageSender",
	HandlerType: ([]*MessageSenderServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Send",
			Handler:    _MessageSender_Concurrency_Send_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "test.proto",
}

type ServerPool struct {
	num  int
	srvs []*chan MessageSenderServer
}

func NewServerPool(num int, ctx context.Context, req *MessageRequest) *ServerPool {
	srvPool := &ServerPool{
		num: num,
	}
	for i := 0; i < num; i++ {
		c := make(chan MessageSenderServer)
		srvPool.srvs = append(srvPool.srvs, &c)
		go func(sender *chan MessageSenderServer) {
			for {
				select {
				case f := <-*sender:
					{
						f.Send(ctx, req)
					}
				default:
				}
			}
		}(&c)
	}
	return srvPool
}
