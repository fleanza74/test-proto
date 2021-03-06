// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package edgekey

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

// EdgeKeyServiceClient is the client API for EdgeKeyService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type EdgeKeyServiceClient interface {
	GenerateKey(ctx context.Context, in *EdgeKeyRequest, opts ...grpc.CallOption) (*EdgeKeyResponse, error)
}

type edgeKeyServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewEdgeKeyServiceClient(cc grpc.ClientConnInterface) EdgeKeyServiceClient {
	return &edgeKeyServiceClient{cc}
}

func (c *edgeKeyServiceClient) GenerateKey(ctx context.Context, in *EdgeKeyRequest, opts ...grpc.CallOption) (*EdgeKeyResponse, error) {
	out := new(EdgeKeyResponse)
	err := c.cc.Invoke(ctx, "/edgekey.EdgeKeyService/GenerateKey", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// EdgeKeyServiceServer is the server API for EdgeKeyService service.
// All implementations must embed UnimplementedEdgeKeyServiceServer
// for forward compatibility
type EdgeKeyServiceServer interface {
	GenerateKey(context.Context, *EdgeKeyRequest) (*EdgeKeyResponse, error)
	mustEmbedUnimplementedEdgeKeyServiceServer()
}

// UnimplementedEdgeKeyServiceServer must be embedded to have forward compatible implementations.
type UnimplementedEdgeKeyServiceServer struct {
}

func (UnimplementedEdgeKeyServiceServer) GenerateKey(context.Context, *EdgeKeyRequest) (*EdgeKeyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GenerateKey not implemented")
}
func (UnimplementedEdgeKeyServiceServer) mustEmbedUnimplementedEdgeKeyServiceServer() {}

// UnsafeEdgeKeyServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to EdgeKeyServiceServer will
// result in compilation errors.
type UnsafeEdgeKeyServiceServer interface {
	mustEmbedUnimplementedEdgeKeyServiceServer()
}

func RegisterEdgeKeyServiceServer(s grpc.ServiceRegistrar, srv EdgeKeyServiceServer) {
	s.RegisterService(&EdgeKeyService_ServiceDesc, srv)
}

func _EdgeKeyService_GenerateKey_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EdgeKeyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EdgeKeyServiceServer).GenerateKey(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/edgekey.EdgeKeyService/GenerateKey",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EdgeKeyServiceServer).GenerateKey(ctx, req.(*EdgeKeyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// EdgeKeyService_ServiceDesc is the grpc.ServiceDesc for EdgeKeyService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var EdgeKeyService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "edgekey.EdgeKeyService",
	HandlerType: (*EdgeKeyServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GenerateKey",
			Handler:    _EdgeKeyService_GenerateKey_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/edgekey/edgekeygenerator.proto",
}
