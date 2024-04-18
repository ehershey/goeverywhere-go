// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v5.26.1
// source: stats.proto

package main

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

const (
	GOE_GetStats_FullMethodName = "/GOE/GetStats"
)

// GOEClient is the client API for GOE service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type GOEClient interface {
	GetStats(ctx context.Context, in *StatsRequest, opts ...grpc.CallOption) (*StatsResponse, error)
}

type gOEClient struct {
	cc grpc.ClientConnInterface
}

func NewGOEClient(cc grpc.ClientConnInterface) GOEClient {
	return &gOEClient{cc}
}

func (c *gOEClient) GetStats(ctx context.Context, in *StatsRequest, opts ...grpc.CallOption) (*StatsResponse, error) {
	out := new(StatsResponse)
	err := c.cc.Invoke(ctx, GOE_GetStats_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// GOEServer is the server API for GOE service.
// All implementations must embed UnimplementedGOEServer
// for forward compatibility
type GOEServer interface {
	GetStats(context.Context, *StatsRequest) (*StatsResponse, error)
	mustEmbedUnimplementedGOEServer()
}

// UnimplementedGOEServer must be embedded to have forward compatible implementations.
type UnimplementedGOEServer struct {
}

func (UnimplementedGOEServer) GetStats(context.Context, *StatsRequest) (*StatsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetStats not implemented")
}
func (UnimplementedGOEServer) mustEmbedUnimplementedGOEServer() {}

// UnsafeGOEServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to GOEServer will
// result in compilation errors.
type UnsafeGOEServer interface {
	mustEmbedUnimplementedGOEServer()
}

func RegisterGOEServer(s grpc.ServiceRegistrar, srv GOEServer) {
	s.RegisterService(&GOE_ServiceDesc, srv)
}

func _GOE_GetStats_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StatsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GOEServer).GetStats(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: GOE_GetStats_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GOEServer).GetStats(ctx, req.(*StatsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// GOE_ServiceDesc is the grpc.ServiceDesc for GOE service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var GOE_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "GOE",
	HandlerType: (*GOEServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetStats",
			Handler:    _GOE_GetStats_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "stats.proto",
}
