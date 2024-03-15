package order

import (
	"context"
	"time"

	"github.com/achwanyusuf/carrent-lib/pkg/errormsg"
	"github.com/achwanyusuf/carrent-lib/pkg/logger"
	"github.com/achwanyusuf/carrent-ordersvc/src/domain/car"
	"github.com/achwanyusuf/carrent-ordersvc/src/domain/order"
	"github.com/achwanyusuf/carrent-ordersvc/src/model"
	"github.com/achwanyusuf/carrent-ordersvc/src/model/grpcmodel"
	"github.com/achwanyusuf/carrent-ordersvc/src/model/psqlmodel"
	"github.com/achwanyusuf/carrent-ordersvc/src/model/svcerr"
	"github.com/gin-gonic/gin"
	"github.com/volatiletech/null/v8"
)

type OrderDep struct {
	log   logger.Logger
	conf  Conf
	order order.OrderInterface
	car   car.CarInterface
}

type Conf struct{}

type OrderInterface interface {
	Create(ctx *gin.Context, v model.CreateOrder) (model.Order, error)
	CreateGRPCProcess(ctx *context.Context, v *grpcmodel.CreateOrderRequest) (*grpcmodel.SingleOrderReply, error)
	GetByParam(ctx *gin.Context, cacheControl string, v model.GetOrdersByParam) ([]model.Order, model.Pagination, error)
	GetByParamGRPCProcess(ctx *context.Context, v *grpcmodel.GetOrderByParamRequest) (*grpcmodel.GetOrderByParamReply, error)
	GetByID(ctx *gin.Context, cacheControl string, id int64) (model.Order, error)
	GetByIDGRPCProcess(ctx *context.Context, v *grpcmodel.GetOrderByIDRequest) (*grpcmodel.SingleOrderReply, error)
	UpdateByID(ctx *gin.Context, id int64, v model.UpdateOrder) (model.Order, error)
	UpdateByIDGRPCProcess(ctx *context.Context, v *grpcmodel.UpdateOrderRequest) (*grpcmodel.SingleOrderReply, error)
	DeleteByID(ctx *gin.Context, id int64, vid int64) error
	DeleteByIDGRPCProccess(ctx *context.Context, v *grpcmodel.DeleteOrderRequest) (*grpcmodel.DeleteOrderReply, error)
}

func New(conf Conf, logger *logger.Logger, order order.OrderInterface, car car.CarInterface) OrderInterface {
	return &OrderDep{
		conf:  conf,
		log:   *logger,
		order: order,
		car:   car,
	}
}

func (c *OrderDep) Create(ctx *gin.Context, v model.CreateOrder) (model.Order, error) {
	var result model.Order
	err := v.Validate()
	if err != nil {
		return result, err
	}

	data, err := c.order.InsertGRPC(ctx, &grpcmodel.CreateOrderRequest{
		CarId:           v.CarID,
		OrderDate:       v.OrderDate.Format(time.RFC3339),
		PickupDate:      v.PickupDate.Format(time.RFC3339),
		DropoffDate:     v.DropoffDate.Format(time.RFC3339),
		PickupLocation:  v.PickupLocation,
		PickupLat:       v.PickupLat,
		PickupLong:      v.PickupLong,
		DropoffLocation: v.DropoffLocation,
		DropoffLat:      v.DropoffLat,
		DropoffLong:     v.DropoffLong,
		CreatedBy:       v.CreatedBy,
	})
	if err != nil {
		return result, err
	}

	return model.TransformSingleOrderReplyToOrder(ctx, data, c.log), nil
}

func (c *OrderDep) CreateGRPCProcess(ctx *context.Context, v *grpcmodel.CreateOrderRequest) (*grpcmodel.SingleOrderReply, error) {
	var result *grpcmodel.SingleOrderReply
	_, err := c.car.GetSingleByParam(ctx, model.MustRevalidate, &model.GetCarByParam{
		ID: null.Int64From(v.CarId),
	})
	if err != nil {
		return &grpcmodel.SingleOrderReply{}, err
	}
	orderDate, err := time.Parse(time.RFC3339, v.OrderDate)
	if err != nil {
		return result, errormsg.WrapErr(svcerr.OrderSVCBadRequest, err, "error parse order date")
	}
	pickupDate, err := time.Parse(time.RFC3339, v.PickupDate)
	if err != nil {
		return result, errormsg.WrapErr(svcerr.OrderSVCBadRequest, err, "error parse pickup date")
	}
	dropoffDate, err := time.Parse(time.RFC3339, v.DropoffDate)
	if err != nil {
		return result, errormsg.WrapErr(svcerr.OrderSVCBadRequest, err, "error parse dropoff date")
	}

	order := &psqlmodel.Order{
		CarID:           int(v.CarId),
		OrderDate:       orderDate,
		PickupDate:      pickupDate,
		DropoffDate:     dropoffDate,
		PickupLocation:  v.PickupLocation,
		PickupLat:       v.PickupLat,
		PickupLong:      v.PickupLong,
		DropoffLocation: v.DropoffLocation,
		DropoffLat:      v.DropoffLat,
		DropoffLong:     v.DropoffLong,
		CreatedBy:       int(v.CreatedBy),
		UpdatedBy:       int(v.CreatedBy),
	}

	err = c.order.Insert(ctx, order)
	if err != nil {
		return result, err
	}

	return model.TransformSingleOrderReply(order), nil
}

