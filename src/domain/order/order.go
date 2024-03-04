package order

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

type OrderDep struct {
	Log   logger.Logger
	DB    *sql.DB
	Redis *goredislib.Client
	Conf  Conf
	Grpc  *grpcclientpool.CPool
}

type Conf struct {
	DefaultPageLimit    int           `mapstructure:"page_limit"`
	RedisExpirationTime time.Duration `mapstructure:"expiration_time"`
}

type OrderInterface interface {
	Insert(ctx *context.Context, data *psqlmodel.Order) error
	GetSingleByParam(ctx *context.Context, cacheControl string, param *model.GetOrderByParam) (psqlmodel.Order, error)
	Update(ctx *context.Context, v *psqlmodel.Order) error
	Delete(ctx *context.Context, v *psqlmodel.Order, id int64, isHardDelete bool) error
	GetByParam(ctx *context.Context, cacheControl string, param *model.GetOrdersByParam) (psqlmodel.OrderSlice, model.Pagination, error)

	// grpc client
	InsertGRPC(ctx context.Context, v *grpcmodel.CreateOrderRequest) (*grpcmodel.SingleOrderReply, error)
	GetByIDGRPC(ctx context.Context, v *grpcmodel.GetOrderByIDRequest) (*grpcmodel.SingleOrderReply, error)
	GetOrderByParam(ctx context.Context, v *grpcmodel.GetOrderByParamRequest) (*grpcmodel.GetOrderByParamReply, error)
	DeleteGRPC(ctx context.Context, v *grpcmodel.DeleteOrderRequest) (*grpcmodel.DeleteOrderReply, error)
	UpdateGRPC(ctx context.Context, v *grpcmodel.UpdateOrderRequest) (*grpcmodel.SingleOrderReply, error)
}

func New(conf Conf, log *logger.Logger, db *sql.DB, rds *goredislib.Client, grpc *grpcclientpool.CPool) OrderInterface {
	return &OrderDep{
		Log:   *log,
		DB:    db,
		Redis: rds,
		Conf:  conf,
		Grpc:  grpc,
	}
}

func (o *OrderDep) Insert(ctx *context.Context, data *psqlmodel.Order) error {
	return o.insertPSQL(ctx, data)
}

func (o *OrderDep) GetSingleByParam(ctx *context.Context, cacheControl string, param *model.GetOrderByParam) (psqlmodel.Order, error) {
	str, err := json.Marshal(param)
	if err != nil {
		return psqlmodel.Order{}, errormsg.WrapErr(svcerr.OrderSVCBadRequest, err, "error marshal param")
	}

	key := fmt.Sprintf(model.GetSingleByParamOrderKey, str)
	if cacheControl != model.MustRevalidate {
		res, err := o.getSingleByParamRedis(ctx, key)
		if err != nil {
			if err == goredislib.Nil {
				res, err := o.getSingleByParamPSQL(ctx, param)
				if err == nil {
					dataStr, err := json.Marshal(&res)
					if err != nil {
						return res, errormsg.WrapErr(svcerr.OrderSVCBadRequest, err, "error get psql")
					}
					err = o.setRedis(ctx, key, string(dataStr))
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

	res, err := o.getSingleByParamPSQL(ctx, param)
	if err == nil {
		dataStr, err := json.Marshal(&res)
		if err != nil {
			return res, errormsg.WrapErr(svcerr.OrderSVCBadRequest, err, "error get psql")
		}
		err = o.setRedis(ctx, key, string(dataStr))
		if err != nil {
			return res, errormsg.WrapErr(svcerr.OrderSVCBadRequest, err, "error set redis")
		}
	}
	return res, err
}

func (o *OrderDep) Update(ctx *context.Context, v *psqlmodel.Order) error {
	return o.updatePSQL(ctx, v)
}

func (o *OrderDep) Delete(ctx *context.Context, v *psqlmodel.Order, id int64, isHardDelete bool) error {
	return o.deletePSQL(ctx, v, id, isHardDelete)
}
func (o *OrderDep) GetByParam(ctx *context.Context, cacheControl string, param *model.GetOrdersByParam) (psqlmodel.OrderSlice, model.Pagination, error) {
	var pg model.Pagination
	var res psqlmodel.OrderSlice

	str, err := json.Marshal(param)
	if err != nil {
		return psqlmodel.OrderSlice{}, pg, errormsg.WrapErr(svcerr.OrderSVCBadRequest, err, "error marshal param")
	}

	key := fmt.Sprintf(model.GetByParamOrderKey, str)
	keyPg := fmt.Sprintf(model.GetByParamOrderPgKey, str)
	if cacheControl != model.MustRevalidate {
		res, err1 := o.getByParamRedis(ctx, key)
		pg, err2 := o.getByParamPaginationRedis(ctx, keyPg)
		if err1 != nil || err2 != nil {
			if err1 == goredislib.Nil || err2 == goredislib.Nil {
				res, pg, err := o.getByParamPSQL(ctx, param)
				if err == nil {
					dataStr, err := json.Marshal(&res)
					if err != nil {
						return res, pg, errormsg.WrapErr(svcerr.OrderSVCBadRequest, err, "error get psql")
					}
					err = o.setRedis(ctx, key, string(dataStr))
					if err != nil {
						return res, pg, errormsg.WrapErr(svcerr.OrderSVCBadRequest, err, "error set redis")
					}
					dataStr, err = json.Marshal(&pg)
					if err != nil {
						return res, pg, errormsg.WrapErr(svcerr.OrderSVCBadRequest, err, "error get psql")
					}
					err = o.setRedis(ctx, key, string(dataStr))
					if err != nil {
						return res, pg, errormsg.WrapErr(svcerr.OrderSVCBadRequest, err, "error set redis")
					}
				}
				return res, pg, err
			}
			return res, pg, errormsg.WrapErr(svcerr.OrderSVCBadRequest, err, "error marshal param")
		}
		return res, pg, nil
	}

	res, pg, err = o.getByParamPSQL(ctx, param)
	if err == nil {
		dataStr, err := json.Marshal(&res)
		if err != nil {
			return res, pg, errormsg.WrapErr(svcerr.OrderSVCBadRequest, err, "error get psql")
		}
		err = o.setRedis(ctx, key, string(dataStr))
		if err != nil {
			return res, pg, errormsg.WrapErr(svcerr.OrderSVCBadRequest, err, "error set redis")
		}
		dataStr, err = json.Marshal(&pg)
		if err != nil {
			return res, pg, errormsg.WrapErr(svcerr.OrderSVCBadRequest, err, "error get psql")
		}
		err = o.setRedis(ctx, key, string(dataStr))
		if err != nil {
			return res, pg, errormsg.WrapErr(svcerr.OrderSVCBadRequest, err, "error set redis")
		}
	}
	return res, pg, err
}
