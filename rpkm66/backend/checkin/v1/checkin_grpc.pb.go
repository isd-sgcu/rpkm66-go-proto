// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.23.4
// source: rpkm66/backend/checkin/v1/checkin.proto

package v1

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
	CheckinService_CheckinVerify_FullMethodName  = "/rpkm66.backend.checkin.v1.CheckinService/CheckinVerify"
	CheckinService_CheckinConfirm_FullMethodName = "/rpkm66.backend.checkin.v1.CheckinService/CheckinConfirm"
)

// CheckinServiceClient is the client API for CheckinService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CheckinServiceClient interface {
	CheckinVerify(ctx context.Context, in *CheckinVerifyRequest, opts ...grpc.CallOption) (*CheckinVerifyResponse, error)
	CheckinConfirm(ctx context.Context, in *CheckinConfirmRequest, opts ...grpc.CallOption) (*CheckinConfirmResponse, error)
}

type checkinServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewCheckinServiceClient(cc grpc.ClientConnInterface) CheckinServiceClient {
	return &checkinServiceClient{cc}
}

func (c *checkinServiceClient) CheckinVerify(ctx context.Context, in *CheckinVerifyRequest, opts ...grpc.CallOption) (*CheckinVerifyResponse, error) {
	out := new(CheckinVerifyResponse)
	err := c.cc.Invoke(ctx, CheckinService_CheckinVerify_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *checkinServiceClient) CheckinConfirm(ctx context.Context, in *CheckinConfirmRequest, opts ...grpc.CallOption) (*CheckinConfirmResponse, error) {
	out := new(CheckinConfirmResponse)
	err := c.cc.Invoke(ctx, CheckinService_CheckinConfirm_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CheckinServiceServer is the server API for CheckinService service.
// All implementations must embed UnimplementedCheckinServiceServer
// for forward compatibility
type CheckinServiceServer interface {
	CheckinVerify(context.Context, *CheckinVerifyRequest) (*CheckinVerifyResponse, error)
	CheckinConfirm(context.Context, *CheckinConfirmRequest) (*CheckinConfirmResponse, error)
	mustEmbedUnimplementedCheckinServiceServer()
}

// UnimplementedCheckinServiceServer must be embedded to have forward compatible implementations.
type UnimplementedCheckinServiceServer struct {
}

func (UnimplementedCheckinServiceServer) CheckinVerify(context.Context, *CheckinVerifyRequest) (*CheckinVerifyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CheckinVerify not implemented")
}
func (UnimplementedCheckinServiceServer) CheckinConfirm(context.Context, *CheckinConfirmRequest) (*CheckinConfirmResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CheckinConfirm not implemented")
}
func (UnimplementedCheckinServiceServer) mustEmbedUnimplementedCheckinServiceServer() {}

// UnsafeCheckinServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CheckinServiceServer will
// result in compilation errors.
type UnsafeCheckinServiceServer interface {
	mustEmbedUnimplementedCheckinServiceServer()
}

func RegisterCheckinServiceServer(s grpc.ServiceRegistrar, srv CheckinServiceServer) {
	s.RegisterService(&CheckinService_ServiceDesc, srv)
}

func _CheckinService_CheckinVerify_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CheckinVerifyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CheckinServiceServer).CheckinVerify(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CheckinService_CheckinVerify_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CheckinServiceServer).CheckinVerify(ctx, req.(*CheckinVerifyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CheckinService_CheckinConfirm_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CheckinConfirmRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CheckinServiceServer).CheckinConfirm(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CheckinService_CheckinConfirm_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CheckinServiceServer).CheckinConfirm(ctx, req.(*CheckinConfirmRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// CheckinService_ServiceDesc is the grpc.ServiceDesc for CheckinService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var CheckinService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "rpkm66.backend.checkin.v1.CheckinService",
	HandlerType: (*CheckinServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CheckinVerify",
			Handler:    _CheckinService_CheckinVerify_Handler,
		},
		{
			MethodName: "CheckinConfirm",
			Handler:    _CheckinService_CheckinConfirm_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "rpkm66/backend/checkin/v1/checkin.proto",
}
