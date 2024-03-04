package car

import (
	"context"
	"encoding/json"

	"github.com/achwanyusuf/carrent-ordersvc/src/model"
	"github.com/achwanyusuf/carrent-ordersvc/src/model/psqlmodel"
)

func (c *CarDep) getSingleByParamRedis(ctx *context.Context, key string) (psqlmodel.Car, error) {
	var res psqlmodel.Car
	data, err := c.Redis.Get(*ctx, key).Result()
	if err != nil {
		return res, err
	}
	err = json.Unmarshal([]byte(data), &res)
	if err != nil {
		return res, err
	}
	return res, nil
}

func (c *CarDep) setRedis(ctx *context.Context, key string, data string) error {
	expTime := c.Conf.RedisExpirationTime
	if c.Conf.RedisExpirationTime == 0 {
		expTime = model.DefaultRedisExpiration
	}
	_, err := c.Redis.Del(*ctx, key).Result()
	if err != nil {
		return err
	}
	_, err = c.Redis.Set(*ctx, key, data, expTime).Result()
	return err
}

func (c *CarDep) getByParamRedis(ctx *context.Context, key string) (psqlmodel.CarSlice, error) {
	var res psqlmodel.CarSlice
	data, err := c.Redis.Get(*ctx, key).Result()
	if err != nil {
		return res, err
	}
	err = json.Unmarshal([]byte(data), &res)
	if err != nil {
		return res, err
	}
	return res, nil
}

func (c *CarDep) getByParamPaginationRedis(ctx *context.Context, key string) (model.Pagination, error) {
	var res model.Pagination
	data, err := c.Redis.Get(*ctx, key).Result()
	if err != nil {
		return res, err
	}
	err = json.Unmarshal([]byte(data), &res)
	if err != nil {
		return res, err
	}
	return res, nil
}
