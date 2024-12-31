// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v5.29.2
// source: proto/blockchain.proto

package proto

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	Blockchain_SayHello_FullMethodName = "/Blockchain/SayHello"
)

// BlockchainClient is the client API for Blockchain service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type BlockchainClient interface {
	// Sends a greeting
	SayHello(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Reply, error)
}

type blockchainClient struct {
	cc grpc.ClientConnInterface
}

func NewBlockchainClient(cc grpc.ClientConnInterface) BlockchainClient {
	return &blockchainClient{cc}
}

func (c *blockchainClient) SayHello(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Reply, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Reply)
	err := c.cc.Invoke(ctx, Blockchain_SayHello_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// BlockchainServer is the server API for Blockchain service.
// All implementations must embed UnimplementedBlockchainServer
// for forward compatibility.
type BlockchainServer interface {
	// Sends a greeting
	SayHello(context.Context, *Request) (*Reply, error)
	mustEmbedUnimplementedBlockchainServer()
}

// UnimplementedBlockchainServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedBlockchainServer struct{}

func (UnimplementedBlockchainServer) SayHello(context.Context, *Request) (*Reply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SayHello not implemented")
}
func (UnimplementedBlockchainServer) mustEmbedUnimplementedBlockchainServer() {}
func (UnimplementedBlockchainServer) testEmbeddedByValue()                    {}

// UnsafeBlockchainServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to BlockchainServer will
// result in compilation errors.
type UnsafeBlockchainServer interface {
	mustEmbedUnimplementedBlockchainServer()
}

func RegisterBlockchainServer(s grpc.ServiceRegistrar, srv BlockchainServer) {
	// If the following call pancis, it indicates UnimplementedBlockchainServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&Blockchain_ServiceDesc, srv)
}

func _Blockchain_SayHello_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BlockchainServer).SayHello(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Blockchain_SayHello_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BlockchainServer).SayHello(ctx, req.(*Request))
	}
	return interceptor(ctx, in, info, handler)
}

// Blockchain_ServiceDesc is the grpc.ServiceDesc for Blockchain service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Blockchain_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "Blockchain",
	HandlerType: (*BlockchainServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SayHello",
			Handler:    _Blockchain_SayHello_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/blockchain.proto",
}