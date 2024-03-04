package car

import (
	"context"
	"time"

	"github.com/achwanyusuf/carrent-lib/pkg/errormsg"
	"github.com/achwanyusuf/carrent-ordersvc/src/model/grpcmodel"
	"github.com/achwanyusuf/carrent-ordersvc/src/model/svcerr"
)

func (c *CarDep) InsertGRPC(ctx context.Context, v *grpcmodel.CreateCarRequest) (*grpcmodel.SingleCarReply, error) {
	client, err := c.Grpc.Get()
	if err != nil {
		return nil, errormsg.WrapErr(svcerr.OrderSVCErrorGRPCClient, err, "error grpc client connection")
	}

	clientService := grpcmodel.NewOrderClient(client.Conn)

	ctx, cancel := context.WithTimeout(ctx, time.Second)
	defer cancel()
	res, err := clientService.CreateCar(ctx, v)
	if err != nil {
		return nil, errormsg.WrapErr(svcerr.OrderSVCErrorGRPCClient, err, "error grpc client")
	}

	client.Release()
	return res, nil
}

func (c *CarDep) UpdateGRPC(ctx context.Context, v *grpcmodel.UpdateCarRequest) (*grpcmodel.SingleCarReply, error) {
	client, err := c.Grpc.Get()
	if err != nil {
		return nil, errormsg.WrapErr(svcerr.OrderSVCErrorGRPCClient, err, "error grpc client connection")
	}

	clientService := grpcmodel.NewOrderClient(client.Conn)

	ctx, cancel := context.WithTimeout(ctx, time.Second)
	defer cancel()
	res, err := clientService.UpdateCar(ctx, v)
	if err != nil {
		return nil, errormsg.WrapErr(svcerr.OrderSVCErrorGRPCClient, err, "error grpc client")
	}

	client.Release()
	return res, nil
}

func (c *CarDep) DeleteGRPC(ctx context.Context, v *grpcmodel.DeleteCarRequest) (*grpcmodel.DeleteCarReply, error) {
	client, err := c.Grpc.Get()
	if err != nil {
		return nil, errormsg.WrapErr(svcerr.OrderSVCErrorGRPCClient, err, "error grpc client connection")
	}

	clientService := grpcmodel.NewOrderClient(client.Conn)

	ctx, cancel := context.WithTimeout(ctx, time.Second)
	defer cancel()
	res, err := clientService.DeleteCar(ctx, v)
	if err != nil {
		return nil, errormsg.WrapErr(svcerr.OrderSVCErrorGRPCClient, err, "error grpc client")
	}

	client.Release()
	return res, nil
}

func (c *CarDep) GetByIDGRPC(ctx context.Context, v *grpcmodel.GetCarByIDRequest) (*grpcmodel.SingleCarReply, error) {
	client, err := c.Grpc.Get()
	if err != nil {
		return nil, errormsg.WrapErr(svcerr.OrderSVCErrorGRPCClient, err, "error grpc client connection")
	}

	clientService := grpcmodel.NewOrderClient(client.Conn)

	ctx, cancel := context.WithTimeout(ctx, time.Second)
	defer cancel()
	res, err := clientService.GetCarByID(ctx, v)
	if err != nil {
		return nil, errormsg.WrapErr(svcerr.OrderSVCErrorGRPCClient, err, "error grpc client")
	}

	client.Release()
	return res, nil
}

func (c *CarDep) GetCarByParam(ctx context.Context, v *grpcmodel.GetCarByParamRequest) (*grpcmodel.GetCarByParamReply, error) {
	client, err := c.Grpc.Get()
	if err != nil {
		return nil, errormsg.WrapErr(svcerr.OrderSVCErrorGRPCClient, err, "error grpc client connection")
	}

	clientService := grpcmodel.NewOrderClient(client.Conn)

	ctx, cancel := context.WithTimeout(ctx, time.Second)
	defer cancel()
	res, err := clientService.GetCarByParam(ctx, v)
	if err != nil {
		return nil, errormsg.WrapErr(svcerr.OrderSVCErrorGRPCClient, err, "error grpc client")
	}

	client.Release()
	return res, nil
}
