// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.25.3
// source: swipes/swipes.proto

package swipes

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
	Swipes_CreateLikes_FullMethodName    = "/Swipes/CreateLikes"
	Swipes_CreateDisLikes_FullMethodName = "/Swipes/CreateDisLikes"
	Swipes_GetLikes_FullMethodName       = "/Swipes/GetLikes"
)

// SwipesClient is the client API for Swipes service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type SwipesClient interface {
	CreateLikes(ctx context.Context, in *LikesRequest, opts ...grpc.CallOption) (*LikesResponse, error)
	CreateDisLikes(ctx context.Context, in *DisLikesRequest, opts ...grpc.CallOption) (*DisLikesResponse, error)
	GetLikes(ctx context.Context, in *GetLikesRequest, opts ...grpc.CallOption) (*GetLikesResponse, error)
}

type swipesClient struct {
	cc grpc.ClientConnInterface
}

func NewSwipesClient(cc grpc.ClientConnInterface) SwipesClient {
	return &swipesClient{cc}
}

func (c *swipesClient) CreateLikes(ctx context.Context, in *LikesRequest, opts ...grpc.CallOption) (*LikesResponse, error) {
	out := new(LikesResponse)
	err := c.cc.Invoke(ctx, Swipes_CreateLikes_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *swipesClient) CreateDisLikes(ctx context.Context, in *DisLikesRequest, opts ...grpc.CallOption) (*DisLikesResponse, error) {
	out := new(DisLikesResponse)
	err := c.cc.Invoke(ctx, Swipes_CreateDisLikes_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *swipesClient) GetLikes(ctx context.Context, in *GetLikesRequest, opts ...grpc.CallOption) (*GetLikesResponse, error) {
	out := new(GetLikesResponse)
	err := c.cc.Invoke(ctx, Swipes_GetLikes_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SwipesServer is the server API for Swipes service.
// All implementations must embed UnimplementedSwipesServer
// for forward compatibility
type SwipesServer interface {
	CreateLikes(context.Context, *LikesRequest) (*LikesResponse, error)
	CreateDisLikes(context.Context, *DisLikesRequest) (*DisLikesResponse, error)
	GetLikes(context.Context, *GetLikesRequest) (*GetLikesResponse, error)
	mustEmbedUnimplementedSwipesServer()
}

// UnimplementedSwipesServer must be embedded to have forward compatible implementations.
type UnimplementedSwipesServer struct {
}

func (UnimplementedSwipesServer) CreateLikes(context.Context, *LikesRequest) (*LikesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateLikes not implemented")
}
func (UnimplementedSwipesServer) CreateDisLikes(context.Context, *DisLikesRequest) (*DisLikesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateDisLikes not implemented")
}
func (UnimplementedSwipesServer) GetLikes(context.Context, *GetLikesRequest) (*GetLikesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetLikes not implemented")
}
func (UnimplementedSwipesServer) mustEmbedUnimplementedSwipesServer() {}

// UnsafeSwipesServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to SwipesServer will
// result in compilation errors.
type UnsafeSwipesServer interface {
	mustEmbedUnimplementedSwipesServer()
}

func RegisterSwipesServer(s grpc.ServiceRegistrar, srv SwipesServer) {
	s.RegisterService(&Swipes_ServiceDesc, srv)
}

func _Swipes_CreateLikes_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LikesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SwipesServer).CreateLikes(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Swipes_CreateLikes_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SwipesServer).CreateLikes(ctx, req.(*LikesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Swipes_CreateDisLikes_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DisLikesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SwipesServer).CreateDisLikes(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Swipes_CreateDisLikes_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SwipesServer).CreateDisLikes(ctx, req.(*DisLikesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Swipes_GetLikes_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetLikesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SwipesServer).GetLikes(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Swipes_GetLikes_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SwipesServer).GetLikes(ctx, req.(*GetLikesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Swipes_ServiceDesc is the grpc.ServiceDesc for Swipes service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Swipes_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "Swipes",
	HandlerType: (*SwipesServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateLikes",
			Handler:    _Swipes_CreateLikes_Handler,
		},
		{
			MethodName: "CreateDisLikes",
			Handler:    _Swipes_CreateDisLikes_Handler,
		},
		{
			MethodName: "GetLikes",
			Handler:    _Swipes_GetLikes_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "swipes/swipes.proto",
}