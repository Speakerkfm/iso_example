// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.19.1
// source: api/address_api.proto

package address_api

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

// AddressServiceClient is the client API for AddressService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AddressServiceClient interface {
	GetUserAddress(ctx context.Context, in *GetUserAddressRequest, opts ...grpc.CallOption) (*GetUserAddressResponse, error)
}

type addressServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewAddressServiceClient(cc grpc.ClientConnInterface) AddressServiceClient {
	return &addressServiceClient{cc}
}

func (c *addressServiceClient) GetUserAddress(ctx context.Context, in *GetUserAddressRequest, opts ...grpc.CallOption) (*GetUserAddressResponse, error) {
	out := new(GetUserAddressResponse)
	err := c.cc.Invoke(ctx, "/AddressService/GetUserAddress", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AddressServiceServer is the server API for AddressService service.
// All implementations must embed UnimplementedAddressServiceServer
// for forward compatibility
type AddressServiceServer interface {
	GetUserAddress(context.Context, *GetUserAddressRequest) (*GetUserAddressResponse, error)
	mustEmbedUnimplementedAddressServiceServer()
}

// UnimplementedAddressServiceServer must be embedded to have forward compatible implementations.
type UnimplementedAddressServiceServer struct {
}

func (UnimplementedAddressServiceServer) GetUserAddress(context.Context, *GetUserAddressRequest) (*GetUserAddressResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUserAddress not implemented")
}
func (UnimplementedAddressServiceServer) mustEmbedUnimplementedAddressServiceServer() {}

// UnsafeAddressServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AddressServiceServer will
// result in compilation errors.
type UnsafeAddressServiceServer interface {
	mustEmbedUnimplementedAddressServiceServer()
}

func RegisterAddressServiceServer(s grpc.ServiceRegistrar, srv AddressServiceServer) {
	s.RegisterService(&AddressService_ServiceDesc, srv)
}

func _AddressService_GetUserAddress_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetUserAddressRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AddressServiceServer).GetUserAddress(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/AddressService/GetUserAddress",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AddressServiceServer).GetUserAddress(ctx, req.(*GetUserAddressRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// AddressService_ServiceDesc is the grpc.ServiceDesc for AddressService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var AddressService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "AddressService",
	HandlerType: (*AddressServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetUserAddress",
			Handler:    _AddressService_GetUserAddress_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/address_api.proto",
}
