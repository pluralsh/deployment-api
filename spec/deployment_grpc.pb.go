// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.12.4
// source: spec/deployment.proto

package deployment

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

// IdentityClient is the client API for Identity service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type IdentityClient interface {
	// This call is meant to retrieve the unique provisioner Identity.
	// This identity will have to be set in DeploymentClaim.DriverName field in order to invoke this specific provisioner.
	DriverGetInfo(ctx context.Context, in *DriverGetInfoRequest, opts ...grpc.CallOption) (*DriverGetInfoResponse, error)
}

type identityClient struct {
	cc grpc.ClientConnInterface
}

func NewIdentityClient(cc grpc.ClientConnInterface) IdentityClient {
	return &identityClient{cc}
}

func (c *identityClient) DriverGetInfo(ctx context.Context, in *DriverGetInfoRequest, opts ...grpc.CallOption) (*DriverGetInfoResponse, error) {
	out := new(DriverGetInfoResponse)
	err := c.cc.Invoke(ctx, "/deployment.v1alpha1.Identity/DriverGetInfo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// IdentityServer is the server API for Identity service.
// All implementations should embed UnimplementedIdentityServer
// for forward compatibility
type IdentityServer interface {
	// This call is meant to retrieve the unique provisioner Identity.
	// This identity will have to be set in DeploymentClaim.DriverName field in order to invoke this specific provisioner.
	DriverGetInfo(context.Context, *DriverGetInfoRequest) (*DriverGetInfoResponse, error)
}

// UnimplementedIdentityServer should be embedded to have forward compatible implementations.
type UnimplementedIdentityServer struct {
}

func (UnimplementedIdentityServer) DriverGetInfo(context.Context, *DriverGetInfoRequest) (*DriverGetInfoResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DriverGetInfo not implemented")
}

// UnsafeIdentityServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to IdentityServer will
// result in compilation errors.
type UnsafeIdentityServer interface {
	mustEmbedUnimplementedIdentityServer()
}

func RegisterIdentityServer(s grpc.ServiceRegistrar, srv IdentityServer) {
	s.RegisterService(&Identity_ServiceDesc, srv)
}

func _Identity_DriverGetInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DriverGetInfoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IdentityServer).DriverGetInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/deployment.v1alpha1.Identity/DriverGetInfo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IdentityServer).DriverGetInfo(ctx, req.(*DriverGetInfoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Identity_ServiceDesc is the grpc.ServiceDesc for Identity service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Identity_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "deployment.v1alpha1.Identity",
	HandlerType: (*IdentityServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "DriverGetInfo",
			Handler:    _Identity_DriverGetInfo_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "spec/deployment.proto",
}

// ProvisionerClient is the client API for Provisioner service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ProvisionerClient interface {
	// This call is made to create the deployment in the backend.
	// This call is idempotent
	//    1. If a deployment that matches both name and parameters already exists, then OK (success) must be returned.
	//    2. If a deployment by same name, but different parameters is provided, then the appropriate error code ALREADY_EXISTS must be returned.
	DriverCreateDeployment(ctx context.Context, in *DriverCreateDeploymentRequest, opts ...grpc.CallOption) (*DriverCreateDeploymentResponse, error)
	// This call is made to delete the deployment in the backend.
	// If the deployment has already been deleted, then no error should be returned.
	DriverDeleteDeployment(ctx context.Context, in *DriverDeleteDeploymentRequest, opts ...grpc.CallOption) (*DriverDeleteDeploymentResponse, error)
	DriverGetDeploymentStatus(ctx context.Context, in *DriverGetDeploymentStatusRequest, opts ...grpc.CallOption) (*DriverGetDeploymentStatusResponse, error)
}

type provisionerClient struct {
	cc grpc.ClientConnInterface
}

func NewProvisionerClient(cc grpc.ClientConnInterface) ProvisionerClient {
	return &provisionerClient{cc}
}

func (c *provisionerClient) DriverCreateDeployment(ctx context.Context, in *DriverCreateDeploymentRequest, opts ...grpc.CallOption) (*DriverCreateDeploymentResponse, error) {
	out := new(DriverCreateDeploymentResponse)
	err := c.cc.Invoke(ctx, "/deployment.v1alpha1.Provisioner/DriverCreateDeployment", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *provisionerClient) DriverDeleteDeployment(ctx context.Context, in *DriverDeleteDeploymentRequest, opts ...grpc.CallOption) (*DriverDeleteDeploymentResponse, error) {
	out := new(DriverDeleteDeploymentResponse)
	err := c.cc.Invoke(ctx, "/deployment.v1alpha1.Provisioner/DriverDeleteDeployment", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *provisionerClient) DriverGetDeploymentStatus(ctx context.Context, in *DriverGetDeploymentStatusRequest, opts ...grpc.CallOption) (*DriverGetDeploymentStatusResponse, error) {
	out := new(DriverGetDeploymentStatusResponse)
	err := c.cc.Invoke(ctx, "/deployment.v1alpha1.Provisioner/DriverGetDeploymentStatus", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ProvisionerServer is the server API for Provisioner service.
// All implementations should embed UnimplementedProvisionerServer
// for forward compatibility
type ProvisionerServer interface {
	// This call is made to create the deployment in the backend.
	// This call is idempotent
	//    1. If a deployment that matches both name and parameters already exists, then OK (success) must be returned.
	//    2. If a deployment by same name, but different parameters is provided, then the appropriate error code ALREADY_EXISTS must be returned.
	DriverCreateDeployment(context.Context, *DriverCreateDeploymentRequest) (*DriverCreateDeploymentResponse, error)
	// This call is made to delete the deployment in the backend.
	// If the deployment has already been deleted, then no error should be returned.
	DriverDeleteDeployment(context.Context, *DriverDeleteDeploymentRequest) (*DriverDeleteDeploymentResponse, error)
	DriverGetDeploymentStatus(context.Context, *DriverGetDeploymentStatusRequest) (*DriverGetDeploymentStatusResponse, error)
}

// UnimplementedProvisionerServer should be embedded to have forward compatible implementations.
type UnimplementedProvisionerServer struct {
}

func (UnimplementedProvisionerServer) DriverCreateDeployment(context.Context, *DriverCreateDeploymentRequest) (*DriverCreateDeploymentResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DriverCreateDeployment not implemented")
}
func (UnimplementedProvisionerServer) DriverDeleteDeployment(context.Context, *DriverDeleteDeploymentRequest) (*DriverDeleteDeploymentResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DriverDeleteDeployment not implemented")
}
func (UnimplementedProvisionerServer) DriverGetDeploymentStatus(context.Context, *DriverGetDeploymentStatusRequest) (*DriverGetDeploymentStatusResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DriverGetDeploymentStatus not implemented")
}

// UnsafeProvisionerServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ProvisionerServer will
// result in compilation errors.
type UnsafeProvisionerServer interface {
	mustEmbedUnimplementedProvisionerServer()
}

func RegisterProvisionerServer(s grpc.ServiceRegistrar, srv ProvisionerServer) {
	s.RegisterService(&Provisioner_ServiceDesc, srv)
}

func _Provisioner_DriverCreateDeployment_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DriverCreateDeploymentRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProvisionerServer).DriverCreateDeployment(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/deployment.v1alpha1.Provisioner/DriverCreateDeployment",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProvisionerServer).DriverCreateDeployment(ctx, req.(*DriverCreateDeploymentRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Provisioner_DriverDeleteDeployment_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DriverDeleteDeploymentRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProvisionerServer).DriverDeleteDeployment(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/deployment.v1alpha1.Provisioner/DriverDeleteDeployment",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProvisionerServer).DriverDeleteDeployment(ctx, req.(*DriverDeleteDeploymentRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Provisioner_DriverGetDeploymentStatus_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DriverGetDeploymentStatusRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProvisionerServer).DriverGetDeploymentStatus(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/deployment.v1alpha1.Provisioner/DriverGetDeploymentStatus",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProvisionerServer).DriverGetDeploymentStatus(ctx, req.(*DriverGetDeploymentStatusRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Provisioner_ServiceDesc is the grpc.ServiceDesc for Provisioner service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Provisioner_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "deployment.v1alpha1.Provisioner",
	HandlerType: (*ProvisionerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "DriverCreateDeployment",
			Handler:    _Provisioner_DriverCreateDeployment_Handler,
		},
		{
			MethodName: "DriverDeleteDeployment",
			Handler:    _Provisioner_DriverDeleteDeployment_Handler,
		},
		{
			MethodName: "DriverGetDeploymentStatus",
			Handler:    _Provisioner_DriverGetDeploymentStatus_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "spec/deployment.proto",
}
