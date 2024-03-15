package car

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"time"

	"github.com/achwanyusuf/carrent-lib/pkg/errormsg"
	"github.com/achwanyusuf/carrent-lib/pkg/grpcclientpool"
	"github.com/achwanyusuf/carrent-lib/pkg/logger"
	"github.com/achwanyusuf/carrent-ordersvc/src/model"
	"github.com/achwanyusuf/carrent-ordersvc/src/model/grpcmodel"
	"github.com/achwanyusuf/carrent-ordersvc/src/model/psqlmodel"
	"github.com/achwanyusuf/carrent-ordersvc/src/model/svcerr"

	goredislib "github.com/redis/go-redis/v9"
)

type CarDep struct {
	Log   logger.Logger
	DB    *sql.DB
	Redis *goredislib.Client
	Conf  Conf
	Grpc  grpcclientpool.CPoolInterface
}

type Conf struct {
	DefaultPageLimit    int           `mapstructure:"page_limit"`
	RedisExpirationTime time.Duration `mapstructure:"expiration_time"`
}

type CarInterface interface {
	Insert(ctx *context.Context, data *psqlmodel.Car) error
	GetSingleByParam(ctx *context.Context, cacheControl string, param *model.GetCarByParam) (psqlmodel.Car, error)
	Update(ctx *context.Context, v *psqlmodel.Car) error
	Delete(ctx *context.Context, v *psqlmodel.Car, id int64, isHardDelete bool) error
	GetByParam(ctx *context.Context, cacheControl string, param *model.GetCarsByParam) (psqlmodel.CarSlice, model.Pagination, error)

	// grpc client
	InsertGRPC(ctx context.Context, v *grpcmodel.CreateCarRequest) (*grpcmodel.SingleCarReply, error)
	GetByIDGRPC(ctx context.Context, v *grpcmodel.GetCarByIDRequest) (*grpcmodel.SingleCarReply, error)
	GetCarByParam(ctx context.Context, v *grpcmodel.GetCarByParamRequest) (*grpcmodel.GetCarByParamReply, error)
	DeleteGRPC(ctx context.Context, v *grpcmodel.DeleteCarRequest) (*grpcmodel.DeleteCarReply, error)
	UpdateGRPC(ctx context.Context, v *grpcmodel.UpdateCarRequest) (*grpcmodel.SingleCarReply, error)
}

func New(conf Conf, log *logger.Logger, db *sql.DB, rds *goredislib.Client, grpc *grpcclientpool.CPoolInterface) CarInterface {
	return &CarDep{
		Log:   *log,
		DB:    db,
		Redis: rds,
		Conf:  conf,
		Grpc:  *grpc,
	}
}

func (c *CarDep) Insert(ctx *context.Context, data *psqlmodel.Car) error {
	return c.insertPSQL(ctx, data)
}

func (c *CarDep) GetSingleByParam(ctx *context.Context, cacheControl string, param *model.GetCarByParam) (psqlmodel.Car, error) {
	str, err := json.Marshal(param)
	if err != nil {
		return psqlmodel.Car{}, errormsg.WrapErr(svcerr.OrderSVCBadRequest, err, "error marshal param")
	}

	key := fmt.Sprintf(model.GetSingleByParamCarKey, str)
	if cacheControl != model.MustRevalidate {
		res, err := c.getSingleByParamRedis(ctx, key)
		if err != nil {
			if err == goredislib.Nil {
				res, err := c.getSingleByParamPSQL(ctx, param)
				if err == nil {
					dataStr, err := json.Marshal(&res)
					if err != nil {
						return res, errormsg.WrapErr(svcerr.OrderSVCBadRequest, err, "error get psql")
					}
					err = c.setRedis(ctx, key, string(dataStr))
					if err != nil {
						return res, errormsg.WrapErr(svcerr.OrderSVCBadRequest, err, "error set redis")
					}
				}
				return res, err
			}
			return res, errormsg.WrapErr(svcerr.OrderSVCBadRequest, err, "error marshal param")
		}
		return res, nil
	}

	res, err := c.getSingleByParamPSQL(ctx, param)
	if err == nil {
		dataStr, err := json.Marshal(&res)
		if err != nil {
			return res, errormsg.WrapErr(svcerr.OrderSVCBadRequest, err, "error get psql")
		}
		err = c.setRedis(ctx, key, string(dataStr))
		if err != nil {
			return res, errormsg.WrapErr(svcerr.OrderSVCBadRequest, err, "error set redis")
		}
	}
	return res, err
}

