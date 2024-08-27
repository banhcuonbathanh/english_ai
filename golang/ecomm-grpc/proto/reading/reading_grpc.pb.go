// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v4.25.1
// source: ecomm-grpc/proto/reading/reading.proto

package reading

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	EcommReading_CreateReading_FullMethodName     = "/proto.EcommReading/CreateReading"
	EcommReading_SaveReading_FullMethodName       = "/proto.EcommReading/SaveReading"
	EcommReading_UpdateReading_FullMethodName     = "/proto.EcommReading/UpdateReading"
	EcommReading_DeleteReading_FullMethodName     = "/proto.EcommReading/DeleteReading"
	EcommReading_FindAllReading_FullMethodName    = "/proto.EcommReading/FindAllReading"
	EcommReading_FindByID_FullMethodName          = "/proto.EcommReading/FindByID"
	EcommReading_FindReadingByPage_FullMethodName = "/proto.EcommReading/FindReadingByPage"
)

// EcommReadingClient is the client API for EcommReading service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type EcommReadingClient interface {
	CreateReading(ctx context.Context, in *ReadingReq, opts ...grpc.CallOption) (*ReadingRes, error)
	SaveReading(ctx context.Context, in *ReadingReq, opts ...grpc.CallOption) (*emptypb.Empty, error)
	UpdateReading(ctx context.Context, in *ReadingReq, opts ...grpc.CallOption) (*ReadingRes, error)
	DeleteReading(ctx context.Context, in *ReadingReq, opts ...grpc.CallOption) (*emptypb.Empty, error)
	FindAllReading(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*ReadingResList, error)
	FindByID(ctx context.Context, in *ReadingReq, opts ...grpc.CallOption) (*ReadingRes, error)
	FindReadingByPage(ctx context.Context, in *PageRequest, opts ...grpc.CallOption) (*ReadingResList, error)
}

type ecommReadingClient struct {
	cc grpc.ClientConnInterface
}

func NewEcommReadingClient(cc grpc.ClientConnInterface) EcommReadingClient {
	return &ecommReadingClient{cc}
}

func (c *ecommReadingClient) CreateReading(ctx context.Context, in *ReadingReq, opts ...grpc.CallOption) (*ReadingRes, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ReadingRes)
	err := c.cc.Invoke(ctx, EcommReading_CreateReading_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ecommReadingClient) SaveReading(ctx context.Context, in *ReadingReq, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, EcommReading_SaveReading_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ecommReadingClient) UpdateReading(ctx context.Context, in *ReadingReq, opts ...grpc.CallOption) (*ReadingRes, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ReadingRes)
	err := c.cc.Invoke(ctx, EcommReading_UpdateReading_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ecommReadingClient) DeleteReading(ctx context.Context, in *ReadingReq, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, EcommReading_DeleteReading_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ecommReadingClient) FindAllReading(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*ReadingResList, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ReadingResList)
	err := c.cc.Invoke(ctx, EcommReading_FindAllReading_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ecommReadingClient) FindByID(ctx context.Context, in *ReadingReq, opts ...grpc.CallOption) (*ReadingRes, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ReadingRes)
	err := c.cc.Invoke(ctx, EcommReading_FindByID_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ecommReadingClient) FindReadingByPage(ctx context.Context, in *PageRequest, opts ...grpc.CallOption) (*ReadingResList, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ReadingResList)
	err := c.cc.Invoke(ctx, EcommReading_FindReadingByPage_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// EcommReadingServer is the server API for EcommReading service.
// All implementations must embed UnimplementedEcommReadingServer
// for forward compatibility.
type EcommReadingServer interface {
	CreateReading(context.Context, *ReadingReq) (*ReadingRes, error)
	SaveReading(context.Context, *ReadingReq) (*emptypb.Empty, error)
	UpdateReading(context.Context, *ReadingReq) (*ReadingRes, error)
	DeleteReading(context.Context, *ReadingReq) (*emptypb.Empty, error)
	FindAllReading(context.Context, *emptypb.Empty) (*ReadingResList, error)
	FindByID(context.Context, *ReadingReq) (*ReadingRes, error)
	FindReadingByPage(context.Context, *PageRequest) (*ReadingResList, error)
	mustEmbedUnimplementedEcommReadingServer()
}

// UnimplementedEcommReadingServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedEcommReadingServer struct{}

func (UnimplementedEcommReadingServer) CreateReading(context.Context, *ReadingReq) (*ReadingRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateReading not implemented")
}
func (UnimplementedEcommReadingServer) SaveReading(context.Context, *ReadingReq) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SaveReading not implemented")
}
func (UnimplementedEcommReadingServer) UpdateReading(context.Context, *ReadingReq) (*ReadingRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateReading not implemented")
}
func (UnimplementedEcommReadingServer) DeleteReading(context.Context, *ReadingReq) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteReading not implemented")
}
func (UnimplementedEcommReadingServer) FindAllReading(context.Context, *emptypb.Empty) (*ReadingResList, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindAllReading not implemented")
}
func (UnimplementedEcommReadingServer) FindByID(context.Context, *ReadingReq) (*ReadingRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindByID not implemented")
}
func (UnimplementedEcommReadingServer) FindReadingByPage(context.Context, *PageRequest) (*ReadingResList, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindReadingByPage not implemented")
}
func (UnimplementedEcommReadingServer) mustEmbedUnimplementedEcommReadingServer() {}
func (UnimplementedEcommReadingServer) testEmbeddedByValue()                      {}

// UnsafeEcommReadingServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to EcommReadingServer will
// result in compilation errors.
type UnsafeEcommReadingServer interface {
	mustEmbedUnimplementedEcommReadingServer()
}

func RegisterEcommReadingServer(s grpc.ServiceRegistrar, srv EcommReadingServer) {
	// If the following call pancis, it indicates UnimplementedEcommReadingServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&EcommReading_ServiceDesc, srv)
}

func _EcommReading_CreateReading_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReadingReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EcommReadingServer).CreateReading(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: EcommReading_CreateReading_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EcommReadingServer).CreateReading(ctx, req.(*ReadingReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _EcommReading_SaveReading_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReadingReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EcommReadingServer).SaveReading(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: EcommReading_SaveReading_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EcommReadingServer).SaveReading(ctx, req.(*ReadingReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _EcommReading_UpdateReading_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReadingReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EcommReadingServer).UpdateReading(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: EcommReading_UpdateReading_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EcommReadingServer).UpdateReading(ctx, req.(*ReadingReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _EcommReading_DeleteReading_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReadingReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EcommReadingServer).DeleteReading(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: EcommReading_DeleteReading_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EcommReadingServer).DeleteReading(ctx, req.(*ReadingReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _EcommReading_FindAllReading_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EcommReadingServer).FindAllReading(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: EcommReading_FindAllReading_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EcommReadingServer).FindAllReading(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _EcommReading_FindByID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReadingReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EcommReadingServer).FindByID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: EcommReading_FindByID_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EcommReadingServer).FindByID(ctx, req.(*ReadingReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _EcommReading_FindReadingByPage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PageRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EcommReadingServer).FindReadingByPage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: EcommReading_FindReadingByPage_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EcommReadingServer).FindReadingByPage(ctx, req.(*PageRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// EcommReading_ServiceDesc is the grpc.ServiceDesc for EcommReading service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var EcommReading_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "proto.EcommReading",
	HandlerType: (*EcommReadingServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateReading",
			Handler:    _EcommReading_CreateReading_Handler,
		},
		{
			MethodName: "SaveReading",
			Handler:    _EcommReading_SaveReading_Handler,
		},
		{
			MethodName: "UpdateReading",
			Handler:    _EcommReading_UpdateReading_Handler,
		},
		{
			MethodName: "DeleteReading",
			Handler:    _EcommReading_DeleteReading_Handler,
		},
		{
			MethodName: "FindAllReading",
			Handler:    _EcommReading_FindAllReading_Handler,
		},
		{
			MethodName: "FindByID",
			Handler:    _EcommReading_FindByID_Handler,
		},
		{
			MethodName: "FindReadingByPage",
			Handler:    _EcommReading_FindReadingByPage_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "ecomm-grpc/proto/reading/reading.proto",
}
