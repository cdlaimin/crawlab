// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.4.0
// - protoc             v5.27.2
// source: services/node_service.proto

package grpc

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.62.0 or later.
const _ = grpc.SupportPackageIsVersion8

const (
	NodeService_Register_FullMethodName      = "/grpc.NodeService/Register"
	NodeService_SendHeartbeat_FullMethodName = "/grpc.NodeService/SendHeartbeat"
	NodeService_Subscribe_FullMethodName     = "/grpc.NodeService/Subscribe"
	NodeService_Unsubscribe_FullMethodName   = "/grpc.NodeService/Unsubscribe"
)

// NodeServiceClient is the client API for NodeService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type NodeServiceClient interface {
	Register(ctx context.Context, in *NodeServiceRegisterRequest, opts ...grpc.CallOption) (*Response, error)
	SendHeartbeat(ctx context.Context, in *NodeServiceSendHeartbeatRequest, opts ...grpc.CallOption) (*Response, error)
	Subscribe(ctx context.Context, in *Request, opts ...grpc.CallOption) (NodeService_SubscribeClient, error)
	Unsubscribe(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error)
}

type nodeServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewNodeServiceClient(cc grpc.ClientConnInterface) NodeServiceClient {
	return &nodeServiceClient{cc}
}

func (c *nodeServiceClient) Register(ctx context.Context, in *NodeServiceRegisterRequest, opts ...grpc.CallOption) (*Response, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Response)
	err := c.cc.Invoke(ctx, NodeService_Register_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *nodeServiceClient) SendHeartbeat(ctx context.Context, in *NodeServiceSendHeartbeatRequest, opts ...grpc.CallOption) (*Response, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Response)
	err := c.cc.Invoke(ctx, NodeService_SendHeartbeat_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *nodeServiceClient) Subscribe(ctx context.Context, in *Request, opts ...grpc.CallOption) (NodeService_SubscribeClient, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	stream, err := c.cc.NewStream(ctx, &NodeService_ServiceDesc.Streams[0], NodeService_Subscribe_FullMethodName, cOpts...)
	if err != nil {
		return nil, err
	}
	x := &nodeServiceSubscribeClient{ClientStream: stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type NodeService_SubscribeClient interface {
	Recv() (*StreamMessage, error)
	grpc.ClientStream
}

type nodeServiceSubscribeClient struct {
	grpc.ClientStream
}

func (x *nodeServiceSubscribeClient) Recv() (*StreamMessage, error) {
	m := new(StreamMessage)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *nodeServiceClient) Unsubscribe(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Response)
	err := c.cc.Invoke(ctx, NodeService_Unsubscribe_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// NodeServiceServer is the server API for NodeService service.
// All implementations must embed UnimplementedNodeServiceServer
// for forward compatibility
type NodeServiceServer interface {
	Register(context.Context, *NodeServiceRegisterRequest) (*Response, error)
	SendHeartbeat(context.Context, *NodeServiceSendHeartbeatRequest) (*Response, error)
	Subscribe(*Request, NodeService_SubscribeServer) error
	Unsubscribe(context.Context, *Request) (*Response, error)
	mustEmbedUnimplementedNodeServiceServer()
}

// UnimplementedNodeServiceServer must be embedded to have forward compatible implementations.
type UnimplementedNodeServiceServer struct {
}

func (UnimplementedNodeServiceServer) Register(context.Context, *NodeServiceRegisterRequest) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Register not implemented")
}
func (UnimplementedNodeServiceServer) SendHeartbeat(context.Context, *NodeServiceSendHeartbeatRequest) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendHeartbeat not implemented")
}
func (UnimplementedNodeServiceServer) Subscribe(*Request, NodeService_SubscribeServer) error {
	return status.Errorf(codes.Unimplemented, "method Subscribe not implemented")
}
func (UnimplementedNodeServiceServer) Unsubscribe(context.Context, *Request) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Unsubscribe not implemented")
}
func (UnimplementedNodeServiceServer) mustEmbedUnimplementedNodeServiceServer() {}

// UnsafeNodeServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to NodeServiceServer will
// result in compilation errors.
type UnsafeNodeServiceServer interface {
	mustEmbedUnimplementedNodeServiceServer()
}

func RegisterNodeServiceServer(s grpc.ServiceRegistrar, srv NodeServiceServer) {
	s.RegisterService(&NodeService_ServiceDesc, srv)
}

func _NodeService_Register_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NodeServiceRegisterRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NodeServiceServer).Register(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: NodeService_Register_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NodeServiceServer).Register(ctx, req.(*NodeServiceRegisterRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _NodeService_SendHeartbeat_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NodeServiceSendHeartbeatRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NodeServiceServer).SendHeartbeat(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: NodeService_SendHeartbeat_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NodeServiceServer).SendHeartbeat(ctx, req.(*NodeServiceSendHeartbeatRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _NodeService_Subscribe_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(Request)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(NodeServiceServer).Subscribe(m, &nodeServiceSubscribeServer{ServerStream: stream})
}

type NodeService_SubscribeServer interface {
	Send(*StreamMessage) error
	grpc.ServerStream
}

type nodeServiceSubscribeServer struct {
	grpc.ServerStream
}

func (x *nodeServiceSubscribeServer) Send(m *StreamMessage) error {
	return x.ServerStream.SendMsg(m)
}

func _NodeService_Unsubscribe_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NodeServiceServer).Unsubscribe(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: NodeService_Unsubscribe_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NodeServiceServer).Unsubscribe(ctx, req.(*Request))
	}
	return interceptor(ctx, in, info, handler)
}

// NodeService_ServiceDesc is the grpc.ServiceDesc for NodeService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var NodeService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "grpc.NodeService",
	HandlerType: (*NodeServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Register",
			Handler:    _NodeService_Register_Handler,
		},
		{
			MethodName: "SendHeartbeat",
			Handler:    _NodeService_SendHeartbeat_Handler,
		},
		{
			MethodName: "Unsubscribe",
			Handler:    _NodeService_Unsubscribe_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Subscribe",
			Handler:       _NodeService_Subscribe_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "services/node_service.proto",
}
