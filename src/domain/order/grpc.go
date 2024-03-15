package order

import (
	"context"
	"time"

	"github.com/achwanyusuf/carrent-lib/pkg/errormsg"
	"github.com/achwanyusuf/carrent-ordersvc/src/model/grpcmodel"
	"github.com/achwanyusuf/carrent-ordersvc/src/model/svcerr"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (o *OrderDep) InsertGRPC(ctx context.Context, v *grpcmodel.CreateOrderRequest) (*grpcmodel.SingleOrderReply, error) {
	client, err := o.Grpc.Get()
	if err != nil {
		return nil, errormsg.WrapErr(svcerr.OrderSVCErrorGRPCClient, err, "error grpc client connection")
	}

	clientService := grpcmodel.NewOrderClient(client.Conn)

	ctx, cancel := context.WithTimeout(ctx, time.Second)
	defer cancel()
	res, err := clientService.CreateOrder(ctx, v)
	if err != nil {
		return nil, errormsg.WrapErr(svcerr.OrderSVCErrorGRPCClient, err, "error grpc client")
	}

	client.Release()
	return res, nil
}

func (o *OrderDep) UpdateGRPC(ctx context.Context, v *grpcmodel.UpdateOrderRequest) (*grpcmodel.SingleOrderReply, error) {
	client, err := o.Grpc.Get()
	if err != nil {
		return nil, errormsg.WrapErr(svcerr.OrderSVCErrorGRPCClient, err, "error grpc client connection")
	}

	clientService := grpcmodel.NewOrderClient(client.Conn)

	ctx, cancel := context.WithTimeout(ctx, time.Second)
	defer cancel()
	res, err := clientService.UpdateOrder(ctx, v)
	if err != nil {
		return nil, errormsg.WrapErr(svcerr.OrderSVCErrorGRPCClient, err, "error grpc client")
	}

	client.Release()
	return res, nil
}

func (o *OrderDep) DeleteGRPC(ctx context.Context, v *grpcmodel.DeleteOrderRequest) (*grpcmodel.DeleteOrderReply, error) {
	client, err := o.Grpc.Get()
	if err != nil {
		return nil, errormsg.WrapErr(svcerr.OrderSVCErrorGRPCClient, err, "error grpc client connection")
	}

	clientService := grpcmodel.NewOrderClient(client.Conn)

	ctx, cancel := context.WithTimeout(ctx, time.Second)
	defer cancel()
	res, err := clientService.DeleteOrder(ctx, v)
	if err != nil {
		return nil, errormsg.WrapErr(svcerr.OrderSVCErrorGRPCClient, err, "error grpc client")
	}

	client.Release()
	return res, nil
}

func (o *OrderDep) GetByIDGRPC(ctx context.Context, v *grpcmodel.GetOrderByIDRequest) (*grpcmodel.SingleOrderReply, error) {
	client, err := o.Grpc.Get()
	if err != nil {
		return nil, errormsg.WrapErr(svcerr.OrderSVCErrorGRPCClient, err, "error grpc client connection")
	}

	clientService := grpcmodel.NewOrderClient(client.Conn)

	ctx, cancel := context.WithTimeout(ctx, time.Second)
	defer cancel()
	res, err := clientService.GetOrderByID(ctx, v)
	if err != nil {
		if st, ok := status.FromError(err); ok {
			if st.Code() == codes.NotFound {
				return nil, errormsg.WrapErr(svcerr.OrderSVCNotFound, err, "data not found")
			}
		}
		return nil, errormsg.WrapErr(svcerr.OrderSVCErrorGRPCClient, err, "error grpc client")
	}

	client.Release()
	return res, nil
}

func (o *OrderDep) GetOrderByParam(ctx context.Context, v *grpcmodel.GetOrderByParamRequest) (*grpcmodel.GetOrderByParamReply, error) {
	client, err := o.Grpc.Get()
	if err != nil {
		return nil, errormsg.WrapErr(svcerr.OrderSVCErrorGRPCClient, err, "error grpc client connection")
	}

	clientService := grpcmodel.NewOrderClient(client.Conn)

	ctx, cancel := context.WithTimeout(ctx, time.Second)
	defer cancel()
	res, err := clientService.GetOrderByParam(ctx, v)
	if err != nil {
		return nil, errormsg.WrapErr(svcerr.OrderSVCErrorGRPCClient, err, "error grpc client")
	}

	client.Release()
	return res, nil
}
