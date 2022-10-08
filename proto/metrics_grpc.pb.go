// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.6
// source: metrics.proto

package __

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// MetricsClient is the client API for Metrics service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MetricsClient interface {
	GetCPU(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*CPU, error)
	GetMemory(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*Memory, error)
	GetStorage(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*Storage, error)
	GetNetwork(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*Network, error)
	GetOS(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*OS, error)
	GetKernel(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*Kernel, error)
}

type metricsClient struct {
	cc grpc.ClientConnInterface
}

func NewMetricsClient(cc grpc.ClientConnInterface) MetricsClient {
	return &metricsClient{cc}
}

func (c *metricsClient) GetCPU(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*CPU, error) {
	out := new(CPU)
	err := c.cc.Invoke(ctx, "/proto.Metrics/GetCPU", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *metricsClient) GetMemory(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*Memory, error) {
	out := new(Memory)
	err := c.cc.Invoke(ctx, "/proto.Metrics/GetMemory", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *metricsClient) GetStorage(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*Storage, error) {
	out := new(Storage)
	err := c.cc.Invoke(ctx, "/proto.Metrics/GetStorage", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *metricsClient) GetNetwork(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*Network, error) {
	out := new(Network)
	err := c.cc.Invoke(ctx, "/proto.Metrics/GetNetwork", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *metricsClient) GetOS(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*OS, error) {
	out := new(OS)
	err := c.cc.Invoke(ctx, "/proto.Metrics/GetOS", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *metricsClient) GetKernel(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*Kernel, error) {
	out := new(Kernel)
	err := c.cc.Invoke(ctx, "/proto.Metrics/GetKernel", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MetricsServer is the server API for Metrics service.
// All implementations must embed UnimplementedMetricsServer
// for forward compatibility
type MetricsServer interface {
	GetCPU(context.Context, *emptypb.Empty) (*CPU, error)
	GetMemory(context.Context, *emptypb.Empty) (*Memory, error)
	GetStorage(context.Context, *emptypb.Empty) (*Storage, error)
	GetNetwork(context.Context, *emptypb.Empty) (*Network, error)
	GetOS(context.Context, *emptypb.Empty) (*OS, error)
	GetKernel(context.Context, *emptypb.Empty) (*Kernel, error)
	mustEmbedUnimplementedMetricsServer()
}

// UnimplementedMetricsServer must be embedded to have forward compatible implementations.
type UnimplementedMetricsServer struct {
}

func (UnimplementedMetricsServer) GetCPU(context.Context, *emptypb.Empty) (*CPU, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCPU not implemented")
}
func (UnimplementedMetricsServer) GetMemory(context.Context, *emptypb.Empty) (*Memory, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetMemory not implemented")
}
func (UnimplementedMetricsServer) GetStorage(context.Context, *emptypb.Empty) (*Storage, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetStorage not implemented")
}
func (UnimplementedMetricsServer) GetNetwork(context.Context, *emptypb.Empty) (*Network, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetNetwork not implemented")
}
func (UnimplementedMetricsServer) GetOS(context.Context, *emptypb.Empty) (*OS, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetOS not implemented")
}
func (UnimplementedMetricsServer) GetKernel(context.Context, *emptypb.Empty) (*Kernel, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetKernel not implemented")
}
func (UnimplementedMetricsServer) mustEmbedUnimplementedMetricsServer() {}

// UnsafeMetricsServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to MetricsServer will
// result in compilation errors.
type UnsafeMetricsServer interface {
	mustEmbedUnimplementedMetricsServer()
}

func RegisterMetricsServer(s grpc.ServiceRegistrar, srv MetricsServer) {
	s.RegisterService(&Metrics_ServiceDesc, srv)
}

func _Metrics_GetCPU_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MetricsServer).GetCPU(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Metrics/GetCPU",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MetricsServer).GetCPU(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _Metrics_GetMemory_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MetricsServer).GetMemory(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Metrics/GetMemory",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MetricsServer).GetMemory(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _Metrics_GetStorage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MetricsServer).GetStorage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Metrics/GetStorage",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MetricsServer).GetStorage(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _Metrics_GetNetwork_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MetricsServer).GetNetwork(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Metrics/GetNetwork",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MetricsServer).GetNetwork(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _Metrics_GetOS_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MetricsServer).GetOS(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Metrics/GetOS",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MetricsServer).GetOS(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _Metrics_GetKernel_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MetricsServer).GetKernel(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Metrics/GetKernel",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MetricsServer).GetKernel(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

// Metrics_ServiceDesc is the grpc.ServiceDesc for Metrics service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Metrics_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "proto.Metrics",
	HandlerType: (*MetricsServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetCPU",
			Handler:    _Metrics_GetCPU_Handler,
		},
		{
			MethodName: "GetMemory",
			Handler:    _Metrics_GetMemory_Handler,
		},
		{
			MethodName: "GetStorage",
			Handler:    _Metrics_GetStorage_Handler,
		},
		{
			MethodName: "GetNetwork",
			Handler:    _Metrics_GetNetwork_Handler,
		},
		{
			MethodName: "GetOS",
			Handler:    _Metrics_GetOS_Handler,
		},
		{
			MethodName: "GetKernel",
			Handler:    _Metrics_GetKernel_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "metrics.proto",
}