// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v4.23.4
// source: api/v1beta1/tracee.proto

package v1beta1

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

// TraceeServiceClient is the client API for TraceeService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type TraceeServiceClient interface {
	GetEventDefinition(ctx context.Context, in *GetEventDefinitionRequest, opts ...grpc.CallOption) (*GetEventDefinitionResponse, error)
	StreamEvents(ctx context.Context, in *StreamEventsRequest, opts ...grpc.CallOption) (TraceeService_StreamEventsClient, error)
	EnableEvent(ctx context.Context, in *EnableEventRequest, opts ...grpc.CallOption) (*EnableEventResponse, error)
	DisableEvent(ctx context.Context, in *DisableEventRequest, opts ...grpc.CallOption) (*DisableEventResponse, error)
	GetVersion(ctx context.Context, in *GetVersionRequest, opts ...grpc.CallOption) (*GetVersionResponse, error)
}

type traceeServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewTraceeServiceClient(cc grpc.ClientConnInterface) TraceeServiceClient {
	return &traceeServiceClient{cc}
}

func (c *traceeServiceClient) GetEventDefinition(ctx context.Context, in *GetEventDefinitionRequest, opts ...grpc.CallOption) (*GetEventDefinitionResponse, error) {
	out := new(GetEventDefinitionResponse)
	err := c.cc.Invoke(ctx, "/tracee.v1beta1.TraceeService/GetEventDefinition", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *traceeServiceClient) StreamEvents(ctx context.Context, in *StreamEventsRequest, opts ...grpc.CallOption) (TraceeService_StreamEventsClient, error) {
	stream, err := c.cc.NewStream(ctx, &TraceeService_ServiceDesc.Streams[0], "/tracee.v1beta1.TraceeService/StreamEvents", opts...)
	if err != nil {
		return nil, err
	}
	x := &traceeServiceStreamEventsClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type TraceeService_StreamEventsClient interface {
	Recv() (*StreamEventsResponse, error)
	grpc.ClientStream
}

type traceeServiceStreamEventsClient struct {
	grpc.ClientStream
}

func (x *traceeServiceStreamEventsClient) Recv() (*StreamEventsResponse, error) {
	m := new(StreamEventsResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *traceeServiceClient) EnableEvent(ctx context.Context, in *EnableEventRequest, opts ...grpc.CallOption) (*EnableEventResponse, error) {
	out := new(EnableEventResponse)
	err := c.cc.Invoke(ctx, "/tracee.v1beta1.TraceeService/EnableEvent", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *traceeServiceClient) DisableEvent(ctx context.Context, in *DisableEventRequest, opts ...grpc.CallOption) (*DisableEventResponse, error) {
	out := new(DisableEventResponse)
	err := c.cc.Invoke(ctx, "/tracee.v1beta1.TraceeService/DisableEvent", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *traceeServiceClient) GetVersion(ctx context.Context, in *GetVersionRequest, opts ...grpc.CallOption) (*GetVersionResponse, error) {
	out := new(GetVersionResponse)
	err := c.cc.Invoke(ctx, "/tracee.v1beta1.TraceeService/GetVersion", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TraceeServiceServer is the server API for TraceeService service.
// All implementations must embed UnimplementedTraceeServiceServer
// for forward compatibility
type TraceeServiceServer interface {
	GetEventDefinition(context.Context, *GetEventDefinitionRequest) (*GetEventDefinitionResponse, error)
	StreamEvents(*StreamEventsRequest, TraceeService_StreamEventsServer) error
	EnableEvent(context.Context, *EnableEventRequest) (*EnableEventResponse, error)
	DisableEvent(context.Context, *DisableEventRequest) (*DisableEventResponse, error)
	GetVersion(context.Context, *GetVersionRequest) (*GetVersionResponse, error)
	mustEmbedUnimplementedTraceeServiceServer()
}

// UnimplementedTraceeServiceServer must be embedded to have forward compatible implementations.
type UnimplementedTraceeServiceServer struct {
}

func (UnimplementedTraceeServiceServer) GetEventDefinition(context.Context, *GetEventDefinitionRequest) (*GetEventDefinitionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetEventDefinition not implemented")
}
func (UnimplementedTraceeServiceServer) StreamEvents(*StreamEventsRequest, TraceeService_StreamEventsServer) error {
	return status.Errorf(codes.Unimplemented, "method StreamEvents not implemented")
}
func (UnimplementedTraceeServiceServer) EnableEvent(context.Context, *EnableEventRequest) (*EnableEventResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method EnableEvent not implemented")
}
func (UnimplementedTraceeServiceServer) DisableEvent(context.Context, *DisableEventRequest) (*DisableEventResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DisableEvent not implemented")
}
func (UnimplementedTraceeServiceServer) GetVersion(context.Context, *GetVersionRequest) (*GetVersionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetVersion not implemented")
}
func (UnimplementedTraceeServiceServer) mustEmbedUnimplementedTraceeServiceServer() {}

// UnsafeTraceeServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to TraceeServiceServer will
// result in compilation errors.
type UnsafeTraceeServiceServer interface {
	mustEmbedUnimplementedTraceeServiceServer()
}

func RegisterTraceeServiceServer(s grpc.ServiceRegistrar, srv TraceeServiceServer) {
	s.RegisterService(&TraceeService_ServiceDesc, srv)
}

func _TraceeService_GetEventDefinition_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetEventDefinitionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TraceeServiceServer).GetEventDefinition(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/tracee.v1beta1.TraceeService/GetEventDefinition",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TraceeServiceServer).GetEventDefinition(ctx, req.(*GetEventDefinitionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TraceeService_StreamEvents_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(StreamEventsRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(TraceeServiceServer).StreamEvents(m, &traceeServiceStreamEventsServer{stream})
}

type TraceeService_StreamEventsServer interface {
	Send(*StreamEventsResponse) error
	grpc.ServerStream
}

type traceeServiceStreamEventsServer struct {
	grpc.ServerStream
}

func (x *traceeServiceStreamEventsServer) Send(m *StreamEventsResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _TraceeService_EnableEvent_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EnableEventRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TraceeServiceServer).EnableEvent(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/tracee.v1beta1.TraceeService/EnableEvent",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TraceeServiceServer).EnableEvent(ctx, req.(*EnableEventRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TraceeService_DisableEvent_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DisableEventRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TraceeServiceServer).DisableEvent(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/tracee.v1beta1.TraceeService/DisableEvent",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TraceeServiceServer).DisableEvent(ctx, req.(*DisableEventRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TraceeService_GetVersion_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetVersionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TraceeServiceServer).GetVersion(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/tracee.v1beta1.TraceeService/GetVersion",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TraceeServiceServer).GetVersion(ctx, req.(*GetVersionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// TraceeService_ServiceDesc is the grpc.ServiceDesc for TraceeService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var TraceeService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "tracee.v1beta1.TraceeService",
	HandlerType: (*TraceeServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetEventDefinition",
			Handler:    _TraceeService_GetEventDefinition_Handler,
		},
		{
			MethodName: "EnableEvent",
			Handler:    _TraceeService_EnableEvent_Handler,
		},
		{
			MethodName: "DisableEvent",
			Handler:    _TraceeService_DisableEvent_Handler,
		},
		{
			MethodName: "GetVersion",
			Handler:    _TraceeService_GetVersion_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "StreamEvents",
			Handler:       _TraceeService_StreamEvents_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "api/v1beta1/tracee.proto",
}
