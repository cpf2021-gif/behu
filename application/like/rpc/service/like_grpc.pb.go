// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.12.4
// source: like.proto

package service

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

// LikeClient is the client API for Like service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type LikeClient interface {
	Thumbup(ctx context.Context, in *ThumbupRequest, opts ...grpc.CallOption) (*ThumbupResponse, error)
	IsThumbup(ctx context.Context, in *IsThumbupRequest, opts ...grpc.CallOption) (*IsThumbupResponse, error)
}

type likeClient struct {
	cc grpc.ClientConnInterface
}

func NewLikeClient(cc grpc.ClientConnInterface) LikeClient {
	return &likeClient{cc}
}

func (c *likeClient) Thumbup(ctx context.Context, in *ThumbupRequest, opts ...grpc.CallOption) (*ThumbupResponse, error) {
	out := new(ThumbupResponse)
	err := c.cc.Invoke(ctx, "/service.Like/Thumbup", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *likeClient) IsThumbup(ctx context.Context, in *IsThumbupRequest, opts ...grpc.CallOption) (*IsThumbupResponse, error) {
	out := new(IsThumbupResponse)
	err := c.cc.Invoke(ctx, "/service.Like/IsThumbup", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// LikeServer is the server API for Like service.
// All implementations must embed UnimplementedLikeServer
// for forward compatibility
type LikeServer interface {
	Thumbup(context.Context, *ThumbupRequest) (*ThumbupResponse, error)
	IsThumbup(context.Context, *IsThumbupRequest) (*IsThumbupResponse, error)
	mustEmbedUnimplementedLikeServer()
}

// UnimplementedLikeServer must be embedded to have forward compatible implementations.
type UnimplementedLikeServer struct {
}

func (UnimplementedLikeServer) Thumbup(context.Context, *ThumbupRequest) (*ThumbupResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Thumbup not implemented")
}
func (UnimplementedLikeServer) IsThumbup(context.Context, *IsThumbupRequest) (*IsThumbupResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method IsThumbup not implemented")
}
func (UnimplementedLikeServer) mustEmbedUnimplementedLikeServer() {}

// UnsafeLikeServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to LikeServer will
// result in compilation errors.
type UnsafeLikeServer interface {
	mustEmbedUnimplementedLikeServer()
}

func RegisterLikeServer(s grpc.ServiceRegistrar, srv LikeServer) {
	s.RegisterService(&Like_ServiceDesc, srv)
}

func _Like_Thumbup_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ThumbupRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LikeServer).Thumbup(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/service.Like/Thumbup",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LikeServer).Thumbup(ctx, req.(*ThumbupRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Like_IsThumbup_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IsThumbupRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LikeServer).IsThumbup(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/service.Like/IsThumbup",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LikeServer).IsThumbup(ctx, req.(*IsThumbupRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Like_ServiceDesc is the grpc.ServiceDesc for Like service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Like_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "service.Like",
	HandlerType: (*LikeServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Thumbup",
			Handler:    _Like_Thumbup_Handler,
		},
		{
			MethodName: "IsThumbup",
			Handler:    _Like_IsThumbup_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "like.proto",
}