func (c *CarDep) Update(ctx *context.Context, v *psqlmodel.Car) error {
	return c.updatePSQL(ctx, v)
}

func (c *CarDep) Delete(ctx *context.Context, v *psqlmodel.Car, id int64, isHardDelete bool) error {
	return c.deletePSQL(ctx, v, id, isHardDelete)
}
func (c *CarDep) GetByParam(ctx *context.Context, cacheControl string, param *model.GetCarsByParam) (psqlmodel.CarSlice, model.Pagination, error) {
	var pg model.Pagination
	var res psqlmodel.CarSlice

	str, err := json.Marshal(param)
	if err != nil {
		return psqlmodel.CarSlice{}, pg, errormsg.WrapErr(svcerr.OrderSVCBadRequest, err, "error marshal param")
	}

	key := fmt.Sprintf(model.GetByParamCarKey, str)
	keyPg := fmt.Sprintf(model.GetByParamCarPgKey, str)
	if cacheControl != model.MustRevalidate {
		res, err1 := c.getByParamRedis(ctx, key)
		pg, err2 := c.getByParamPaginationRedis(ctx, keyPg)
		if err1 != nil || err2 != nil {
			if err1 == goredislib.Nil || err2 == goredislib.Nil {
				res, pg, err := c.getByParamPSQL(ctx, param)
				if err == nil {
					dataStr, err := json.Marshal(&res)
					if err != nil {
						return res, pg, errormsg.WrapErr(svcerr.OrderSVCBadRequest, err, "error get psql")
					}
					err = c.setRedis(ctx, key, string(dataStr))
					if err != nil {
						return res, pg, errormsg.WrapErr(svcerr.OrderSVCBadRequest, err, "error set redis")
					}
					dataStr, err = json.Marshal(&pg)
					if err != nil {
						return res, pg, errormsg.WrapErr(svcerr.OrderSVCBadRequest, err, "error get psql")
					}
					err = c.setRedis(ctx, keyPg, string(dataStr))
					if err != nil {
						return res, pg, errormsg.WrapErr(svcerr.OrderSVCBadRequest, err, "error set redis")
					}
				}
				return res, pg, err
			}
			if err1 != nil {
				return res, pg, errormsg.WrapErr(svcerr.OrderSVCBadRequest, err1, "error redis")
			}
			if err2 != nil {
				return res, pg, errormsg.WrapErr(svcerr.OrderSVCBadRequest, err2, "error redis")
			}
		}
		return res, pg, nil
	}

	res, pg, err = c.getByParamPSQL(ctx, param)
	if err == nil {
		dataStr, err := json.Marshal(&res)
		if err != nil {
			return res, pg, errormsg.WrapErr(svcerr.OrderSVCBadRequest, err, "error get psql")
		}
		err = c.setRedis(ctx, key, string(dataStr))
		if err != nil {
			return res, pg, errormsg.WrapErr(svcerr.OrderSVCBadRequest, err, "error set redis")
		}
		dataStr, err = json.Marshal(&pg)
		if err != nil {
			return res, pg, errormsg.WrapErr(svcerr.OrderSVCBadRequest, err, "error get psql")
		}
		err = c.setRedis(ctx, keyPg, string(dataStr))
		if err != nil {
			return res, pg, errormsg.WrapErr(svcerr.OrderSVCBadRequest, err, "error set redis")
		}
	}
	return res, pg, err
}
