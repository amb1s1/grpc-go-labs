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

// ElejirEquipoClient is the client API for ElejirEquipo service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ElejirEquipoClient interface {
	GetCiudad(ctx context.Context, in *CiudadRequest, opts ...grpc.CallOption) (*CiudadRepuesta, error)
}

type elejirEquipoClient struct {
	cc grpc.ClientConnInterface
}

func NewElejirEquipoClient(cc grpc.ClientConnInterface) ElejirEquipoClient {
	return &elejirEquipoClient{cc}
}

func (c *elejirEquipoClient) GetCiudad(ctx context.Context, in *CiudadRequest, opts ...grpc.CallOption) (*CiudadRepuesta, error) {
	out := new(CiudadRepuesta)
	err := c.cc.Invoke(ctx, "/lidom.ElejirEquipo/GetCiudad", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ElejirEquipoServer is the server API for ElejirEquipo service.
// All implementations must embed UnimplementedElejirEquipoServer
// for forward compatibility
type ElejirEquipoServer interface {
	GetCiudad(context.Context, *CiudadRequest) (*CiudadRepuesta, error)
	mustEmbedUnimplementedElejirEquipoServer()
}

// UnimplementedElejirEquipoServer must be embedded to have forward compatible implementations.
type UnimplementedElejirEquipoServer struct {
}

func (UnimplementedElejirEquipoServer) GetCiudad(context.Context, *CiudadRequest) (*CiudadRepuesta, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCiudad not implemented")
}
func (UnimplementedElejirEquipoServer) mustEmbedUnimplementedElejirEquipoServer() {}

// UnsafeElejirEquipoServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ElejirEquipoServer will
// result in compilation errors.
type UnsafeElejirEquipoServer interface {
	mustEmbedUnimplementedElejirEquipoServer()
}

func RegisterElejirEquipoServer(s grpc.ServiceRegistrar, srv ElejirEquipoServer) {
	s.RegisterService(&ElejirEquipo_ServiceDesc, srv)
}

func _ElejirEquipo_GetCiudad_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CiudadRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ElejirEquipoServer).GetCiudad(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/lidom.ElejirEquipo/GetCiudad",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ElejirEquipoServer).GetCiudad(ctx, req.(*CiudadRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ElejirEquipo_ServiceDesc is the grpc.ServiceDesc for ElejirEquipo service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ElejirEquipo_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "lidom.ElejirEquipo",
	HandlerType: (*ElejirEquipoServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetCiudad",
			Handler:    _ElejirEquipo_GetCiudad_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/lidom.proto",
}
