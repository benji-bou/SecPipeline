// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package pluginctl

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

// SecPipelinePluginsClient is the client API for SecPipelinePlugins service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type SecPipelinePluginsClient interface {
	GetInputSchema(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*InputSchema, error)
	Config(ctx context.Context, in *RunInputConfig, opts ...grpc.CallOption) (*Empty, error)
	Input(ctx context.Context, opts ...grpc.CallOption) (SecPipelinePlugins_InputClient, error)
	Output(ctx context.Context, in *Empty, opts ...grpc.CallOption) (SecPipelinePlugins_OutputClient, error)
}

type secPipelinePluginsClient struct {
	cc grpc.ClientConnInterface
}

func NewSecPipelinePluginsClient(cc grpc.ClientConnInterface) SecPipelinePluginsClient {
	return &secPipelinePluginsClient{cc}
}

func (c *secPipelinePluginsClient) GetInputSchema(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*InputSchema, error) {
	out := new(InputSchema)
	err := c.cc.Invoke(ctx, "/pluginctl.SecPipelinePlugins/GetInputSchema", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *secPipelinePluginsClient) Config(ctx context.Context, in *RunInputConfig, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, "/pluginctl.SecPipelinePlugins/Config", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *secPipelinePluginsClient) Input(ctx context.Context, opts ...grpc.CallOption) (SecPipelinePlugins_InputClient, error) {
	stream, err := c.cc.NewStream(ctx, &SecPipelinePlugins_ServiceDesc.Streams[0], "/pluginctl.SecPipelinePlugins/Input", opts...)
	if err != nil {
		return nil, err
	}
	x := &secPipelinePluginsInputClient{stream}
	return x, nil
}

type SecPipelinePlugins_InputClient interface {
	Send(*DataStream) error
	CloseAndRecv() (*Empty, error)
	grpc.ClientStream
}

type secPipelinePluginsInputClient struct {
	grpc.ClientStream
}

func (x *secPipelinePluginsInputClient) Send(m *DataStream) error {
	return x.ClientStream.SendMsg(m)
}

func (x *secPipelinePluginsInputClient) CloseAndRecv() (*Empty, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(Empty)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *secPipelinePluginsClient) Output(ctx context.Context, in *Empty, opts ...grpc.CallOption) (SecPipelinePlugins_OutputClient, error) {
	stream, err := c.cc.NewStream(ctx, &SecPipelinePlugins_ServiceDesc.Streams[1], "/pluginctl.SecPipelinePlugins/Output", opts...)
	if err != nil {
		return nil, err
	}
	x := &secPipelinePluginsOutputClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type SecPipelinePlugins_OutputClient interface {
	Recv() (*DataStream, error)
	grpc.ClientStream
}

type secPipelinePluginsOutputClient struct {
	grpc.ClientStream
}

func (x *secPipelinePluginsOutputClient) Recv() (*DataStream, error) {
	m := new(DataStream)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// SecPipelinePluginsServer is the server API for SecPipelinePlugins service.
// All implementations must embed UnimplementedSecPipelinePluginsServer
// for forward compatibility
type SecPipelinePluginsServer interface {
	GetInputSchema(context.Context, *Empty) (*InputSchema, error)
	Config(context.Context, *RunInputConfig) (*Empty, error)
	Input(SecPipelinePlugins_InputServer) error
	Output(*Empty, SecPipelinePlugins_OutputServer) error
	mustEmbedUnimplementedSecPipelinePluginsServer()
}

// UnimplementedSecPipelinePluginsServer must be embedded to have forward compatible implementations.
type UnimplementedSecPipelinePluginsServer struct {
}

func (UnimplementedSecPipelinePluginsServer) GetInputSchema(context.Context, *Empty) (*InputSchema, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetInputSchema not implemented")
}
func (UnimplementedSecPipelinePluginsServer) Config(context.Context, *RunInputConfig) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Config not implemented")
}
func (UnimplementedSecPipelinePluginsServer) Input(SecPipelinePlugins_InputServer) error {
	return status.Errorf(codes.Unimplemented, "method Input not implemented")
}
func (UnimplementedSecPipelinePluginsServer) Output(*Empty, SecPipelinePlugins_OutputServer) error {
	return status.Errorf(codes.Unimplemented, "method Output not implemented")
}
func (UnimplementedSecPipelinePluginsServer) mustEmbedUnimplementedSecPipelinePluginsServer() {}

// UnsafeSecPipelinePluginsServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to SecPipelinePluginsServer will
// result in compilation errors.
type UnsafeSecPipelinePluginsServer interface {
	mustEmbedUnimplementedSecPipelinePluginsServer()
}

func RegisterSecPipelinePluginsServer(s grpc.ServiceRegistrar, srv SecPipelinePluginsServer) {
	s.RegisterService(&SecPipelinePlugins_ServiceDesc, srv)
}

func _SecPipelinePlugins_GetInputSchema_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SecPipelinePluginsServer).GetInputSchema(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pluginctl.SecPipelinePlugins/GetInputSchema",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SecPipelinePluginsServer).GetInputSchema(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _SecPipelinePlugins_Config_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RunInputConfig)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SecPipelinePluginsServer).Config(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pluginctl.SecPipelinePlugins/Config",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SecPipelinePluginsServer).Config(ctx, req.(*RunInputConfig))
	}
	return interceptor(ctx, in, info, handler)
}

func _SecPipelinePlugins_Input_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(SecPipelinePluginsServer).Input(&secPipelinePluginsInputServer{stream})
}

type SecPipelinePlugins_InputServer interface {
	SendAndClose(*Empty) error
	Recv() (*DataStream, error)
	grpc.ServerStream
}

type secPipelinePluginsInputServer struct {
	grpc.ServerStream
}

func (x *secPipelinePluginsInputServer) SendAndClose(m *Empty) error {
	return x.ServerStream.SendMsg(m)
}

func (x *secPipelinePluginsInputServer) Recv() (*DataStream, error) {
	m := new(DataStream)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _SecPipelinePlugins_Output_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(Empty)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(SecPipelinePluginsServer).Output(m, &secPipelinePluginsOutputServer{stream})
}

type SecPipelinePlugins_OutputServer interface {
	Send(*DataStream) error
	grpc.ServerStream
}

type secPipelinePluginsOutputServer struct {
	grpc.ServerStream
}

func (x *secPipelinePluginsOutputServer) Send(m *DataStream) error {
	return x.ServerStream.SendMsg(m)
}

// SecPipelinePlugins_ServiceDesc is the grpc.ServiceDesc for SecPipelinePlugins service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var SecPipelinePlugins_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "pluginctl.SecPipelinePlugins",
	HandlerType: (*SecPipelinePluginsServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetInputSchema",
			Handler:    _SecPipelinePlugins_GetInputSchema_Handler,
		},
		{
			MethodName: "Config",
			Handler:    _SecPipelinePlugins_Config_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Input",
			Handler:       _SecPipelinePlugins_Input_Handler,
			ClientStreams: true,
		},
		{
			StreamName:    "Output",
			Handler:       _SecPipelinePlugins_Output_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "pluginctl/plugins.proto",
}
