// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v4.25.1
// source: quanqr/proto_qr/set/set.proto

package set

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
	SetService_GetSetProtoList_FullMethodName   = "/set_qr.SetService/GetSetProtoList"
	SetService_GetSetProtoDetail_FullMethodName = "/set_qr.SetService/GetSetProtoDetail"
	SetService_CreateSetProto_FullMethodName    = "/set_qr.SetService/CreateSetProto"
	SetService_UpdateSetProto_FullMethodName    = "/set_qr.SetService/UpdateSetProto"
	SetService_DeleteSetProto_FullMethodName    = "/set_qr.SetService/DeleteSetProto"
)

// SetServiceClient is the client API for SetService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type SetServiceClient interface {
	GetSetProtoList(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*SetProtoListResponse, error)
	GetSetProtoDetail(ctx context.Context, in *SetProtoIdParam, opts ...grpc.CallOption) (*SetProtoResponse, error)
	CreateSetProto(ctx context.Context, in *CreateSetProtoRequest, opts ...grpc.CallOption) (*SetProtoResponse, error)
	UpdateSetProto(ctx context.Context, in *UpdateSetProtoRequest, opts ...grpc.CallOption) (*SetProtoResponse, error)
	DeleteSetProto(ctx context.Context, in *SetProtoIdParam, opts ...grpc.CallOption) (*SetProtoResponse, error)
}

type setServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewSetServiceClient(cc grpc.ClientConnInterface) SetServiceClient {
	return &setServiceClient{cc}
}

func (c *setServiceClient) GetSetProtoList(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*SetProtoListResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(SetProtoListResponse)
	err := c.cc.Invoke(ctx, SetService_GetSetProtoList_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *setServiceClient) GetSetProtoDetail(ctx context.Context, in *SetProtoIdParam, opts ...grpc.CallOption) (*SetProtoResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(SetProtoResponse)
	err := c.cc.Invoke(ctx, SetService_GetSetProtoDetail_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *setServiceClient) CreateSetProto(ctx context.Context, in *CreateSetProtoRequest, opts ...grpc.CallOption) (*SetProtoResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(SetProtoResponse)
	err := c.cc.Invoke(ctx, SetService_CreateSetProto_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *setServiceClient) UpdateSetProto(ctx context.Context, in *UpdateSetProtoRequest, opts ...grpc.CallOption) (*SetProtoResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(SetProtoResponse)
	err := c.cc.Invoke(ctx, SetService_UpdateSetProto_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *setServiceClient) DeleteSetProto(ctx context.Context, in *SetProtoIdParam, opts ...grpc.CallOption) (*SetProtoResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(SetProtoResponse)
	err := c.cc.Invoke(ctx, SetService_DeleteSetProto_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SetServiceServer is the server API for SetService service.
// All implementations must embed UnimplementedSetServiceServer
// for forward compatibility.
type SetServiceServer interface {
	GetSetProtoList(context.Context, *emptypb.Empty) (*SetProtoListResponse, error)
	GetSetProtoDetail(context.Context, *SetProtoIdParam) (*SetProtoResponse, error)
	CreateSetProto(context.Context, *CreateSetProtoRequest) (*SetProtoResponse, error)
	UpdateSetProto(context.Context, *UpdateSetProtoRequest) (*SetProtoResponse, error)
	DeleteSetProto(context.Context, *SetProtoIdParam) (*SetProtoResponse, error)
	mustEmbedUnimplementedSetServiceServer()
}

// UnimplementedSetServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedSetServiceServer struct{}

func (UnimplementedSetServiceServer) GetSetProtoList(context.Context, *emptypb.Empty) (*SetProtoListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetSetProtoList not implemented")
}
func (UnimplementedSetServiceServer) GetSetProtoDetail(context.Context, *SetProtoIdParam) (*SetProtoResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetSetProtoDetail not implemented")
}
func (UnimplementedSetServiceServer) CreateSetProto(context.Context, *CreateSetProtoRequest) (*SetProtoResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateSetProto not implemented")
}
func (UnimplementedSetServiceServer) UpdateSetProto(context.Context, *UpdateSetProtoRequest) (*SetProtoResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateSetProto not implemented")
}
func (UnimplementedSetServiceServer) DeleteSetProto(context.Context, *SetProtoIdParam) (*SetProtoResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteSetProto not implemented")
}
func (UnimplementedSetServiceServer) mustEmbedUnimplementedSetServiceServer() {}
func (UnimplementedSetServiceServer) testEmbeddedByValue()                    {}

// UnsafeSetServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to SetServiceServer will
// result in compilation errors.
type UnsafeSetServiceServer interface {
	mustEmbedUnimplementedSetServiceServer()
}

func RegisterSetServiceServer(s grpc.ServiceRegistrar, srv SetServiceServer) {
	// If the following call pancis, it indicates UnimplementedSetServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&SetService_ServiceDesc, srv)
}

func _SetService_GetSetProtoList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SetServiceServer).GetSetProtoList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SetService_GetSetProtoList_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SetServiceServer).GetSetProtoList(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _SetService_GetSetProtoDetail_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SetProtoIdParam)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SetServiceServer).GetSetProtoDetail(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SetService_GetSetProtoDetail_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SetServiceServer).GetSetProtoDetail(ctx, req.(*SetProtoIdParam))
	}
	return interceptor(ctx, in, info, handler)
}

func _SetService_CreateSetProto_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateSetProtoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SetServiceServer).CreateSetProto(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SetService_CreateSetProto_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SetServiceServer).CreateSetProto(ctx, req.(*CreateSetProtoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SetService_UpdateSetProto_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateSetProtoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SetServiceServer).UpdateSetProto(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SetService_UpdateSetProto_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SetServiceServer).UpdateSetProto(ctx, req.(*UpdateSetProtoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SetService_DeleteSetProto_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SetProtoIdParam)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SetServiceServer).DeleteSetProto(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SetService_DeleteSetProto_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SetServiceServer).DeleteSetProto(ctx, req.(*SetProtoIdParam))
	}
	return interceptor(ctx, in, info, handler)
}

// SetService_ServiceDesc is the grpc.ServiceDesc for SetService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var SetService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "set_qr.SetService",
	HandlerType: (*SetServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetSetProtoList",
			Handler:    _SetService_GetSetProtoList_Handler,
		},
		{
			MethodName: "GetSetProtoDetail",
			Handler:    _SetService_GetSetProtoDetail_Handler,
		},
		{
			MethodName: "CreateSetProto",
			Handler:    _SetService_CreateSetProto_Handler,
		},
		{
			MethodName: "UpdateSetProto",
			Handler:    _SetService_UpdateSetProto_Handler,
		},
		{
			MethodName: "DeleteSetProto",
			Handler:    _SetService_DeleteSetProto_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "quanqr/proto_qr/set/set.proto",
}
