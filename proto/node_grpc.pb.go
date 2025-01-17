// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v5.29.2
// source: proto/node.proto

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
	Node_ResisterNode_FullMethodName   = "/Node/ResisterNode"
	Node_Sync_FullMethodName           = "/Node/Sync"
	Node_Upload_FullMethodName         = "/Node/Upload"
	Node_Bloadcast_FullMethodName      = "/Node/Bloadcast"
	Node_BloadcastBlock_FullMethodName = "/Node/BloadcastBlock"
)

// NodeClient is the client API for Node service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type NodeClient interface {
	ResisterNode(ctx context.Context, in *ClientInfo, opts ...grpc.CallOption) (*NodeInfo, error)
	Sync(ctx context.Context, in *SyncInfo, opts ...grpc.CallOption) (*SyncReply, error)
	Upload(ctx context.Context, opts ...grpc.CallOption) (grpc.ClientStreamingClient[FileChunk, UploadStatus], error)
	Bloadcast(ctx context.Context, in *Transaction, opts ...grpc.CallOption) (*Verify, error)
	BloadcastBlock(ctx context.Context, in *Block, opts ...grpc.CallOption) (*VerifyBlock, error)
}

type nodeClient struct {
	cc grpc.ClientConnInterface
}

func NewNodeClient(cc grpc.ClientConnInterface) NodeClient {
	return &nodeClient{cc}
}

func (c *nodeClient) ResisterNode(ctx context.Context, in *ClientInfo, opts ...grpc.CallOption) (*NodeInfo, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(NodeInfo)
	err := c.cc.Invoke(ctx, Node_ResisterNode_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *nodeClient) Sync(ctx context.Context, in *SyncInfo, opts ...grpc.CallOption) (*SyncReply, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(SyncReply)
	err := c.cc.Invoke(ctx, Node_Sync_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *nodeClient) Upload(ctx context.Context, opts ...grpc.CallOption) (grpc.ClientStreamingClient[FileChunk, UploadStatus], error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	stream, err := c.cc.NewStream(ctx, &Node_ServiceDesc.Streams[0], Node_Upload_FullMethodName, cOpts...)
	if err != nil {
		return nil, err
	}
	x := &grpc.GenericClientStream[FileChunk, UploadStatus]{ClientStream: stream}
	return x, nil
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type Node_UploadClient = grpc.ClientStreamingClient[FileChunk, UploadStatus]

func (c *nodeClient) Bloadcast(ctx context.Context, in *Transaction, opts ...grpc.CallOption) (*Verify, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Verify)
	err := c.cc.Invoke(ctx, Node_Bloadcast_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *nodeClient) BloadcastBlock(ctx context.Context, in *Block, opts ...grpc.CallOption) (*VerifyBlock, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(VerifyBlock)
	err := c.cc.Invoke(ctx, Node_BloadcastBlock_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// NodeServer is the server API for Node service.
// All implementations must embed UnimplementedNodeServer
// for forward compatibility.
type NodeServer interface {
	ResisterNode(context.Context, *ClientInfo) (*NodeInfo, error)
	Sync(context.Context, *SyncInfo) (*SyncReply, error)
	Upload(grpc.ClientStreamingServer[FileChunk, UploadStatus]) error
	Bloadcast(context.Context, *Transaction) (*Verify, error)
	BloadcastBlock(context.Context, *Block) (*VerifyBlock, error)
	mustEmbedUnimplementedNodeServer()
}

// UnimplementedNodeServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedNodeServer struct{}

func (UnimplementedNodeServer) ResisterNode(context.Context, *ClientInfo) (*NodeInfo, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ResisterNode not implemented")
}
func (UnimplementedNodeServer) Sync(context.Context, *SyncInfo) (*SyncReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Sync not implemented")
}
func (UnimplementedNodeServer) Upload(grpc.ClientStreamingServer[FileChunk, UploadStatus]) error {
	return status.Errorf(codes.Unimplemented, "method Upload not implemented")
}
func (UnimplementedNodeServer) Bloadcast(context.Context, *Transaction) (*Verify, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Bloadcast not implemented")
}
func (UnimplementedNodeServer) BloadcastBlock(context.Context, *Block) (*VerifyBlock, error) {
	return nil, status.Errorf(codes.Unimplemented, "method BloadcastBlock not implemented")
}
func (UnimplementedNodeServer) mustEmbedUnimplementedNodeServer() {}
func (UnimplementedNodeServer) testEmbeddedByValue()              {}

// UnsafeNodeServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to NodeServer will
// result in compilation errors.
type UnsafeNodeServer interface {
	mustEmbedUnimplementedNodeServer()
}

func RegisterNodeServer(s grpc.ServiceRegistrar, srv NodeServer) {
	// If the following call pancis, it indicates UnimplementedNodeServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&Node_ServiceDesc, srv)
}

func _Node_ResisterNode_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ClientInfo)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NodeServer).ResisterNode(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Node_ResisterNode_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NodeServer).ResisterNode(ctx, req.(*ClientInfo))
	}
	return interceptor(ctx, in, info, handler)
}

func _Node_Sync_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SyncInfo)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NodeServer).Sync(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Node_Sync_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NodeServer).Sync(ctx, req.(*SyncInfo))
	}
	return interceptor(ctx, in, info, handler)
}

func _Node_Upload_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(NodeServer).Upload(&grpc.GenericServerStream[FileChunk, UploadStatus]{ServerStream: stream})
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type Node_UploadServer = grpc.ClientStreamingServer[FileChunk, UploadStatus]

func _Node_Bloadcast_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Transaction)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NodeServer).Bloadcast(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Node_Bloadcast_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NodeServer).Bloadcast(ctx, req.(*Transaction))
	}
	return interceptor(ctx, in, info, handler)
}

func _Node_BloadcastBlock_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Block)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NodeServer).BloadcastBlock(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Node_BloadcastBlock_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NodeServer).BloadcastBlock(ctx, req.(*Block))
	}
	return interceptor(ctx, in, info, handler)
}

// Node_ServiceDesc is the grpc.ServiceDesc for Node service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Node_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "Node",
	HandlerType: (*NodeServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ResisterNode",
			Handler:    _Node_ResisterNode_Handler,
		},
		{
			MethodName: "Sync",
			Handler:    _Node_Sync_Handler,
		},
		{
			MethodName: "Bloadcast",
			Handler:    _Node_Bloadcast_Handler,
		},
		{
			MethodName: "BloadcastBlock",
			Handler:    _Node_BloadcastBlock_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Upload",
			Handler:       _Node_Upload_Handler,
			ClientStreams: true,
		},
	},
	Metadata: "proto/node.proto",
}
