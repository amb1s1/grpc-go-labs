// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package proto

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

// CLIClient is the client API for CLI service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CLIClient interface {
	GetLS(ctx context.Context, in *LSRequest, opts ...grpc.CallOption) (*LSResponse, error)
}

type cLIClient struct {
	cc grpc.ClientConnInterface
}

func NewCLIClient(cc grpc.ClientConnInterface) CLIClient {
	return &cLIClient{cc}
}

func (c *cLIClient) GetLS(ctx context.Context, in *LSRequest, opts ...grpc.CallOption) (*LSResponse, error) {
	out := new(LSResponse)
	err := c.cc.Invoke(ctx, "/cli.CLI/GetLS", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CLIServer is the server API for CLI service.
// All implementations must embed UnimplementedCLIServer
// for forward compatibility
type CLIServer interface {
	GetLS(context.Context, *LSRequest) (*LSResponse, error)
	mustEmbedUnimplementedCLIServer()
}

// UnimplementedCLIServer must be embedded to have forward compatible implementations.
type UnimplementedCLIServer struct {
}

func (UnimplementedCLIServer) GetLS(context.Context, *LSRequest) (*LSResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetLS not implemented")
}
func (UnimplementedCLIServer) mustEmbedUnimplementedCLIServer() {}

// UnsafeCLIServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CLIServer will
// result in compilation errors.
type UnsafeCLIServer interface {
	mustEmbedUnimplementedCLIServer()
}

func RegisterCLIServer(s grpc.ServiceRegistrar, srv CLIServer) {
	s.RegisterService(&CLI_ServiceDesc, srv)
}

func _CLI_GetLS_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LSRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CLIServer).GetLS(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cli.CLI/GetLS",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CLIServer).GetLS(ctx, req.(*LSRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// CLI_ServiceDesc is the grpc.ServiceDesc for CLI service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var CLI_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "cli.CLI",
	HandlerType: (*CLIServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetLS",
			Handler:    _CLI_GetLS_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/cli.proto",
}
