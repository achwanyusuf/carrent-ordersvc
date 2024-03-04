package grpc

import (
	"context"

	"github.com/achwanyusuf/carrent-lib/pkg/logger"
	"github.com/achwanyusuf/carrent-ordersvc/src/model/grpcmodel"
	"github.com/achwanyusuf/carrent-ordersvc/src/usecase"
)

type GrpcDep struct {
	grpcmodel.UnimplementedOrderServer
	Conf    Config
	Log     *logger.Logger
	Usecase *usecase.UsecaseInterface
}

type Config struct{}

type GRPCInterface interface {
	CreateCar(ctx context.Context, v *grpcmodel.CreateCarRequest) (*grpcmodel.SingleCarReply, error)
	UpdateCar(ctx context.Context, v *grpcmodel.UpdateCarRequest) (*grpcmodel.SingleCarReply, error)
	GetCarByParam(ctx context.Context, v *grpcmodel.GetCarByParamRequest) (*grpcmodel.GetCarByParamReply, error)
	DeleteCar(ctx context.Context, v *grpcmodel.DeleteCarRequest) (*grpcmodel.DeleteCarReply, error)
	GetCarByID(ctx context.Context, v *grpcmodel.GetCarByIDRequest) (*grpcmodel.SingleCarReply, error)

	CreateOrder(ctx context.Context, v *grpcmodel.CreateOrderRequest) (*grpcmodel.SingleOrderReply, error)
	UpdateOrder(ctx context.Context, v *grpcmodel.UpdateOrderRequest) (*grpcmodel.SingleOrderReply, error)
	DeleteOrder(ctx context.Context, v *grpcmodel.DeleteOrderRequest) (*grpcmodel.DeleteOrderReply, error)
	GetOrderByID(ctx context.Context, v *grpcmodel.GetOrderByIDRequest) (*grpcmodel.SingleOrderReply, error)
	GetOrderByParam(ctx context.Context, v *grpcmodel.GetOrderByParamRequest) (*grpcmodel.GetOrderByParamReply, error)
}

func New(conf Config, log *logger.Logger, usecase *usecase.UsecaseInterface) *GrpcDep {
	return &GrpcDep{
		Conf:    conf,
		Log:     log,
		Usecase: usecase,
	}
}

func (g *GrpcDep) CreateCar(ctx context.Context, v *grpcmodel.CreateCarRequest) (*grpcmodel.SingleCarReply, error) {
	car, err := g.Usecase.Car.CreateGRPCProcess(&ctx, v)
	if err != nil {
		return &grpcmodel.SingleCarReply{}, err
	}

	return car, nil
}

func (g *GrpcDep) UpdateCar(ctx context.Context, v *grpcmodel.UpdateCarRequest) (*grpcmodel.SingleCarReply, error) {
	car, err := g.Usecase.Car.UpdateByIDGRPCProcess(&ctx, v)
	if err != nil {
		return &grpcmodel.SingleCarReply{}, err
	}

	return car, nil
}

func (g *GrpcDep) GetCarByParam(ctx context.Context, v *grpcmodel.GetCarByParamRequest) (*grpcmodel.GetCarByParamReply, error) {
	car, err := g.Usecase.Car.GetByParamGRPCProcess(&ctx, v)
	if err != nil {
		return &grpcmodel.GetCarByParamReply{}, err
	}

	return car, nil
}

func (g *GrpcDep) DeleteCar(ctx context.Context, v *grpcmodel.DeleteCarRequest) (*grpcmodel.DeleteCarReply, error) {
	car, err := g.Usecase.Car.DeleteByIDGRPCProccess(&ctx, v)
	if err != nil {
		return &grpcmodel.DeleteCarReply{}, err
	}

	return car, nil
}
func (g *GrpcDep) GetCarByID(ctx context.Context, v *grpcmodel.GetCarByIDRequest) (*grpcmodel.SingleCarReply, error) {
	car, err := g.Usecase.Car.GetByIDGRPCProcess(&ctx, v)
	if err != nil {
		return &grpcmodel.SingleCarReply{}, err
	}

	return car, nil
}

func (g *GrpcDep) CreateOrder(ctx context.Context, v *grpcmodel.CreateOrderRequest) (*grpcmodel.SingleOrderReply, error) {
	order, err := g.Usecase.Order.CreateGRPCProcess(&ctx, v)
	if err != nil {
		return &grpcmodel.SingleOrderReply{}, err
	}

	return order, nil
}

func (g *GrpcDep) UpdateOrder(ctx context.Context, v *grpcmodel.UpdateOrderRequest) (*grpcmodel.SingleOrderReply, error) {
	order, err := g.Usecase.Order.UpdateByIDGRPCProcess(&ctx, v)
	if err != nil {
		return &grpcmodel.SingleOrderReply{}, err
	}

	return order, nil
}

func (g *GrpcDep) DeleteOrder(ctx context.Context, v *grpcmodel.DeleteOrderRequest) (*grpcmodel.DeleteOrderReply, error) {
	order, err := g.Usecase.Order.DeleteByIDGRPCProccess(&ctx, v)
	if err != nil {
		return &grpcmodel.DeleteOrderReply{}, err
	}

	return order, nil
}
func (g *GrpcDep) GetOrderByID(ctx context.Context, v *grpcmodel.GetOrderByIDRequest) (*grpcmodel.SingleOrderReply, error) {
	order, err := g.Usecase.Order.GetByIDGRPCProcess(&ctx, v)
	if err != nil {
		return &grpcmodel.SingleOrderReply{}, err
	}

	return order, nil
}
func (g *GrpcDep) GetOrderByParam(ctx context.Context, v *grpcmodel.GetOrderByParamRequest) (*grpcmodel.GetOrderByParamReply, error) {
	order, err := g.Usecase.Order.GetByParamGRPCProcess(&ctx, v)
	if err != nil {
		return &grpcmodel.GetOrderByParamReply{}, err
	}

	return order, nil
}
