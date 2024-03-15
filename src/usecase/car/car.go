package car

import (
	"context"

	"github.com/achwanyusuf/carrent-lib/pkg/errormsg"
	"github.com/achwanyusuf/carrent-lib/pkg/logger"
	"github.com/achwanyusuf/carrent-ordersvc/src/domain/car"
	"github.com/achwanyusuf/carrent-ordersvc/src/model"
	"github.com/achwanyusuf/carrent-ordersvc/src/model/grpcmodel"
	"github.com/achwanyusuf/carrent-ordersvc/src/model/psqlmodel"
	"github.com/achwanyusuf/carrent-ordersvc/src/model/svcerr"
	"github.com/gin-gonic/gin"
	"github.com/volatiletech/null/v8"
)

type CarDep struct {
	log  logger.Logger
	conf Conf
	car  car.CarInterface
}

type Conf struct{}

type CarInterface interface {
	Create(ctx *gin.Context, v model.CreateCar) (model.Car, error)
	CreateGRPCProcess(ctx *context.Context, v *grpcmodel.CreateCarRequest) (*grpcmodel.SingleCarReply, error)
	GetByParam(ctx *gin.Context, cacheControl string, v model.GetCarsByParam) ([]model.Car, model.Pagination, error)
	GetByParamGRPCProcess(ctx *context.Context, v *grpcmodel.GetCarByParamRequest) (*grpcmodel.GetCarByParamReply, error)
	GetByID(ctx *gin.Context, cacheControl string, id int64) (model.Car, error)
	GetByIDGRPCProcess(ctx *context.Context, v *grpcmodel.GetCarByIDRequest) (*grpcmodel.SingleCarReply, error)
	UpdateByID(ctx *gin.Context, id int64, v model.UpdateCar) (model.Car, error)
	UpdateByIDGRPCProcess(ctx *context.Context, v *grpcmodel.UpdateCarRequest) (*grpcmodel.SingleCarReply, error)
	DeleteByID(ctx *gin.Context, id int64, vid int64) error
	DeleteByIDGRPCProccess(ctx *context.Context, v *grpcmodel.DeleteCarRequest) (*grpcmodel.DeleteCarReply, error)
}

func New(conf Conf, logger *logger.Logger, car car.CarInterface) CarInterface {
	return &CarDep{
		conf: conf,
		log:  *logger,
		car:  car,
	}
}

func (c *CarDep) Create(ctx *gin.Context, v model.CreateCar) (model.Car, error) {
	var result model.Car
	err := v.Validate()
	if err != nil {
		return result, err
	}

	data, err := c.car.InsertGRPC(ctx, &grpcmodel.CreateCarRequest{
		CarName:   v.CarName,
		DayRate:   v.DayRate,
		MonthRate: v.MonthRate,
		Image:     v.Image,
		CreatedBy: v.CreatedBy,
	})
	if err != nil {
		return result, err
	}

	return model.TransformSingleCarReplyToCar(ctx, data, c.log), nil
}

func (c *CarDep) CreateGRPCProcess(ctx *context.Context, v *grpcmodel.CreateCarRequest) (*grpcmodel.SingleCarReply, error) {
	var result *grpcmodel.SingleCarReply
	car := &psqlmodel.Car{
		CarName:   v.CarName,
		DayRate:   v.DayRate,
		MonthRate: v.MonthRate,
		Image:     v.Image,
		CreatedBy: int(v.CreatedBy),
		UpdatedBy: int(v.CreatedBy),
	}

	err := c.car.Insert(ctx, car)
	if err != nil {
		return result, err
	}

	return model.TransformSingleCarReply(car), nil
}

func (c *CarDep) GetByParam(ctx *gin.Context, cacheControl string, v model.GetCarsByParam) ([]model.Car, model.Pagination, error) {
	param := v.FillGrpcClient()
	param.CacheControl = cacheControl
	carSlice, err := c.car.GetCarByParam(ctx, param)
	if err != nil {
		return []model.Car{}, model.Pagination{}, errormsg.WrapErr(svcerr.OrderSVCBadRequest, err, "error get by param")
	}
	cars, pagination := model.TransformCarByParamReplyToCar(ctx, carSlice, c.log)
	return cars, pagination, nil
}

