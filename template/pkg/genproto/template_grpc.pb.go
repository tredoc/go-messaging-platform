// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v4.25.1
// source: template.proto

package template

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

// TemplateClient is the client API for Template service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type TemplateClient interface {
	CreateTemplate(ctx context.Context, in *CreateTemplateRequest, opts ...grpc.CallOption) (*CreateTemplateResponse, error)
	GetTemplateByID(ctx context.Context, in *GetTemplateByIDRequest, opts ...grpc.CallOption) (*GetTemplateByIDResponse, error)
	DeleteTemplateByID(ctx context.Context, in *DeleteTemplateByIDRequest, opts ...grpc.CallOption) (*DeleteTemplateByIDResponse, error)
}

type templateClient struct {
	cc grpc.ClientConnInterface
}

func NewTemplateClient(cc grpc.ClientConnInterface) TemplateClient {
	return &templateClient{cc}
}

func (c *templateClient) CreateTemplate(ctx context.Context, in *CreateTemplateRequest, opts ...grpc.CallOption) (*CreateTemplateResponse, error) {
	out := new(CreateTemplateResponse)
	err := c.cc.Invoke(ctx, "/template.Template/CreateTemplate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *templateClient) GetTemplateByID(ctx context.Context, in *GetTemplateByIDRequest, opts ...grpc.CallOption) (*GetTemplateByIDResponse, error) {
	out := new(GetTemplateByIDResponse)
	err := c.cc.Invoke(ctx, "/template.Template/GetTemplateByID", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *templateClient) DeleteTemplateByID(ctx context.Context, in *DeleteTemplateByIDRequest, opts ...grpc.CallOption) (*DeleteTemplateByIDResponse, error) {
	out := new(DeleteTemplateByIDResponse)
	err := c.cc.Invoke(ctx, "/template.Template/DeleteTemplateByID", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TemplateServer is the server API for Template service.
// All implementations should embed UnimplementedTemplateServer
// for forward compatibility
type TemplateServer interface {
	CreateTemplate(context.Context, *CreateTemplateRequest) (*CreateTemplateResponse, error)
	GetTemplateByID(context.Context, *GetTemplateByIDRequest) (*GetTemplateByIDResponse, error)
	DeleteTemplateByID(context.Context, *DeleteTemplateByIDRequest) (*DeleteTemplateByIDResponse, error)
}

// UnimplementedTemplateServer should be embedded to have forward compatible implementations.
type UnimplementedTemplateServer struct {
}

func (UnimplementedTemplateServer) CreateTemplate(context.Context, *CreateTemplateRequest) (*CreateTemplateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateTemplate not implemented")
}
func (UnimplementedTemplateServer) GetTemplateByID(context.Context, *GetTemplateByIDRequest) (*GetTemplateByIDResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTemplateByID not implemented")
}
func (UnimplementedTemplateServer) DeleteTemplateByID(context.Context, *DeleteTemplateByIDRequest) (*DeleteTemplateByIDResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteTemplateByID not implemented")
}

// UnsafeTemplateServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to TemplateServer will
// result in compilation errors.
type UnsafeTemplateServer interface {
	mustEmbedUnimplementedTemplateServer()
}

func RegisterTemplateServer(s grpc.ServiceRegistrar, srv TemplateServer) {
	s.RegisterService(&Template_ServiceDesc, srv)
}

func _Template_CreateTemplate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateTemplateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TemplateServer).CreateTemplate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/template.Template/CreateTemplate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TemplateServer).CreateTemplate(ctx, req.(*CreateTemplateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Template_GetTemplateByID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetTemplateByIDRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TemplateServer).GetTemplateByID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/template.Template/GetTemplateByID",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TemplateServer).GetTemplateByID(ctx, req.(*GetTemplateByIDRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Template_DeleteTemplateByID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteTemplateByIDRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TemplateServer).DeleteTemplateByID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/template.Template/DeleteTemplateByID",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TemplateServer).DeleteTemplateByID(ctx, req.(*DeleteTemplateByIDRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Template_ServiceDesc is the grpc.ServiceDesc for Template service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Template_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "template.Template",
	HandlerType: (*TemplateServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateTemplate",
			Handler:    _Template_CreateTemplate_Handler,
		},
		{
			MethodName: "GetTemplateByID",
			Handler:    _Template_GetTemplateByID_Handler,
		},
		{
			MethodName: "DeleteTemplateByID",
			Handler:    _Template_DeleteTemplateByID_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "template.proto",
}
