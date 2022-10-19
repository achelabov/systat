// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.6
// source: stats.proto

package systat

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

// StatsClient is the client API for Stats service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type StatsClient interface {
	GetBatteries(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (Stats_GetBatteriesClient, error)
	GetCpus(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (Stats_GetCpusClient, error)
}

type statsClient struct {
	cc grpc.ClientConnInterface
}

func NewStatsClient(cc grpc.ClientConnInterface) StatsClient {
	return &statsClient{cc}
}

func (c *statsClient) GetBatteries(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (Stats_GetBatteriesClient, error) {
	stream, err := c.cc.NewStream(ctx, &Stats_ServiceDesc.Streams[0], "/Stats/GetBatteries", opts...)
	if err != nil {
		return nil, err
	}
	x := &statsGetBatteriesClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Stats_GetBatteriesClient interface {
	Recv() (*BatteriesResponse, error)
	grpc.ClientStream
}

type statsGetBatteriesClient struct {
	grpc.ClientStream
}

func (x *statsGetBatteriesClient) Recv() (*BatteriesResponse, error) {
	m := new(BatteriesResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *statsClient) GetCpus(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (Stats_GetCpusClient, error) {
	stream, err := c.cc.NewStream(ctx, &Stats_ServiceDesc.Streams[1], "/Stats/GetCpus", opts...)
	if err != nil {
		return nil, err
	}
	x := &statsGetCpusClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Stats_GetCpusClient interface {
	Recv() (*CpusResponse, error)
	grpc.ClientStream
}

type statsGetCpusClient struct {
	grpc.ClientStream
}

func (x *statsGetCpusClient) Recv() (*CpusResponse, error) {
	m := new(CpusResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// StatsServer is the server API for Stats service.
// All implementations must embed UnimplementedStatsServer
// for forward compatibility
type StatsServer interface {
	GetBatteries(*emptypb.Empty, Stats_GetBatteriesServer) error
	GetCpus(*emptypb.Empty, Stats_GetCpusServer) error
	mustEmbedUnimplementedStatsServer()
}

// UnimplementedStatsServer must be embedded to have forward compatible implementations.
type UnimplementedStatsServer struct {
}

func (UnimplementedStatsServer) GetBatteries(*emptypb.Empty, Stats_GetBatteriesServer) error {
	return status.Errorf(codes.Unimplemented, "method GetBatteries not implemented")
}
func (UnimplementedStatsServer) GetCpus(*emptypb.Empty, Stats_GetCpusServer) error {
	return status.Errorf(codes.Unimplemented, "method GetCpus not implemented")
}
func (UnimplementedStatsServer) mustEmbedUnimplementedStatsServer() {}

// UnsafeStatsServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to StatsServer will
// result in compilation errors.
type UnsafeStatsServer interface {
	mustEmbedUnimplementedStatsServer()
}

func RegisterStatsServer(s grpc.ServiceRegistrar, srv StatsServer) {
	s.RegisterService(&Stats_ServiceDesc, srv)
}

func _Stats_GetBatteries_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(emptypb.Empty)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(StatsServer).GetBatteries(m, &statsGetBatteriesServer{stream})
}

type Stats_GetBatteriesServer interface {
	Send(*BatteriesResponse) error
	grpc.ServerStream
}

type statsGetBatteriesServer struct {
	grpc.ServerStream
}

func (x *statsGetBatteriesServer) Send(m *BatteriesResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _Stats_GetCpus_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(emptypb.Empty)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(StatsServer).GetCpus(m, &statsGetCpusServer{stream})
}

type Stats_GetCpusServer interface {
	Send(*CpusResponse) error
	grpc.ServerStream
}

type statsGetCpusServer struct {
	grpc.ServerStream
}

func (x *statsGetCpusServer) Send(m *CpusResponse) error {
	return x.ServerStream.SendMsg(m)
}

// Stats_ServiceDesc is the grpc.ServiceDesc for Stats service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Stats_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "Stats",
	HandlerType: (*StatsServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "GetBatteries",
			Handler:       _Stats_GetBatteries_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "GetCpus",
			Handler:       _Stats_GetCpus_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "stats.proto",
}