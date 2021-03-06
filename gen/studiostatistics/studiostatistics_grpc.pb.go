// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package studiostatistics

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

// StudioStatisticsServiceClient is the client API for StudioStatisticsService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type StudioStatisticsServiceClient interface {
	UploadStatistics(ctx context.Context, in *StudioStatisticsRequest, opts ...grpc.CallOption) (*StudioStatisticsResponse, error)
}

type studioStatisticsServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewStudioStatisticsServiceClient(cc grpc.ClientConnInterface) StudioStatisticsServiceClient {
	return &studioStatisticsServiceClient{cc}
}

func (c *studioStatisticsServiceClient) UploadStatistics(ctx context.Context, in *StudioStatisticsRequest, opts ...grpc.CallOption) (*StudioStatisticsResponse, error) {
	out := new(StudioStatisticsResponse)
	err := c.cc.Invoke(ctx, "/studiostatistics.StudioStatisticsService/UploadStatistics", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// StudioStatisticsServiceServer is the server API for StudioStatisticsService service.
// All implementations must embed UnimplementedStudioStatisticsServiceServer
// for forward compatibility
type StudioStatisticsServiceServer interface {
	UploadStatistics(context.Context, *StudioStatisticsRequest) (*StudioStatisticsResponse, error)
	mustEmbedUnimplementedStudioStatisticsServiceServer()
}

// UnimplementedStudioStatisticsServiceServer must be embedded to have forward compatible implementations.
type UnimplementedStudioStatisticsServiceServer struct {
}

func (UnimplementedStudioStatisticsServiceServer) UploadStatistics(context.Context, *StudioStatisticsRequest) (*StudioStatisticsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UploadStatistics not implemented")
}
func (UnimplementedStudioStatisticsServiceServer) mustEmbedUnimplementedStudioStatisticsServiceServer() {
}

// UnsafeStudioStatisticsServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to StudioStatisticsServiceServer will
// result in compilation errors.
type UnsafeStudioStatisticsServiceServer interface {
	mustEmbedUnimplementedStudioStatisticsServiceServer()
}

func RegisterStudioStatisticsServiceServer(s grpc.ServiceRegistrar, srv StudioStatisticsServiceServer) {
	s.RegisterService(&StudioStatisticsService_ServiceDesc, srv)
}

func _StudioStatisticsService_UploadStatistics_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StudioStatisticsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StudioStatisticsServiceServer).UploadStatistics(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/studiostatistics.StudioStatisticsService/UploadStatistics",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StudioStatisticsServiceServer).UploadStatistics(ctx, req.(*StudioStatisticsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// StudioStatisticsService_ServiceDesc is the grpc.ServiceDesc for StudioStatisticsService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var StudioStatisticsService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "studiostatistics.StudioStatisticsService",
	HandlerType: (*StudioStatisticsServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "UploadStatistics",
			Handler:    _StudioStatisticsService_UploadStatistics_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/studiostatistics/studiostatistics.proto",
}
