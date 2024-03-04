// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.12.4
// source: order.proto

package grpcmodel

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
	Order_CreateOrder_FullMethodName     = "/order.Order/CreateOrder"
	Order_UpdateOrder_FullMethodName     = "/order.Order/UpdateOrder"
	Order_DeleteOrder_FullMethodName     = "/order.Order/DeleteOrder"
	Order_GetOrderByID_FullMethodName    = "/order.Order/GetOrderByID"
	Order_GetOrderByParam_FullMethodName = "/order.Order/GetOrderByParam"
	Order_CreateCar_FullMethodName       = "/order.Order/CreateCar"
	Order_UpdateCar_FullMethodName       = "/order.Order/UpdateCar"
	Order_DeleteCar_FullMethodName       = "/order.Order/DeleteCar"
	Order_GetCarByID_FullMethodName      = "/order.Order/GetCarByID"
	Order_GetCarByParam_FullMethodName   = "/order.Order/GetCarByParam"
)

// OrderClient is the client API for Order service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type OrderClient interface {
	CreateOrder(ctx context.Context, in *CreateOrderRequest, opts ...grpc.CallOption) (*SingleOrderReply, error)
	UpdateOrder(ctx context.Context, in *UpdateOrderRequest, opts ...grpc.CallOption) (*SingleOrderReply, error)
	DeleteOrder(ctx context.Context, in *DeleteOrderRequest, opts ...grpc.CallOption) (*DeleteOrderReply, error)
	GetOrderByID(ctx context.Context, in *GetOrderByIDRequest, opts ...grpc.CallOption) (*SingleOrderReply, error)
	GetOrderByParam(ctx context.Context, in *GetOrderByParamRequest, opts ...grpc.CallOption) (*GetOrderByParamReply, error)
	CreateCar(ctx context.Context, in *CreateCarRequest, opts ...grpc.CallOption) (*SingleCarReply, error)
	UpdateCar(ctx context.Context, in *UpdateCarRequest, opts ...grpc.CallOption) (*SingleCarReply, error)
	DeleteCar(ctx context.Context, in *DeleteCarRequest, opts ...grpc.CallOption) (*DeleteCarReply, error)
	GetCarByID(ctx context.Context, in *GetCarByIDRequest, opts ...grpc.CallOption) (*SingleCarReply, error)
	GetCarByParam(ctx context.Context, in *GetCarByParamRequest, opts ...grpc.CallOption) (*GetCarByParamReply, error)
}

type orderClient struct {
	cc grpc.ClientConnInterface
}

func NewOrderClient(cc grpc.ClientConnInterface) OrderClient {
	return &orderClient{cc}
}

