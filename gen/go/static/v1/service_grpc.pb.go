// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v4.23.1
// source: service.proto

package pb_static

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

// StaticClient is the client API for Static service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type StaticClient interface {
	Upload(ctx context.Context, opts ...grpc.CallOption) (Static_UploadClient, error)
	Download(ctx context.Context, in *DownloadRequest, opts ...grpc.CallOption) (Static_DownloadClient, error)
	Delete(ctx context.Context, in *DeleteRequest, opts ...grpc.CallOption) (*DeleteResponse, error)
}

type staticClient struct {
	cc grpc.ClientConnInterface
}

func NewStaticClient(cc grpc.ClientConnInterface) StaticClient {
	return &staticClient{cc}
}

func (c *staticClient) Upload(ctx context.Context, opts ...grpc.CallOption) (Static_UploadClient, error) {
	stream, err := c.cc.NewStream(ctx, &Static_ServiceDesc.Streams[0], "/static.v1.Static/Upload", opts...)
	if err != nil {
		return nil, err
	}
	x := &staticUploadClient{stream}
	return x, nil
}

type Static_UploadClient interface {
	Send(*UploadRequest) error
	CloseAndRecv() (*UploadResponse, error)
	grpc.ClientStream
}

type staticUploadClient struct {
	grpc.ClientStream
}

func (x *staticUploadClient) Send(m *UploadRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *staticUploadClient) CloseAndRecv() (*UploadResponse, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(UploadResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *staticClient) Download(ctx context.Context, in *DownloadRequest, opts ...grpc.CallOption) (Static_DownloadClient, error) {
	stream, err := c.cc.NewStream(ctx, &Static_ServiceDesc.Streams[1], "/static.v1.Static/Download", opts...)
	if err != nil {
		return nil, err
	}
	x := &staticDownloadClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Static_DownloadClient interface {
	Recv() (*DownloadResponse, error)
	grpc.ClientStream
}

type staticDownloadClient struct {
	grpc.ClientStream
}

func (x *staticDownloadClient) Recv() (*DownloadResponse, error) {
	m := new(DownloadResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *staticClient) Delete(ctx context.Context, in *DeleteRequest, opts ...grpc.CallOption) (*DeleteResponse, error) {
	out := new(DeleteResponse)
	err := c.cc.Invoke(ctx, "/static.v1.Static/Delete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// StaticServer is the server API for Static service.
// All implementations must embed UnimplementedStaticServer
// for forward compatibility
type StaticServer interface {
	Upload(Static_UploadServer) error
	Download(*DownloadRequest, Static_DownloadServer) error
	Delete(context.Context, *DeleteRequest) (*DeleteResponse, error)
	mustEmbedUnimplementedStaticServer()
}

// UnimplementedStaticServer must be embedded to have forward compatible implementations.
type UnimplementedStaticServer struct {
}

func (UnimplementedStaticServer) Upload(Static_UploadServer) error {
	return status.Errorf(codes.Unimplemented, "method Upload not implemented")
}
func (UnimplementedStaticServer) Download(*DownloadRequest, Static_DownloadServer) error {
	return status.Errorf(codes.Unimplemented, "method Download not implemented")
}
func (UnimplementedStaticServer) Delete(context.Context, *DeleteRequest) (*DeleteResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}
func (UnimplementedStaticServer) mustEmbedUnimplementedStaticServer() {}

// UnsafeStaticServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to StaticServer will
// result in compilation errors.
type UnsafeStaticServer interface {
	mustEmbedUnimplementedStaticServer()
}

func RegisterStaticServer(s grpc.ServiceRegistrar, srv StaticServer) {
	s.RegisterService(&Static_ServiceDesc, srv)
}

func _Static_Upload_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(StaticServer).Upload(&staticUploadServer{stream})
}

type Static_UploadServer interface {
	SendAndClose(*UploadResponse) error
	Recv() (*UploadRequest, error)
	grpc.ServerStream
}

type staticUploadServer struct {
	grpc.ServerStream
}

func (x *staticUploadServer) SendAndClose(m *UploadResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *staticUploadServer) Recv() (*UploadRequest, error) {
	m := new(UploadRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _Static_Download_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(DownloadRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(StaticServer).Download(m, &staticDownloadServer{stream})
}

type Static_DownloadServer interface {
	Send(*DownloadResponse) error
	grpc.ServerStream
}

type staticDownloadServer struct {
	grpc.ServerStream
}

func (x *staticDownloadServer) Send(m *DownloadResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _Static_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StaticServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/static.v1.Static/Delete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StaticServer).Delete(ctx, req.(*DeleteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Static_ServiceDesc is the grpc.ServiceDesc for Static service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Static_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "static.v1.Static",
	HandlerType: (*StaticServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Delete",
			Handler:    _Static_Delete_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Upload",
			Handler:       _Static_Upload_Handler,
			ClientStreams: true,
		},
		{
			StreamName:    "Download",
			Handler:       _Static_Download_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "service.proto",
}
