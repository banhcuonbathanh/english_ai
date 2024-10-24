// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v4.25.1
// source: quanqr/proto_qr/guest/guest.proto

package guest

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
	GuestService_GuestLoginGRPC_FullMethodName        = "/guest.GuestService/GuestLoginGRPC"
	GuestService_GuestLogoutGRPC_FullMethodName       = "/guest.GuestService/GuestLogoutGRPC"
	GuestService_GuestRefreshTokenGRPC_FullMethodName = "/guest.GuestService/GuestRefreshTokenGRPC"
	GuestService_GuestCreateOrdersGRPC_FullMethodName = "/guest.GuestService/GuestCreateOrdersGRPC"
	GuestService_GuestGetOrdersGRPC_FullMethodName    = "/guest.GuestService/GuestGetOrdersGRPC"
	GuestService_GuestCreateSession_FullMethodName    = "/guest.GuestService/GuestCreateSession"
	GuestService_GuestGetSession_FullMethodName       = "/guest.GuestService/GuestGetSession"
	GuestService_GuestRevokeSession_FullMethodName    = "/guest.GuestService/GuestRevokeSession"
	GuestService_GuestDeleteSession_FullMethodName    = "/guest.GuestService/GuestDeleteSession"
)

// GuestServiceClient is the client API for GuestService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type GuestServiceClient interface {
	GuestLoginGRPC(ctx context.Context, in *GuestLoginRequest, opts ...grpc.CallOption) (*GuestLoginResponse, error)
	GuestLogoutGRPC(ctx context.Context, in *LogoutRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	GuestRefreshTokenGRPC(ctx context.Context, in *RefreshTokenRequest, opts ...grpc.CallOption) (*RefreshTokenResponse, error)
	GuestCreateOrdersGRPC(ctx context.Context, in *GuestCreateOrderRequest, opts ...grpc.CallOption) (*OrdersResponse, error)
	GuestGetOrdersGRPC(ctx context.Context, in *GuestGetOrdersGRPCRequest, opts ...grpc.CallOption) (*ListOrdersResponse, error)
	GuestCreateSession(ctx context.Context, in *GuestSessionReq, opts ...grpc.CallOption) (*GuestSessionRes, error)
	GuestGetSession(ctx context.Context, in *GuestSessionReq, opts ...grpc.CallOption) (*GuestSessionRes, error)
	GuestRevokeSession(ctx context.Context, in *GuestSessionReq, opts ...grpc.CallOption) (*GuestSessionRes, error)
	GuestDeleteSession(ctx context.Context, in *GuestSessionReq, opts ...grpc.CallOption) (*GuestSessionRes, error)
}

type guestServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewGuestServiceClient(cc grpc.ClientConnInterface) GuestServiceClient {
	return &guestServiceClient{cc}
}

