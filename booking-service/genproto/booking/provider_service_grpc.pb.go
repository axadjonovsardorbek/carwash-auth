// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v5.26.1
// source: provider_service.proto

package booking

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
	ProviderServiceService_Create_FullMethodName  = "/booking.ProviderServiceService/Create"
	ProviderServiceService_GetById_FullMethodName = "/booking.ProviderServiceService/GetById"
	ProviderServiceService_GetAll_FullMethodName  = "/booking.ProviderServiceService/GetAll"
	ProviderServiceService_Delete_FullMethodName  = "/booking.ProviderServiceService/Delete"
)

// ProviderServiceServiceClient is the client API for ProviderServiceService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ProviderServiceServiceClient interface {
	Create(ctx context.Context, in *ProviderServiceRes, opts ...grpc.CallOption) (*Void, error)
	GetById(ctx context.Context, in *ById, opts ...grpc.CallOption) (*ProviderServiceGetByIdRes, error)
	GetAll(ctx context.Context, in *ProviderServiceGetAllReq, opts ...grpc.CallOption) (*ProviderServiceGetAllRes, error)
	Delete(ctx context.Context, in *ById, opts ...grpc.CallOption) (*Void, error)
}

type providerServiceServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewProviderServiceServiceClient(cc grpc.ClientConnInterface) ProviderServiceServiceClient {
	return &providerServiceServiceClient{cc}
}

func (c *providerServiceServiceClient) Create(ctx context.Context, in *ProviderServiceRes, opts ...grpc.CallOption) (*Void, error) {
	out := new(Void)
	err := c.cc.Invoke(ctx, ProviderServiceService_Create_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *providerServiceServiceClient) GetById(ctx context.Context, in *ById, opts ...grpc.CallOption) (*ProviderServiceGetByIdRes, error) {
	out := new(ProviderServiceGetByIdRes)
	err := c.cc.Invoke(ctx, ProviderServiceService_GetById_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *providerServiceServiceClient) GetAll(ctx context.Context, in *ProviderServiceGetAllReq, opts ...grpc.CallOption) (*ProviderServiceGetAllRes, error) {
	out := new(ProviderServiceGetAllRes)
	err := c.cc.Invoke(ctx, ProviderServiceService_GetAll_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *providerServiceServiceClient) Delete(ctx context.Context, in *ById, opts ...grpc.CallOption) (*Void, error) {
	out := new(Void)
	err := c.cc.Invoke(ctx, ProviderServiceService_Delete_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ProviderServiceServiceServer is the server API for ProviderServiceService service.
// All implementations must embed UnimplementedProviderServiceServiceServer
// for forward compatibility
type ProviderServiceServiceServer interface {
	Create(context.Context, *ProviderServiceRes) (*Void, error)
	GetById(context.Context, *ById) (*ProviderServiceGetByIdRes, error)
	GetAll(context.Context, *ProviderServiceGetAllReq) (*ProviderServiceGetAllRes, error)
	Delete(context.Context, *ById) (*Void, error)
	mustEmbedUnimplementedProviderServiceServiceServer()
}

// UnimplementedProviderServiceServiceServer must be embedded to have forward compatible implementations.
type UnimplementedProviderServiceServiceServer struct {
}

func (UnimplementedProviderServiceServiceServer) Create(context.Context, *ProviderServiceRes) (*Void, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (UnimplementedProviderServiceServiceServer) GetById(context.Context, *ById) (*ProviderServiceGetByIdRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetById not implemented")
}
func (UnimplementedProviderServiceServiceServer) GetAll(context.Context, *ProviderServiceGetAllReq) (*ProviderServiceGetAllRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAll not implemented")
}
func (UnimplementedProviderServiceServiceServer) Delete(context.Context, *ById) (*Void, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}
func (UnimplementedProviderServiceServiceServer) mustEmbedUnimplementedProviderServiceServiceServer() {
}

// UnsafeProviderServiceServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ProviderServiceServiceServer will
// result in compilation errors.
type UnsafeProviderServiceServiceServer interface {
	mustEmbedUnimplementedProviderServiceServiceServer()
}

func RegisterProviderServiceServiceServer(s grpc.ServiceRegistrar, srv ProviderServiceServiceServer) {
	s.RegisterService(&ProviderServiceService_ServiceDesc, srv)
}

func _ProviderServiceService_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ProviderServiceRes)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProviderServiceServiceServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ProviderServiceService_Create_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProviderServiceServiceServer).Create(ctx, req.(*ProviderServiceRes))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProviderServiceService_GetById_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ById)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProviderServiceServiceServer).GetById(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ProviderServiceService_GetById_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProviderServiceServiceServer).GetById(ctx, req.(*ById))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProviderServiceService_GetAll_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ProviderServiceGetAllReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProviderServiceServiceServer).GetAll(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ProviderServiceService_GetAll_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProviderServiceServiceServer).GetAll(ctx, req.(*ProviderServiceGetAllReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProviderServiceService_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ById)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProviderServiceServiceServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ProviderServiceService_Delete_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProviderServiceServiceServer).Delete(ctx, req.(*ById))
	}
	return interceptor(ctx, in, info, handler)
}

// ProviderServiceService_ServiceDesc is the grpc.ServiceDesc for ProviderServiceService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ProviderServiceService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "booking.ProviderServiceService",
	HandlerType: (*ProviderServiceServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Create",
			Handler:    _ProviderServiceService_Create_Handler,
		},
		{
			MethodName: "GetById",
			Handler:    _ProviderServiceService_GetById_Handler,
		},
		{
			MethodName: "GetAll",
			Handler:    _ProviderServiceService_GetAll_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _ProviderServiceService_Delete_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "provider_service.proto",
}
