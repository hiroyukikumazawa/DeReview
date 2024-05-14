// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             (unknown)
// source: dereview/dereview/query.proto

package dereview

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
	Query_Params_FullMethodName        = "/dereview.dereview.Query/Params"
	Query_HelloDereview_FullMethodName = "/dereview.dereview.Query/HelloDereview"
	Query_GetPost_FullMethodName       = "/dereview.dereview.Query/GetPost"
	Query_ListPost_FullMethodName      = "/dereview.dereview.Query/ListPost"
)

// QueryClient is the client API for Query service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type QueryClient interface {
	// Parameters queries the parameters of the module.
	Params(ctx context.Context, in *QueryParamsRequest, opts ...grpc.CallOption) (*QueryParamsResponse, error)
	// Queries a list of HelloDereview items.
	HelloDereview(ctx context.Context, in *QueryHelloDereviewRequest, opts ...grpc.CallOption) (*QueryHelloDereviewResponse, error)
	// Queries a list of Post items.
	GetPost(ctx context.Context, in *QueryGetPostRequest, opts ...grpc.CallOption) (*QueryGetPostResponse, error)
	ListPost(ctx context.Context, in *QueryAllPostRequest, opts ...grpc.CallOption) (*QueryAllPostResponse, error)
}

type queryClient struct {
	cc grpc.ClientConnInterface
}

func NewQueryClient(cc grpc.ClientConnInterface) QueryClient {
	return &queryClient{cc}
}

func (c *queryClient) Params(ctx context.Context, in *QueryParamsRequest, opts ...grpc.CallOption) (*QueryParamsResponse, error) {
	out := new(QueryParamsResponse)
	err := c.cc.Invoke(ctx, Query_Params_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) HelloDereview(ctx context.Context, in *QueryHelloDereviewRequest, opts ...grpc.CallOption) (*QueryHelloDereviewResponse, error) {
	out := new(QueryHelloDereviewResponse)
	err := c.cc.Invoke(ctx, Query_HelloDereview_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) GetPost(ctx context.Context, in *QueryGetPostRequest, opts ...grpc.CallOption) (*QueryGetPostResponse, error) {
	out := new(QueryGetPostResponse)
	err := c.cc.Invoke(ctx, Query_GetPost_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) ListPost(ctx context.Context, in *QueryAllPostRequest, opts ...grpc.CallOption) (*QueryAllPostResponse, error) {
	out := new(QueryAllPostResponse)
	err := c.cc.Invoke(ctx, Query_ListPost_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// QueryServer is the server API for Query service.
// All implementations must embed UnimplementedQueryServer
// for forward compatibility
type QueryServer interface {
	// Parameters queries the parameters of the module.
	Params(context.Context, *QueryParamsRequest) (*QueryParamsResponse, error)
	// Queries a list of HelloDereview items.
	HelloDereview(context.Context, *QueryHelloDereviewRequest) (*QueryHelloDereviewResponse, error)
	// Queries a list of Post items.
	GetPost(context.Context, *QueryGetPostRequest) (*QueryGetPostResponse, error)
	ListPost(context.Context, *QueryAllPostRequest) (*QueryAllPostResponse, error)
	mustEmbedUnimplementedQueryServer()
}

// UnimplementedQueryServer must be embedded to have forward compatible implementations.
type UnimplementedQueryServer struct {
}

func (UnimplementedQueryServer) Params(context.Context, *QueryParamsRequest) (*QueryParamsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Params not implemented")
}
func (UnimplementedQueryServer) HelloDereview(context.Context, *QueryHelloDereviewRequest) (*QueryHelloDereviewResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method HelloDereview not implemented")
}
func (UnimplementedQueryServer) GetPost(context.Context, *QueryGetPostRequest) (*QueryGetPostResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPost not implemented")
}
func (UnimplementedQueryServer) ListPost(context.Context, *QueryAllPostRequest) (*QueryAllPostResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListPost not implemented")
}
func (UnimplementedQueryServer) mustEmbedUnimplementedQueryServer() {}

// UnsafeQueryServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to QueryServer will
// result in compilation errors.
type UnsafeQueryServer interface {
	mustEmbedUnimplementedQueryServer()
}

func RegisterQueryServer(s grpc.ServiceRegistrar, srv QueryServer) {
	s.RegisterService(&Query_ServiceDesc, srv)
}

func _Query_Params_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryParamsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).Params(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Query_Params_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).Params(ctx, req.(*QueryParamsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_HelloDereview_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryHelloDereviewRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).HelloDereview(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Query_HelloDereview_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).HelloDereview(ctx, req.(*QueryHelloDereviewRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_GetPost_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryGetPostRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).GetPost(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Query_GetPost_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).GetPost(ctx, req.(*QueryGetPostRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_ListPost_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryAllPostRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).ListPost(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Query_ListPost_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).ListPost(ctx, req.(*QueryAllPostRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Query_ServiceDesc is the grpc.ServiceDesc for Query service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Query_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "dereview.dereview.Query",
	HandlerType: (*QueryServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Params",
			Handler:    _Query_Params_Handler,
		},
		{
			MethodName: "HelloDereview",
			Handler:    _Query_HelloDereview_Handler,
		},
		{
			MethodName: "GetPost",
			Handler:    _Query_GetPost_Handler,
		},
		{
			MethodName: "ListPost",
			Handler:    _Query_ListPost_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "dereview/dereview/query.proto",
}