func (c *CarDep) GetByParamGRPCProcess(ctx *context.Context, v *grpcmodel.GetCarByParamRequest) (*grpcmodel.GetCarByParamReply, error) {
	param := model.TransformGetCarByParamRequestToCarParam(v)
	carSlice, pagination, err := c.car.GetByParam(ctx, v.CacheControl, &param)
	if err != nil {
		return &grpcmodel.GetCarByParamReply{}, errormsg.WrapErr(svcerr.OrderSVCBadRequest, err, "error get by param")
	}
	return model.TransformCarToGetCarByParamReply(&carSlice, pagination), nil
}

func (c *CarDep) GetByID(ctx *gin.Context, cacheControl string, id int64) (model.Car, error) {
	car, err := c.car.GetByIDGRPC(ctx, &grpcmodel.GetCarByIDRequest{
		Id:           id,
		CacheControl: cacheControl,
	})
	if err != nil {
		return model.Car{}, err
	}
	return model.TransformSingleCarReplyToCar(ctx, car, c.log), nil
}

func (c *CarDep) GetByIDGRPCProcess(ctx *context.Context, v *grpcmodel.GetCarByIDRequest) (*grpcmodel.SingleCarReply, error) {
	car, err := c.car.GetSingleByParam(ctx, v.CacheControl, &model.GetCarByParam{
		ID: null.Int64From(v.Id),
	})
	if err != nil {
		return &grpcmodel.SingleCarReply{}, err
	}
	return model.TransformSingleCarReply(&car), nil
}

func (c *CarDep) UpdateByID(ctx *gin.Context, id int64, v model.UpdateCar) (model.Car, error) {
	updateData := v.FillGrpcClient()
	updateData.Id = id
	updateData.UpdatedBy = v.UpdatedBy
	car, err := c.car.UpdateGRPC(ctx, updateData)
	if err != nil {
		return model.Car{}, err
	}

	return model.TransformSingleCarReplyToCar(ctx, car, c.log), nil
}

func (c *CarDep) UpdateByIDGRPCProcess(ctx *context.Context, v *grpcmodel.UpdateCarRequest) (*grpcmodel.SingleCarReply, error) {
	car, err := c.car.GetSingleByParam(ctx, model.MustRevalidate, &model.GetCarByParam{
		ID: null.NewInt64(v.Id, true),
	})
	if err != nil {
		return &grpcmodel.SingleCarReply{}, err
	}

	if v.CarName == nil && v.DayRate == nil && v.Image == nil && v.MonthRate == nil {
		return model.TransformSingleCarReply(&car), nil
	}

	model.FillUpdateCar(&car, v)
	car.UpdatedBy = int(v.UpdatedBy)

	err = c.car.Update(ctx, &car)
	if err != nil {
		return &grpcmodel.SingleCarReply{}, err
	}

	return model.TransformSingleCarReply(&car), nil
}

func (c *CarDep) DeleteByID(ctx *gin.Context, id int64, vid int64) error {
	_, err := c.car.DeleteGRPC(ctx, &grpcmodel.DeleteCarRequest{
		Id:        vid,
		DeletedBy: id,
	})

	return err
}

func (c *CarDep) DeleteByIDGRPCProccess(ctx *context.Context, v *grpcmodel.DeleteCarRequest) (*grpcmodel.DeleteCarReply, error) {
	car, err := c.car.GetSingleByParam(ctx, model.MustRevalidate, &model.GetCarByParam{
		ID: null.NewInt64(v.Id, true),
	})
	if err != nil {
		return &grpcmodel.DeleteCarReply{}, err
	}
	err = c.car.Delete(ctx, &car, v.DeletedBy, false)
	if err != nil {
		return &grpcmodel.DeleteCarReply{}, err
	}
	return &grpcmodel.DeleteCarReply{
		Id: v.Id,
	}, nil
}
