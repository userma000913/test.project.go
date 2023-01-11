// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.8
// source: hello_grpc.proto

package say_hello

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

// SayHelloServiceClient is the client API for SayHelloService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type SayHelloServiceClient interface {
	SayHello(ctx context.Context, in *SayHelloRequest, opts ...grpc.CallOption) (*SayHelloResponse, error)
}

type sayHelloServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewSayHelloServiceClient(cc grpc.ClientConnInterface) SayHelloServiceClient {
	return &sayHelloServiceClient{cc}
}

func (c *sayHelloServiceClient) SayHello(ctx context.Context, in *SayHelloRequest, opts ...grpc.CallOption) (*SayHelloResponse, error) {
	out := new(SayHelloResponse)
	err := c.cc.Invoke(ctx, "/say_hello.SayHelloService/SayHello", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SayHelloServiceServer is the server API for SayHelloService service.
// All implementations must embed UnimplementedSayHelloServiceServer
// for forward compatibility
type SayHelloServiceServer interface {
	SayHello(context.Context, *SayHelloRequest) (*SayHelloResponse, error)
	mustEmbedUnimplementedSayHelloServiceServer()
}

// UnimplementedSayHelloServiceServer must be embedded to have forward compatible implementations.
type UnimplementedSayHelloServiceServer struct {
}

func (UnimplementedSayHelloServiceServer) SayHello(context.Context, *SayHelloRequest) (*SayHelloResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SayHello not implemented")
}
func (UnimplementedSayHelloServiceServer) mustEmbedUnimplementedSayHelloServiceServer() {}

// UnsafeSayHelloServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to SayHelloServiceServer will
// result in compilation errors.
type UnsafeSayHelloServiceServer interface {
	mustEmbedUnimplementedSayHelloServiceServer()
}

func RegisterSayHelloServiceServer(s grpc.ServiceRegistrar, srv SayHelloServiceServer) {
	s.RegisterService(&SayHelloService_ServiceDesc, srv)
}

func _SayHelloService_SayHello_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SayHelloRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SayHelloServiceServer).SayHello(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/say_hello.SayHelloService/SayHello",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SayHelloServiceServer).SayHello(ctx, req.(*SayHelloRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// SayHelloService_ServiceDesc is the grpc.ServiceDesc for SayHelloService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var SayHelloService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "say_hello.SayHelloService",
	HandlerType: (*SayHelloServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SayHello",
			Handler:    _SayHelloService_SayHello_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "hello_grpc.proto",
}
