package order

import (
	"context"
	"encoding/json"

	"github.com/achwanyusuf/carrent-ordersvc/src/model"
	"github.com/achwanyusuf/carrent-ordersvc/src/model/psqlmodel"
)

func (o *OrderDep) getSingleByParamRedis(ctx *context.Context, key string) (psqlmodel.Order, error) {
	var res psqlmodel.Order
	data, err := o.Redis.Get(*ctx, key).Result()
	if err != nil {
		return res, err
	}
	err = json.Unmarshal([]byte(data), &res)
	if err != nil {
		return res, err
	}
	return res, nil
}

func (o *OrderDep) setRedis(ctx *context.Context, key string, data string) error {
	expTime := o.Conf.RedisExpirationTime
	if o.Conf.RedisExpirationTime == 0 {
		expTime = model.DefaultRedisExpiration
	}
	_, err := o.Redis.Del(*ctx, key).Result()
	if err != nil {
		return err
	}
	_, err = o.Redis.Set(*ctx, key, data, expTime).Result()
	return err
}

func (o *OrderDep) getByParamRedis(ctx *context.Context, key string) (psqlmodel.OrderSlice, error) {
	var res psqlmodel.OrderSlice
	data, err := o.Redis.Get(*ctx, key).Result()
	if err != nil {
		return res, err
	}
	err = json.Unmarshal([]byte(data), &res)
	if err != nil {
		return res, err
	}
	return res, nil
}

func (o *OrderDep) getByParamPaginationRedis(ctx *context.Context, key string) (model.Pagination, error) {
	var res model.Pagination
	data, err := o.Redis.Get(*ctx, key).Result()
	if err != nil {
		return res, err
	}
	err = json.Unmarshal([]byte(data), &res)
	if err != nil {
		return res, err
	}
	return res, nil
}
