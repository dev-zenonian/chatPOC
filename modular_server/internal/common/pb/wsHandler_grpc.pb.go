// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.15.8
// source: wsHandler.proto

package pb

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
	WSHandlerService_PassMessageToClient_FullMethodName = "/proto.WSHandlerService/PassMessageToClient"
)

// WSHandlerServiceClient is the client API for WSHandlerService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type WSHandlerServiceClient interface {
	PassMessageToClient(ctx context.Context, in *PassMessageToClientRequest, opts ...grpc.CallOption) (*PassMessageToClientResponse, error)
}

type wSHandlerServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewWSHandlerServiceClient(cc grpc.ClientConnInterface) WSHandlerServiceClient {
	return &wSHandlerServiceClient{cc}
}

func (c *wSHandlerServiceClient) PassMessageToClient(ctx context.Context, in *PassMessageToClientRequest, opts ...grpc.CallOption) (*PassMessageToClientResponse, error) {
	out := new(PassMessageToClientResponse)
	err := c.cc.Invoke(ctx, WSHandlerService_PassMessageToClient_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// WSHandlerServiceServer is the server API for WSHandlerService service.
// All implementations must embed UnimplementedWSHandlerServiceServer
// for forward compatibility
type WSHandlerServiceServer interface {
	PassMessageToClient(context.Context, *PassMessageToClientRequest) (*PassMessageToClientResponse, error)
	mustEmbedUnimplementedWSHandlerServiceServer()
}

// UnimplementedWSHandlerServiceServer must be embedded to have forward compatible implementations.
type UnimplementedWSHandlerServiceServer struct {
}

func (UnimplementedWSHandlerServiceServer) PassMessageToClient(context.Context, *PassMessageToClientRequest) (*PassMessageToClientResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PassMessageToClient not implemented")
}
func (UnimplementedWSHandlerServiceServer) mustEmbedUnimplementedWSHandlerServiceServer() {}

// UnsafeWSHandlerServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to WSHandlerServiceServer will
// result in compilation errors.
type UnsafeWSHandlerServiceServer interface {
	mustEmbedUnimplementedWSHandlerServiceServer()
}

func RegisterWSHandlerServiceServer(s grpc.ServiceRegistrar, srv WSHandlerServiceServer) {
	s.RegisterService(&WSHandlerService_ServiceDesc, srv)
}

func _WSHandlerService_PassMessageToClient_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PassMessageToClientRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WSHandlerServiceServer).PassMessageToClient(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: WSHandlerService_PassMessageToClient_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WSHandlerServiceServer).PassMessageToClient(ctx, req.(*PassMessageToClientRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// WSHandlerService_ServiceDesc is the grpc.ServiceDesc for WSHandlerService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var WSHandlerService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "proto.WSHandlerService",
	HandlerType: (*WSHandlerServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "PassMessageToClient",
			Handler:    _WSHandlerService_PassMessageToClient_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "wsHandler.proto",
}