func (c *orderClient) CreateOrder(ctx context.Context, in *CreateOrderRequest, opts ...grpc.CallOption) (*SingleOrderReply, error) {
	out := new(SingleOrderReply)
	err := c.cc.Invoke(ctx, Order_CreateOrder_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *orderClient) UpdateOrder(ctx context.Context, in *UpdateOrderRequest, opts ...grpc.CallOption) (*SingleOrderReply, error) {
	out := new(SingleOrderReply)
	err := c.cc.Invoke(ctx, Order_UpdateOrder_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *orderClient) DeleteOrder(ctx context.Context, in *DeleteOrderRequest, opts ...grpc.CallOption) (*DeleteOrderReply, error) {
	out := new(DeleteOrderReply)
	err := c.cc.Invoke(ctx, Order_DeleteOrder_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *orderClient) GetOrderByID(ctx context.Context, in *GetOrderByIDRequest, opts ...grpc.CallOption) (*SingleOrderReply, error) {
	out := new(SingleOrderReply)
	err := c.cc.Invoke(ctx, Order_GetOrderByID_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *orderClient) GetOrderByParam(ctx context.Context, in *GetOrderByParamRequest, opts ...grpc.CallOption) (*GetOrderByParamReply, error) {
	out := new(GetOrderByParamReply)
	err := c.cc.Invoke(ctx, Order_GetOrderByParam_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *orderClient) CreateCar(ctx context.Context, in *CreateCarRequest, opts ...grpc.CallOption) (*SingleCarReply, error) {
	out := new(SingleCarReply)
	err := c.cc.Invoke(ctx, Order_CreateCar_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *orderClient) UpdateCar(ctx context.Context, in *UpdateCarRequest, opts ...grpc.CallOption) (*SingleCarReply, error) {
	out := new(SingleCarReply)
	err := c.cc.Invoke(ctx, Order_UpdateCar_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *orderClient) DeleteCar(ctx context.Context, in *DeleteCarRequest, opts ...grpc.CallOption) (*DeleteCarReply, error) {
	out := new(DeleteCarReply)
	err := c.cc.Invoke(ctx, Order_DeleteCar_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *orderClient) GetCarByID(ctx context.Context, in *GetCarByIDRequest, opts ...grpc.CallOption) (*SingleCarReply, error) {
	out := new(SingleCarReply)
	err := c.cc.Invoke(ctx, Order_GetCarByID_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *orderClient) GetCarByParam(ctx context.Context, in *GetCarByParamRequest, opts ...grpc.CallOption) (*GetCarByParamReply, error) {
	out := new(GetCarByParamReply)
	err := c.cc.Invoke(ctx, Order_GetCarByParam_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// OrderServer is the server API for Order service.
// All implementations must embed UnimplementedOrderServer
// for forward compatibility
type OrderServer interface {
	CreateOrder(context.Context, *CreateOrderRequest) (*SingleOrderReply, error)
	UpdateOrder(context.Context, *UpdateOrderRequest) (*SingleOrderReply, error)
	DeleteOrder(context.Context, *DeleteOrderRequest) (*DeleteOrderReply, error)
	GetOrderByID(context.Context, *GetOrderByIDRequest) (*SingleOrderReply, error)
	GetOrderByParam(context.Context, *GetOrderByParamRequest) (*GetOrderByParamReply, error)
	CreateCar(context.Context, *CreateCarRequest) (*SingleCarReply, error)
	UpdateCar(context.Context, *UpdateCarRequest) (*SingleCarReply, error)
	DeleteCar(context.Context, *DeleteCarRequest) (*DeleteCarReply, error)
	GetCarByID(context.Context, *GetCarByIDRequest) (*SingleCarReply, error)
	GetCarByParam(context.Context, *GetCarByParamRequest) (*GetCarByParamReply, error)
	mustEmbedUnimplementedOrderServer()
}

// UnimplementedOrderServer must be embedded to have forward compatible implementations.
type UnimplementedOrderServer struct {
}

func (UnimplementedOrderServer) CreateOrder(context.Context, *CreateOrderRequest) (*SingleOrderReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateOrder not implemented")
}
func (UnimplementedOrderServer) UpdateOrder(context.Context, *UpdateOrderRequest) (*SingleOrderReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateOrder not implemented")
}
func (UnimplementedOrderServer) DeleteOrder(context.Context, *DeleteOrderRequest) (*DeleteOrderReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteOrder not implemented")
}
func (UnimplementedOrderServer) GetOrderByID(context.Context, *GetOrderByIDRequest) (*SingleOrderReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetOrderByID not implemented")
}
func (UnimplementedOrderServer) GetOrderByParam(context.Context, *GetOrderByParamRequest) (*GetOrderByParamReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetOrderByParam not implemented")
}
func (UnimplementedOrderServer) CreateCar(context.Context, *CreateCarRequest) (*SingleCarReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateCar not implemented")
}
func (UnimplementedOrderServer) UpdateCar(context.Context, *UpdateCarRequest) (*SingleCarReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateCar not implemented")
}
func (UnimplementedOrderServer) DeleteCar(context.Context, *DeleteCarRequest) (*DeleteCarReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteCar not implemented")
}
func (UnimplementedOrderServer) GetCarByID(context.Context, *GetCarByIDRequest) (*SingleCarReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCarByID not implemented")
}
func (UnimplementedOrderServer) GetCarByParam(context.Context, *GetCarByParamRequest) (*GetCarByParamReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCarByParam not implemented")
}
func (UnimplementedOrderServer) mustEmbedUnimplementedOrderServer() {}

// UnsafeOrderServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to OrderServer will
// result in compilation errors.
type UnsafeOrderServer interface {
	mustEmbedUnimplementedOrderServer()
}

func RegisterOrderServer(s grpc.ServiceRegistrar, srv OrderServer) {
	s.RegisterService(&Order_ServiceDesc, srv)
}

func _Order_CreateOrder_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateOrderRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrderServer).CreateOrder(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Order_CreateOrder_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrderServer).CreateOrder(ctx, req.(*CreateOrderRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Order_UpdateOrder_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateOrderRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrderServer).UpdateOrder(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Order_UpdateOrder_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrderServer).UpdateOrder(ctx, req.(*UpdateOrderRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Order_DeleteOrder_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteOrderRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrderServer).DeleteOrder(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Order_DeleteOrder_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrderServer).DeleteOrder(ctx, req.(*DeleteOrderRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Order_GetOrderByID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetOrderByIDRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrderServer).GetOrderByID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Order_GetOrderByID_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrderServer).GetOrderByID(ctx, req.(*GetOrderByIDRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Order_GetOrderByParam_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetOrderByParamRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrderServer).GetOrderByParam(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Order_GetOrderByParam_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrderServer).GetOrderByParam(ctx, req.(*GetOrderByParamRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Order_CreateCar_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateCarRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrderServer).CreateCar(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Order_CreateCar_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrderServer).CreateCar(ctx, req.(*CreateCarRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Order_UpdateCar_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateCarRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrderServer).UpdateCar(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Order_UpdateCar_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrderServer).UpdateCar(ctx, req.(*UpdateCarRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Order_DeleteCar_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteCarRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrderServer).DeleteCar(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Order_DeleteCar_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrderServer).DeleteCar(ctx, req.(*DeleteCarRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Order_GetCarByID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetCarByIDRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrderServer).GetCarByID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Order_GetCarByID_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrderServer).GetCarByID(ctx, req.(*GetCarByIDRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Order_GetCarByParam_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetCarByParamRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrderServer).GetCarByParam(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Order_GetCarByParam_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrderServer).GetCarByParam(ctx, req.(*GetCarByParamRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Order_ServiceDesc is the grpc.ServiceDesc for Order service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Order_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "order.Order",
	HandlerType: (*OrderServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateOrder",
			Handler:    _Order_CreateOrder_Handler,
		},
		{
			MethodName: "UpdateOrder",
			Handler:    _Order_UpdateOrder_Handler,
		},
		{
			MethodName: "DeleteOrder",
			Handler:    _Order_DeleteOrder_Handler,
		},
		{
			MethodName: "GetOrderByID",
			Handler:    _Order_GetOrderByID_Handler,
		},
		{
			MethodName: "GetOrderByParam",
			Handler:    _Order_GetOrderByParam_Handler,
		},
		{
			MethodName: "CreateCar",
			Handler:    _Order_CreateCar_Handler,
		},
		{
			MethodName: "UpdateCar",
			Handler:    _Order_UpdateCar_Handler,
		},
		{
			MethodName: "DeleteCar",
			Handler:    _Order_DeleteCar_Handler,
		},
		{
			MethodName: "GetCarByID",
			Handler:    _Order_GetCarByID_Handler,
		},
		{
			MethodName: "GetCarByParam",
			Handler:    _Order_GetCarByParam_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "order.proto",
}
