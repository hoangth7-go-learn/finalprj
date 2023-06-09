// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.22.4
// source: proto/flights.proto

package flights

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
	FlightService_UpsertFlight_FullMethodName       = "/flight.FlightService/UpsertFlight"
	FlightService_GetFlightsList_FullMethodName     = "/flight.FlightService/GetFlightsList"
	FlightService_GetFlightById_FullMethodName      = "/flight.FlightService/GetFlightById"
	FlightService_ChangeFlightStatus_FullMethodName = "/flight.FlightService/ChangeFlightStatus"
	FlightService_BookFlight_FullMethodName         = "/flight.FlightService/BookFlight"
)

// FlightServiceClient is the client API for FlightService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type FlightServiceClient interface {
	UpsertFlight(ctx context.Context, in *Flight, opts ...grpc.CallOption) (*FlightId, error)
	GetFlightsList(ctx context.Context, in *FlightQuery, opts ...grpc.CallOption) (*FlightList, error)
	GetFlightById(ctx context.Context, in *FlightId, opts ...grpc.CallOption) (*Flight, error)
	ChangeFlightStatus(ctx context.Context, in *FlightStatusRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	BookFlight(ctx context.Context, in *BookFlightRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
}

type flightServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewFlightServiceClient(cc grpc.ClientConnInterface) FlightServiceClient {
	return &flightServiceClient{cc}
}

func (c *flightServiceClient) UpsertFlight(ctx context.Context, in *Flight, opts ...grpc.CallOption) (*FlightId, error) {
	out := new(FlightId)
	err := c.cc.Invoke(ctx, FlightService_UpsertFlight_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *flightServiceClient) GetFlightsList(ctx context.Context, in *FlightQuery, opts ...grpc.CallOption) (*FlightList, error) {
	out := new(FlightList)
	err := c.cc.Invoke(ctx, FlightService_GetFlightsList_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *flightServiceClient) GetFlightById(ctx context.Context, in *FlightId, opts ...grpc.CallOption) (*Flight, error) {
	out := new(Flight)
	err := c.cc.Invoke(ctx, FlightService_GetFlightById_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *flightServiceClient) ChangeFlightStatus(ctx context.Context, in *FlightStatusRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, FlightService_ChangeFlightStatus_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *flightServiceClient) BookFlight(ctx context.Context, in *BookFlightRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, FlightService_BookFlight_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// FlightServiceServer is the server API for FlightService service.
// All implementations must embed UnimplementedFlightServiceServer
// for forward compatibility
type FlightServiceServer interface {
	UpsertFlight(context.Context, *Flight) (*FlightId, error)
	GetFlightsList(context.Context, *FlightQuery) (*FlightList, error)
	GetFlightById(context.Context, *FlightId) (*Flight, error)
	ChangeFlightStatus(context.Context, *FlightStatusRequest) (*emptypb.Empty, error)
	BookFlight(context.Context, *BookFlightRequest) (*emptypb.Empty, error)
	mustEmbedUnimplementedFlightServiceServer()
}

// UnimplementedFlightServiceServer must be embedded to have forward compatible implementations.
type UnimplementedFlightServiceServer struct {
}

func (UnimplementedFlightServiceServer) UpsertFlight(context.Context, *Flight) (*FlightId, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpsertFlight not implemented")
}
func (UnimplementedFlightServiceServer) GetFlightsList(context.Context, *FlightQuery) (*FlightList, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetFlightsList not implemented")
}
func (UnimplementedFlightServiceServer) GetFlightById(context.Context, *FlightId) (*Flight, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetFlightById not implemented")
}
func (UnimplementedFlightServiceServer) ChangeFlightStatus(context.Context, *FlightStatusRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ChangeFlightStatus not implemented")
}
func (UnimplementedFlightServiceServer) BookFlight(context.Context, *BookFlightRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method BookFlight not implemented")
}
func (UnimplementedFlightServiceServer) mustEmbedUnimplementedFlightServiceServer() {}

// UnsafeFlightServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to FlightServiceServer will
// result in compilation errors.
type UnsafeFlightServiceServer interface {
	mustEmbedUnimplementedFlightServiceServer()
}

func RegisterFlightServiceServer(s grpc.ServiceRegistrar, srv FlightServiceServer) {
	s.RegisterService(&FlightService_ServiceDesc, srv)
}

func _FlightService_UpsertFlight_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Flight)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FlightServiceServer).UpsertFlight(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: FlightService_UpsertFlight_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FlightServiceServer).UpsertFlight(ctx, req.(*Flight))
	}
	return interceptor(ctx, in, info, handler)
}

func _FlightService_GetFlightsList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FlightQuery)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FlightServiceServer).GetFlightsList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: FlightService_GetFlightsList_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FlightServiceServer).GetFlightsList(ctx, req.(*FlightQuery))
	}
	return interceptor(ctx, in, info, handler)
}

func _FlightService_GetFlightById_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FlightId)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FlightServiceServer).GetFlightById(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: FlightService_GetFlightById_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FlightServiceServer).GetFlightById(ctx, req.(*FlightId))
	}
	return interceptor(ctx, in, info, handler)
}

func _FlightService_ChangeFlightStatus_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FlightStatusRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FlightServiceServer).ChangeFlightStatus(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: FlightService_ChangeFlightStatus_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FlightServiceServer).ChangeFlightStatus(ctx, req.(*FlightStatusRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _FlightService_BookFlight_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BookFlightRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FlightServiceServer).BookFlight(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: FlightService_BookFlight_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FlightServiceServer).BookFlight(ctx, req.(*BookFlightRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// FlightService_ServiceDesc is the grpc.ServiceDesc for FlightService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var FlightService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "flight.FlightService",
	HandlerType: (*FlightServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "UpsertFlight",
			Handler:    _FlightService_UpsertFlight_Handler,
		},
		{
			MethodName: "GetFlightsList",
			Handler:    _FlightService_GetFlightsList_Handler,
		},
		{
			MethodName: "GetFlightById",
			Handler:    _FlightService_GetFlightById_Handler,
		},
		{
			MethodName: "ChangeFlightStatus",
			Handler:    _FlightService_ChangeFlightStatus_Handler,
		},
		{
			MethodName: "BookFlight",
			Handler:    _FlightService_BookFlight_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/flights.proto",
}
