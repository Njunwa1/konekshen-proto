// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.25.3
// source: matches/matches.proto

package matches

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
	Matcher_CreateMatch_FullMethodName = "/Matcher/CreateMatch"
	Matcher_GetMatches_FullMethodName  = "/Matcher/GetMatches"
)

// MatcherClient is the client API for Matcher service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MatcherClient interface {
	CreateMatch(ctx context.Context, in *CreateMatchRequest, opts ...grpc.CallOption) (*CreateMatchResponse, error)
	GetMatches(ctx context.Context, in *GetMatchesRequest, opts ...grpc.CallOption) (*GetMatchesResponse, error)
}

type matcherClient struct {
	cc grpc.ClientConnInterface
}

func NewMatcherClient(cc grpc.ClientConnInterface) MatcherClient {
	return &matcherClient{cc}
}

func (c *matcherClient) CreateMatch(ctx context.Context, in *CreateMatchRequest, opts ...grpc.CallOption) (*CreateMatchResponse, error) {
	out := new(CreateMatchResponse)
	err := c.cc.Invoke(ctx, Matcher_CreateMatch_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *matcherClient) GetMatches(ctx context.Context, in *GetMatchesRequest, opts ...grpc.CallOption) (*GetMatchesResponse, error) {
	out := new(GetMatchesResponse)
	err := c.cc.Invoke(ctx, Matcher_GetMatches_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MatcherServer is the server API for Matcher service.
// All implementations must embed UnimplementedMatcherServer
// for forward compatibility
type MatcherServer interface {
	CreateMatch(context.Context, *CreateMatchRequest) (*CreateMatchResponse, error)
	GetMatches(context.Context, *GetMatchesRequest) (*GetMatchesResponse, error)
	mustEmbedUnimplementedMatcherServer()
}

// UnimplementedMatcherServer must be embedded to have forward compatible implementations.
type UnimplementedMatcherServer struct {
}

func (UnimplementedMatcherServer) CreateMatch(context.Context, *CreateMatchRequest) (*CreateMatchResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateMatch not implemented")
}
func (UnimplementedMatcherServer) GetMatches(context.Context, *GetMatchesRequest) (*GetMatchesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetMatches not implemented")
}
func (UnimplementedMatcherServer) mustEmbedUnimplementedMatcherServer() {}

// UnsafeMatcherServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to MatcherServer will
// result in compilation errors.
type UnsafeMatcherServer interface {
	mustEmbedUnimplementedMatcherServer()
}

func RegisterMatcherServer(s grpc.ServiceRegistrar, srv MatcherServer) {
	s.RegisterService(&Matcher_ServiceDesc, srv)
}

func _Matcher_CreateMatch_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateMatchRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MatcherServer).CreateMatch(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Matcher_CreateMatch_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MatcherServer).CreateMatch(ctx, req.(*CreateMatchRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Matcher_GetMatches_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetMatchesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MatcherServer).GetMatches(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Matcher_GetMatches_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MatcherServer).GetMatches(ctx, req.(*GetMatchesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Matcher_ServiceDesc is the grpc.ServiceDesc for Matcher service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Matcher_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "Matcher",
	HandlerType: (*MatcherServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateMatch",
			Handler:    _Matcher_CreateMatch_Handler,
		},
		{
			MethodName: "GetMatches",
			Handler:    _Matcher_GetMatches_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "matches/matches.proto",
}
