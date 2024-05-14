// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             (unknown)
// source: dereview/dereview/tx.proto

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
	Msg_UpdateParams_FullMethodName = "/dereview.dereview.Msg/UpdateParams"
	Msg_CreatePost_FullMethodName   = "/dereview.dereview.Msg/CreatePost"
	Msg_UpdatePost_FullMethodName   = "/dereview.dereview.Msg/UpdatePost"
	Msg_DeletePost_FullMethodName   = "/dereview.dereview.Msg/DeletePost"
	Msg_CreatePost1_FullMethodName  = "/dereview.dereview.Msg/CreatePost1"
	Msg_UpdatePost1_FullMethodName  = "/dereview.dereview.Msg/UpdatePost1"
	Msg_DeletePost1_FullMethodName  = "/dereview.dereview.Msg/DeletePost1"
)

// MsgClient is the client API for Msg service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MsgClient interface {
	// UpdateParams defines a (governance) operation for updating the module
	// parameters. The authority defaults to the x/gov module account.
	UpdateParams(ctx context.Context, in *MsgUpdateParams, opts ...grpc.CallOption) (*MsgUpdateParamsResponse, error)
	CreatePost(ctx context.Context, in *MsgCreatePost, opts ...grpc.CallOption) (*MsgCreatePostResponse, error)
	UpdatePost(ctx context.Context, in *MsgUpdatePost, opts ...grpc.CallOption) (*MsgUpdatePostResponse, error)
	DeletePost(ctx context.Context, in *MsgDeletePost, opts ...grpc.CallOption) (*MsgDeletePostResponse, error)
	CreatePost1(ctx context.Context, in *MsgCreatePost1, opts ...grpc.CallOption) (*MsgCreatePost1Response, error)
	UpdatePost1(ctx context.Context, in *MsgUpdatePost1, opts ...grpc.CallOption) (*MsgUpdatePost1Response, error)
	DeletePost1(ctx context.Context, in *MsgDeletePost1, opts ...grpc.CallOption) (*MsgDeletePost1Response, error)
}

type msgClient struct {
	cc grpc.ClientConnInterface
}

func NewMsgClient(cc grpc.ClientConnInterface) MsgClient {
	return &msgClient{cc}
}