func (c *OrderDep) GetByParam(ctx *gin.Context, cacheControl string, v model.GetOrdersByParam) ([]model.Order, model.Pagination, error) {
	param := v.FillGrpcClient()
	param.CacheControl = cacheControl
	orderSlice, err := c.order.GetOrderByParam(ctx, param)
	if err != nil {
		return []model.Order{}, model.Pagination{}, errormsg.WrapErr(svcerr.OrderSVCBadRequest, err, "error get by param")
	}
	orders, pagination := model.TransformOrderByParamReplyToOrder(ctx, orderSlice, c.log)
	return orders, pagination, nil
}

func (c *OrderDep) GetByParamGRPCProcess(ctx *context.Context, v *grpcmodel.GetOrderByParamRequest) (*grpcmodel.GetOrderByParamReply, error) {
	param := model.TransformGetOrderByParamRequestToOrderParam(*ctx, v, c.log)
	orderSlice, pagination, err := c.order.GetByParam(ctx, v.CacheControl, &param)
	if err != nil {
		return &grpcmodel.GetOrderByParamReply{}, errormsg.WrapErr(svcerr.OrderSVCBadRequest, err, "error get by param")
	}
	return model.TransformOrderToGetOrderByParamReply(&orderSlice, pagination), nil
}

func (c *OrderDep) GetByID(ctx *gin.Context, cacheControl string, id int64) (model.Order, error) {
	order, err := c.order.GetByIDGRPC(ctx, &grpcmodel.GetOrderByIDRequest{
		Id:           id,
		CacheControl: cacheControl,
	})
	if err != nil {
		return model.Order{}, err
	}
	return model.TransformSingleOrderReplyToOrder(ctx, order, c.log), nil
}

func (c *OrderDep) GetByIDGRPCProcess(ctx *context.Context, v *grpcmodel.GetOrderByIDRequest) (*grpcmodel.SingleOrderReply, error) {
	order, err := c.order.GetSingleByParam(ctx, v.CacheControl, &model.GetOrderByParam{
		ID: null.Int64From(v.Id),
	})
	if err != nil {
		return &grpcmodel.SingleOrderReply{}, err
	}
	return model.TransformSingleOrderReply(&order), nil
}

func (c *OrderDep) UpdateByID(ctx *gin.Context, id int64, v model.UpdateOrder) (model.Order, error) {
	updateData := v.FillGrpcClient()
	updateData.Id = id
	updateData.UpdatedBy = v.UpdatedBy
	order, err := c.order.UpdateGRPC(ctx, updateData)
	if err != nil {
		return model.Order{}, err
	}

	return model.TransformSingleOrderReplyToOrder(ctx, order, c.log), nil
}

func (c *OrderDep) UpdateByIDGRPCProcess(ctx *context.Context, v *grpcmodel.UpdateOrderRequest) (*grpcmodel.SingleOrderReply, error) {
	order, err := c.order.GetSingleByParam(ctx, model.MustRevalidate, &model.GetOrderByParam{
		ID: null.NewInt64(v.Id, true),
	})
	if err != nil {
		return &grpcmodel.SingleOrderReply{}, err
	}

	if v.CarId == nil && v.OrderDate == nil && v.PickupDate == nil && v.DropoffDate == nil &&
		v.PickupLocation == nil && v.PickupLat == nil && v.PickupLong == nil && v.DropoffLocation == nil &&
		v.DropoffLat == nil && v.DropoffLong == nil {
		return model.TransformSingleOrderReply(&order), nil
	}

	model.FillUpdateOrder(*ctx, c.log, &order, v)
	order.UpdatedBy = int(v.UpdatedBy)

	err = c.order.Update(ctx, &order)
	if err != nil {
		return &grpcmodel.SingleOrderReply{}, err
	}

	return model.TransformSingleOrderReply(&order), nil
}

func (c *OrderDep) DeleteByID(ctx *gin.Context, id int64, vid int64) error {
	_, err := c.order.DeleteGRPC(ctx, &grpcmodel.DeleteOrderRequest{
		Id:        vid,
		DeletedBy: id,
	})

	return err
}

func (c *OrderDep) DeleteByIDGRPCProccess(ctx *context.Context, v *grpcmodel.DeleteOrderRequest) (*grpcmodel.DeleteOrderReply, error) {
	order, err := c.order.GetSingleByParam(ctx, model.MustRevalidate, &model.GetOrderByParam{
		ID: null.NewInt64(v.Id, true),
	})
	if err != nil {
		return &grpcmodel.DeleteOrderReply{}, err
	}
	err = c.order.Delete(ctx, &order, v.DeletedBy, false)
	if err != nil {
		return &grpcmodel.DeleteOrderReply{}, err
	}
	return &grpcmodel.DeleteOrderReply{
		Id: v.Id,
	}, nil
}
