// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.24.0
// source: source.proto

package source

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

const (
	Source_Stream_FullMethodName               = "/source.Source/Stream"
	Source_HandleSingleDispatch_FullMethodName = "/source.Source/HandleSingleDispatch"
	Source_Metadata_FullMethodName             = "/source.Source/Metadata"
)

// SourceClient is the client API for Source service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type SourceClient interface {
	Stream(ctx context.Context, in *StreamRequest, opts ...grpc.CallOption) (Source_StreamClient, error)
	HandleSingleDispatch(ctx context.Context, in *SingleDispatchRequest, opts ...grpc.CallOption) (*SingleDispatchResponse, error)
	Metadata(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*MetadataResponse, error)
}

type sourceClient struct {
	cc grpc.ClientConnInterface
}

func NewSourceClient(cc grpc.ClientConnInterface) SourceClient {
	return &sourceClient{cc}
}

func (c *sourceClient) Stream(ctx context.Context, in *StreamRequest, opts ...grpc.CallOption) (Source_StreamClient, error) {
	stream, err := c.cc.NewStream(ctx, &Source_ServiceDesc.Streams[0], Source_Stream_FullMethodName, opts...)
	if err != nil {
		return nil, err
	}
	x := &sourceStreamClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Source_StreamClient interface {
	Recv() (*StreamResponse, error)
	grpc.ClientStream
}

type sourceStreamClient struct {
	grpc.ClientStream
}

func (x *sourceStreamClient) Recv() (*StreamResponse, error) {
	m := new(StreamResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *sourceClient) HandleSingleDispatch(ctx context.Context, in *SingleDispatchRequest, opts ...grpc.CallOption) (*SingleDispatchResponse, error) {
	out := new(SingleDispatchResponse)
	err := c.cc.Invoke(ctx, Source_HandleSingleDispatch_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sourceClient) Metadata(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*MetadataResponse, error) {
	out := new(MetadataResponse)
	err := c.cc.Invoke(ctx, Source_Metadata_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SourceServer is the server API for Source service.
// All implementations must embed UnimplementedSourceServer
// for forward compatibility
type SourceServer interface {
	Stream(*StreamRequest, Source_StreamServer) error
	HandleSingleDispatch(context.Context, *SingleDispatchRequest) (*SingleDispatchResponse, error)
	Metadata(context.Context, *emptypb.Empty) (*MetadataResponse, error)
	mustEmbedUnimplementedSourceServer()
}

// UnimplementedSourceServer must be embedded to have forward compatible implementations.
type UnimplementedSourceServer struct {
}

func (UnimplementedSourceServer) Stream(*StreamRequest, Source_StreamServer) error {
	return status.Errorf(codes.Unimplemented, "method Stream not implemented")
}
func (UnimplementedSourceServer) HandleSingleDispatch(context.Context, *SingleDispatchRequest) (*SingleDispatchResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method HandleSingleDispatch not implemented")
}
func (UnimplementedSourceServer) Metadata(context.Context, *emptypb.Empty) (*MetadataResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Metadata not implemented")
}
func (UnimplementedSourceServer) mustEmbedUnimplementedSourceServer() {}

// UnsafeSourceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to SourceServer will
// result in compilation errors.
type UnsafeSourceServer interface {
	mustEmbedUnimplementedSourceServer()
}

func RegisterSourceServer(s grpc.ServiceRegistrar, srv SourceServer) {
	s.RegisterService(&Source_ServiceDesc, srv)
}

func _Source_Stream_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(StreamRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(SourceServer).Stream(m, &sourceStreamServer{stream})
}

type Source_StreamServer interface {
	Send(*StreamResponse) error
	grpc.ServerStream
}

type sourceStreamServer struct {
	grpc.ServerStream
}

func (x *sourceStreamServer) Send(m *StreamResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _Source_HandleSingleDispatch_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SingleDispatchRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SourceServer).HandleSingleDispatch(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Source_HandleSingleDispatch_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SourceServer).HandleSingleDispatch(ctx, req.(*SingleDispatchRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Source_Metadata_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SourceServer).Metadata(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Source_Metadata_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SourceServer).Metadata(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

// Source_ServiceDesc is the grpc.ServiceDesc for Source service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Source_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "source.Source",
	HandlerType: (*SourceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "HandleSingleDispatch",
			Handler:    _Source_HandleSingleDispatch_Handler,
		},
		{
			MethodName: "Metadata",
			Handler:    _Source_Metadata_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Stream",
			Handler:       _Source_Stream_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "source.proto",
}
