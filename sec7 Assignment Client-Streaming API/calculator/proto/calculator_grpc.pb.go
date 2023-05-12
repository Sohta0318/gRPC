// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.12
// source: calculator.proto

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

// CalculatorClient is the client API for Calculator service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CalculatorClient interface {
	Calculate(ctx context.Context, opts ...grpc.CallOption) (Calculator_CalculateClient, error)
}

type calculatorClient struct {
	cc grpc.ClientConnInterface
}

func NewCalculatorClient(cc grpc.ClientConnInterface) CalculatorClient {
	return &calculatorClient{cc}
}

func (c *calculatorClient) Calculate(ctx context.Context, opts ...grpc.CallOption) (Calculator_CalculateClient, error) {
	stream, err := c.cc.NewStream(ctx, &Calculator_ServiceDesc.Streams[0], "/Calculator/Calculate", opts...)
	if err != nil {
		return nil, err
	}
	x := &calculatorCalculateClient{stream}
	return x, nil
}

type Calculator_CalculateClient interface {
	Send(*CalculatorRequest) error
	CloseAndRecv() (*CalculatorResponse, error)
	grpc.ClientStream
}

type calculatorCalculateClient struct {
	grpc.ClientStream
}

func (x *calculatorCalculateClient) Send(m *CalculatorRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *calculatorCalculateClient) CloseAndRecv() (*CalculatorResponse, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(CalculatorResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// CalculatorServer is the server API for Calculator service.
// All implementations must embed UnimplementedCalculatorServer
// for forward compatibility
type CalculatorServer interface {
	Calculate(Calculator_CalculateServer) error
	mustEmbedUnimplementedCalculatorServer()
}

// UnimplementedCalculatorServer must be embedded to have forward compatible implementations.
type UnimplementedCalculatorServer struct {
}

func (UnimplementedCalculatorServer) Calculate(Calculator_CalculateServer) error {
	return status.Errorf(codes.Unimplemented, "method Calculate not implemented")
}
func (UnimplementedCalculatorServer) mustEmbedUnimplementedCalculatorServer() {}

// UnsafeCalculatorServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CalculatorServer will
// result in compilation errors.
type UnsafeCalculatorServer interface {
	mustEmbedUnimplementedCalculatorServer()
}

func RegisterCalculatorServer(s grpc.ServiceRegistrar, srv CalculatorServer) {
	s.RegisterService(&Calculator_ServiceDesc, srv)
}

func _Calculator_Calculate_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(CalculatorServer).Calculate(&calculatorCalculateServer{stream})
}

type Calculator_CalculateServer interface {
	SendAndClose(*CalculatorResponse) error
	Recv() (*CalculatorRequest, error)
	grpc.ServerStream
}

type calculatorCalculateServer struct {
	grpc.ServerStream
}

func (x *calculatorCalculateServer) SendAndClose(m *CalculatorResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *calculatorCalculateServer) Recv() (*CalculatorRequest, error) {
	m := new(CalculatorRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// Calculator_ServiceDesc is the grpc.ServiceDesc for Calculator service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Calculator_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "Calculator",
	HandlerType: (*CalculatorServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Calculate",
			Handler:       _Calculator_Calculate_Handler,
			ClientStreams: true,
		},
	},
	Metadata: "calculator.proto",
}