func (c *msgClient) UpdateParams(ctx context.Context, in *MsgUpdateParams, opts ...grpc.CallOption) (*MsgUpdateParamsResponse, error) {
	out := new(MsgUpdateParamsResponse)
	err := c.cc.Invoke(ctx, Msg_UpdateParams_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) CreatePost(ctx context.Context, in *MsgCreatePost, opts ...grpc.CallOption) (*MsgCreatePostResponse, error) {
	out := new(MsgCreatePostResponse)
	err := c.cc.Invoke(ctx, Msg_CreatePost_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) UpdatePost(ctx context.Context, in *MsgUpdatePost, opts ...grpc.CallOption) (*MsgUpdatePostResponse, error) {
	out := new(MsgUpdatePostResponse)
	err := c.cc.Invoke(ctx, Msg_UpdatePost_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) DeletePost(ctx context.Context, in *MsgDeletePost, opts ...grpc.CallOption) (*MsgDeletePostResponse, error) {
	out := new(MsgDeletePostResponse)
	err := c.cc.Invoke(ctx, Msg_DeletePost_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) CreatePost1(ctx context.Context, in *MsgCreatePost1, opts ...grpc.CallOption) (*MsgCreatePost1Response, error) {
	out := new(MsgCreatePost1Response)
	err := c.cc.Invoke(ctx, Msg_CreatePost1_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) UpdatePost1(ctx context.Context, in *MsgUpdatePost1, opts ...grpc.CallOption) (*MsgUpdatePost1Response, error) {
	out := new(MsgUpdatePost1Response)
	err := c.cc.Invoke(ctx, Msg_UpdatePost1_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) DeletePost1(ctx context.Context, in *MsgDeletePost1, opts ...grpc.CallOption) (*MsgDeletePost1Response, error) {
	out := new(MsgDeletePost1Response)
	err := c.cc.Invoke(ctx, Msg_DeletePost1_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MsgServer is the server API for Msg service.
// All implementations must embed UnimplementedMsgServer
// for forward compatibility
type MsgServer interface {
	// UpdateParams defines a (governance) operation for updating the module
	// parameters. The authority defaults to the x/gov module account.
	UpdateParams(context.Context, *MsgUpdateParams) (*MsgUpdateParamsResponse, error)
	CreatePost(context.Context, *MsgCreatePost) (*MsgCreatePostResponse, error)
	UpdatePost(context.Context, *MsgUpdatePost) (*MsgUpdatePostResponse, error)
	DeletePost(context.Context, *MsgDeletePost) (*MsgDeletePostResponse, error)
	CreatePost1(context.Context, *MsgCreatePost1) (*MsgCreatePost1Response, error)
	UpdatePost1(context.Context, *MsgUpdatePost1) (*MsgUpdatePost1Response, error)
	DeletePost1(context.Context, *MsgDeletePost1) (*MsgDeletePost1Response, error)
	mustEmbedUnimplementedMsgServer()
}

// UnimplementedMsgServer must be embedded to have forward compatible implementations.
type UnimplementedMsgServer struct {
}

func (UnimplementedMsgServer) UpdateParams(context.Context, *MsgUpdateParams) (*MsgUpdateParamsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateParams not implemented")
}
func (UnimplementedMsgServer) CreatePost(context.Context, *MsgCreatePost) (*MsgCreatePostResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreatePost not implemented")
}
func (UnimplementedMsgServer) UpdatePost(context.Context, *MsgUpdatePost) (*MsgUpdatePostResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdatePost not implemented")
}
func (UnimplementedMsgServer) DeletePost(context.Context, *MsgDeletePost) (*MsgDeletePostResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeletePost not implemented")
}
func (UnimplementedMsgServer) CreatePost1(context.Context, *MsgCreatePost1) (*MsgCreatePost1Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreatePost1 not implemented")
}
func (UnimplementedMsgServer) UpdatePost1(context.Context, *MsgUpdatePost1) (*MsgUpdatePost1Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdatePost1 not implemented")
}
func (UnimplementedMsgServer) DeletePost1(context.Context, *MsgDeletePost1) (*MsgDeletePost1Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeletePost1 not implemented")
}
func (UnimplementedMsgServer) mustEmbedUnimplementedMsgServer() {}

// UnsafeMsgServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to MsgServer will
// result in compilation errors.
type UnsafeMsgServer interface {
	mustEmbedUnimplementedMsgServer()
}

func RegisterMsgServer(s grpc.ServiceRegistrar, srv MsgServer) {
	s.RegisterService(&Msg_ServiceDesc, srv)
}

func _Msg_UpdateParams_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgUpdateParams)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).UpdateParams(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Msg_UpdateParams_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).UpdateParams(ctx, req.(*MsgUpdateParams))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_CreatePost_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgCreatePost)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).CreatePost(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Msg_CreatePost_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).CreatePost(ctx, req.(*MsgCreatePost))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_UpdatePost_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgUpdatePost)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).UpdatePost(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Msg_UpdatePost_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).UpdatePost(ctx, req.(*MsgUpdatePost))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_DeletePost_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgDeletePost)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).DeletePost(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Msg_DeletePost_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).DeletePost(ctx, req.(*MsgDeletePost))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_CreatePost1_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgCreatePost1)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).CreatePost1(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Msg_CreatePost1_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).CreatePost1(ctx, req.(*MsgCreatePost1))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_UpdatePost1_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgUpdatePost1)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).UpdatePost1(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Msg_UpdatePost1_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).UpdatePost1(ctx, req.(*MsgUpdatePost1))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_DeletePost1_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgDeletePost1)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).DeletePost1(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Msg_DeletePost1_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).DeletePost1(ctx, req.(*MsgDeletePost1))
	}
	return interceptor(ctx, in, info, handler)
}

// Msg_ServiceDesc is the grpc.ServiceDesc for Msg service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Msg_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "dereview.dereview.Msg",
	HandlerType: (*MsgServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "UpdateParams",
			Handler:    _Msg_UpdateParams_Handler,
		},
		{
			MethodName: "CreatePost",
			Handler:    _Msg_CreatePost_Handler,
		},
		{
			MethodName: "UpdatePost",
			Handler:    _Msg_UpdatePost_Handler,
		},
		{
			MethodName: "DeletePost",
			Handler:    _Msg_DeletePost_Handler,
		},
		{
			MethodName: "CreatePost1",
			Handler:    _Msg_CreatePost1_Handler,
		},
		{
			MethodName: "UpdatePost1",
			Handler:    _Msg_UpdatePost1_Handler,
		},
		{
			MethodName: "DeletePost1",
			Handler:    _Msg_DeletePost1_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "dereview/dereview/tx.proto",
}
