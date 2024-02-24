// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v4.24.4
// source: chromadb/proto/logservice.proto

package logservicepb

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

// LogServiceClient is the client API for LogService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type LogServiceClient interface {
	PushLogs(ctx context.Context, in *PushLogsRequest, opts ...grpc.CallOption) (*PushLogsResponse, error)
	PullLogs(ctx context.Context, in *PullLogsRequest, opts ...grpc.CallOption) (*PullLogsResponse, error)
}

type logServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewLogServiceClient(cc grpc.ClientConnInterface) LogServiceClient {
	return &logServiceClient{cc}
}

func (c *logServiceClient) PushLogs(ctx context.Context, in *PushLogsRequest, opts ...grpc.CallOption) (*PushLogsResponse, error) {
	out := new(PushLogsResponse)
	err := c.cc.Invoke(ctx, "/chroma.LogService/PushLogs", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *logServiceClient) PullLogs(ctx context.Context, in *PullLogsRequest, opts ...grpc.CallOption) (*PullLogsResponse, error) {
	out := new(PullLogsResponse)
	err := c.cc.Invoke(ctx, "/chroma.LogService/PullLogs", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// LogServiceServer is the server API for LogService service.
// All implementations must embed UnimplementedLogServiceServer
// for forward compatibility
type LogServiceServer interface {
	PushLogs(context.Context, *PushLogsRequest) (*PushLogsResponse, error)
	PullLogs(context.Context, *PullLogsRequest) (*PullLogsResponse, error)
	mustEmbedUnimplementedLogServiceServer()
}

// UnimplementedLogServiceServer must be embedded to have forward compatible implementations.
type UnimplementedLogServiceServer struct {
}

func (UnimplementedLogServiceServer) PushLogs(context.Context, *PushLogsRequest) (*PushLogsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PushLogs not implemented")
}
func (UnimplementedLogServiceServer) PullLogs(context.Context, *PullLogsRequest) (*PullLogsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PullLogs not implemented")
}
func (UnimplementedLogServiceServer) mustEmbedUnimplementedLogServiceServer() {}

// UnsafeLogServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to LogServiceServer will
// result in compilation errors.
type UnsafeLogServiceServer interface {
	mustEmbedUnimplementedLogServiceServer()
}

func RegisterLogServiceServer(s grpc.ServiceRegistrar, srv LogServiceServer) {
	s.RegisterService(&LogService_ServiceDesc, srv)
}

func _LogService_PushLogs_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PushLogsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LogServiceServer).PushLogs(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/chroma.LogService/PushLogs",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LogServiceServer).PushLogs(ctx, req.(*PushLogsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _LogService_PullLogs_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PullLogsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LogServiceServer).PullLogs(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/chroma.LogService/PullLogs",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LogServiceServer).PullLogs(ctx, req.(*PullLogsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// LogService_ServiceDesc is the grpc.ServiceDesc for LogService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var LogService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "chroma.LogService",
	HandlerType: (*LogServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "PushLogs",
			Handler:    _LogService_PushLogs_Handler,
		},
		{
			MethodName: "PullLogs",
			Handler:    _LogService_PullLogs_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "chromadb/proto/logservice.proto",
}