func (c *guestServiceClient) GuestLoginGRPC(ctx context.Context, in *GuestLoginRequest, opts ...grpc.CallOption) (*GuestLoginResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GuestLoginResponse)
	err := c.cc.Invoke(ctx, GuestService_GuestLoginGRPC_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *guestServiceClient) GuestLogoutGRPC(ctx context.Context, in *LogoutRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, GuestService_GuestLogoutGRPC_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *guestServiceClient) GuestRefreshTokenGRPC(ctx context.Context, in *RefreshTokenRequest, opts ...grpc.CallOption) (*RefreshTokenResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(RefreshTokenResponse)
	err := c.cc.Invoke(ctx, GuestService_GuestRefreshTokenGRPC_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *guestServiceClient) GuestCreateOrdersGRPC(ctx context.Context, in *GuestCreateOrderRequest, opts ...grpc.CallOption) (*OrdersResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(OrdersResponse)
	err := c.cc.Invoke(ctx, GuestService_GuestCreateOrdersGRPC_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *guestServiceClient) GuestGetOrdersGRPC(ctx context.Context, in *GuestGetOrdersGRPCRequest, opts ...grpc.CallOption) (*ListOrdersResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ListOrdersResponse)
	err := c.cc.Invoke(ctx, GuestService_GuestGetOrdersGRPC_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *guestServiceClient) GuestCreateSession(ctx context.Context, in *GuestSessionReq, opts ...grpc.CallOption) (*GuestSessionRes, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GuestSessionRes)
	err := c.cc.Invoke(ctx, GuestService_GuestCreateSession_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *guestServiceClient) GuestGetSession(ctx context.Context, in *GuestSessionReq, opts ...grpc.CallOption) (*GuestSessionRes, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GuestSessionRes)
	err := c.cc.Invoke(ctx, GuestService_GuestGetSession_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *guestServiceClient) GuestRevokeSession(ctx context.Context, in *GuestSessionReq, opts ...grpc.CallOption) (*GuestSessionRes, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GuestSessionRes)
	err := c.cc.Invoke(ctx, GuestService_GuestRevokeSession_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *guestServiceClient) GuestDeleteSession(ctx context.Context, in *GuestSessionReq, opts ...grpc.CallOption) (*GuestSessionRes, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GuestSessionRes)
	err := c.cc.Invoke(ctx, GuestService_GuestDeleteSession_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// GuestServiceServer is the server API for GuestService service.
// All implementations must embed UnimplementedGuestServiceServer
// for forward compatibility.
type GuestServiceServer interface {
	GuestLoginGRPC(context.Context, *GuestLoginRequest) (*GuestLoginResponse, error)
	GuestLogoutGRPC(context.Context, *LogoutRequest) (*emptypb.Empty, error)
	GuestRefreshTokenGRPC(context.Context, *RefreshTokenRequest) (*RefreshTokenResponse, error)
	GuestCreateOrdersGRPC(context.Context, *GuestCreateOrderRequest) (*OrdersResponse, error)
	GuestGetOrdersGRPC(context.Context, *GuestGetOrdersGRPCRequest) (*ListOrdersResponse, error)
	GuestCreateSession(context.Context, *GuestSessionReq) (*GuestSessionRes, error)
	GuestGetSession(context.Context, *GuestSessionReq) (*GuestSessionRes, error)
	GuestRevokeSession(context.Context, *GuestSessionReq) (*GuestSessionRes, error)
	GuestDeleteSession(context.Context, *GuestSessionReq) (*GuestSessionRes, error)
	mustEmbedUnimplementedGuestServiceServer()
}

// UnimplementedGuestServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedGuestServiceServer struct{}

func (UnimplementedGuestServiceServer) GuestLoginGRPC(context.Context, *GuestLoginRequest) (*GuestLoginResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GuestLoginGRPC not implemented")
}
func (UnimplementedGuestServiceServer) GuestLogoutGRPC(context.Context, *LogoutRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GuestLogoutGRPC not implemented")
}
func (UnimplementedGuestServiceServer) GuestRefreshTokenGRPC(context.Context, *RefreshTokenRequest) (*RefreshTokenResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GuestRefreshTokenGRPC not implemented")
}
func (UnimplementedGuestServiceServer) GuestCreateOrdersGRPC(context.Context, *GuestCreateOrderRequest) (*OrdersResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GuestCreateOrdersGRPC not implemented")
}
func (UnimplementedGuestServiceServer) GuestGetOrdersGRPC(context.Context, *GuestGetOrdersGRPCRequest) (*ListOrdersResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GuestGetOrdersGRPC not implemented")
}
func (UnimplementedGuestServiceServer) GuestCreateSession(context.Context, *GuestSessionReq) (*GuestSessionRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GuestCreateSession not implemented")
}
func (UnimplementedGuestServiceServer) GuestGetSession(context.Context, *GuestSessionReq) (*GuestSessionRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GuestGetSession not implemented")
}
func (UnimplementedGuestServiceServer) GuestRevokeSession(context.Context, *GuestSessionReq) (*GuestSessionRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GuestRevokeSession not implemented")
}
func (UnimplementedGuestServiceServer) GuestDeleteSession(context.Context, *GuestSessionReq) (*GuestSessionRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GuestDeleteSession not implemented")
}
func (UnimplementedGuestServiceServer) mustEmbedUnimplementedGuestServiceServer() {}
func (UnimplementedGuestServiceServer) testEmbeddedByValue()                      {}

// UnsafeGuestServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to GuestServiceServer will
// result in compilation errors.
type UnsafeGuestServiceServer interface {
	mustEmbedUnimplementedGuestServiceServer()
}

func RegisterGuestServiceServer(s grpc.ServiceRegistrar, srv GuestServiceServer) {
	// If the following call pancis, it indicates UnimplementedGuestServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&GuestService_ServiceDesc, srv)
}

func _GuestService_GuestLoginGRPC_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GuestLoginRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GuestServiceServer).GuestLoginGRPC(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: GuestService_GuestLoginGRPC_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GuestServiceServer).GuestLoginGRPC(ctx, req.(*GuestLoginRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GuestService_GuestLogoutGRPC_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LogoutRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GuestServiceServer).GuestLogoutGRPC(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: GuestService_GuestLogoutGRPC_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GuestServiceServer).GuestLogoutGRPC(ctx, req.(*LogoutRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GuestService_GuestRefreshTokenGRPC_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RefreshTokenRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GuestServiceServer).GuestRefreshTokenGRPC(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: GuestService_GuestRefreshTokenGRPC_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GuestServiceServer).GuestRefreshTokenGRPC(ctx, req.(*RefreshTokenRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GuestService_GuestCreateOrdersGRPC_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GuestCreateOrderRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GuestServiceServer).GuestCreateOrdersGRPC(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: GuestService_GuestCreateOrdersGRPC_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GuestServiceServer).GuestCreateOrdersGRPC(ctx, req.(*GuestCreateOrderRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GuestService_GuestGetOrdersGRPC_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GuestGetOrdersGRPCRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GuestServiceServer).GuestGetOrdersGRPC(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: GuestService_GuestGetOrdersGRPC_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GuestServiceServer).GuestGetOrdersGRPC(ctx, req.(*GuestGetOrdersGRPCRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GuestService_GuestCreateSession_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GuestSessionReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GuestServiceServer).GuestCreateSession(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: GuestService_GuestCreateSession_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GuestServiceServer).GuestCreateSession(ctx, req.(*GuestSessionReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _GuestService_GuestGetSession_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GuestSessionReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GuestServiceServer).GuestGetSession(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: GuestService_GuestGetSession_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GuestServiceServer).GuestGetSession(ctx, req.(*GuestSessionReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _GuestService_GuestRevokeSession_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GuestSessionReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GuestServiceServer).GuestRevokeSession(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: GuestService_GuestRevokeSession_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GuestServiceServer).GuestRevokeSession(ctx, req.(*GuestSessionReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _GuestService_GuestDeleteSession_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GuestSessionReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GuestServiceServer).GuestDeleteSession(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: GuestService_GuestDeleteSession_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GuestServiceServer).GuestDeleteSession(ctx, req.(*GuestSessionReq))
	}
	return interceptor(ctx, in, info, handler)
}

// GuestService_ServiceDesc is the grpc.ServiceDesc for GuestService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var GuestService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "guest.GuestService",
	HandlerType: (*GuestServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GuestLoginGRPC",
			Handler:    _GuestService_GuestLoginGRPC_Handler,
		},
		{
			MethodName: "GuestLogoutGRPC",
			Handler:    _GuestService_GuestLogoutGRPC_Handler,
		},
		{
			MethodName: "GuestRefreshTokenGRPC",
			Handler:    _GuestService_GuestRefreshTokenGRPC_Handler,
		},
		{
			MethodName: "GuestCreateOrdersGRPC",
			Handler:    _GuestService_GuestCreateOrdersGRPC_Handler,
		},
		{
			MethodName: "GuestGetOrdersGRPC",
			Handler:    _GuestService_GuestGetOrdersGRPC_Handler,
		},
		{
			MethodName: "GuestCreateSession",
			Handler:    _GuestService_GuestCreateSession_Handler,
		},
		{
			MethodName: "GuestGetSession",
			Handler:    _GuestService_GuestGetSession_Handler,
		},
		{
			MethodName: "GuestRevokeSession",
			Handler:    _GuestService_GuestRevokeSession_Handler,
		},
		{
			MethodName: "GuestDeleteSession",
			Handler:    _GuestService_GuestDeleteSession_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "quanqr/proto_qr/guest/guest.proto",
}
