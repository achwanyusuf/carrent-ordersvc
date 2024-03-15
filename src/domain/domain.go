package domain

import (
	"database/sql"

	"github.com/achwanyusuf/carrent-lib/pkg/grpcclientpool"
	"github.com/achwanyusuf/carrent-lib/pkg/logger"
	"github.com/achwanyusuf/carrent-ordersvc/src/domain/car"
	"github.com/achwanyusuf/carrent-ordersvc/src/domain/order"
	goredislib "github.com/redis/go-redis/v9"
)

type DomainDep struct {
	Conf  Config
	Log   *logger.Logger
	DB    *sql.DB
	Redis *goredislib.Client
	Grpc  *grpcclientpool.CPoolInterface
}

type Config struct {
	Car   car.Conf
	Order order.Conf
}

type DomainInterface struct {
	Car   car.CarInterface
	Order order.OrderInterface
}

func New(d *DomainDep) *DomainInterface {
	return &DomainInterface{
		car.New(d.Conf.Car, d.Log, d.DB, d.Redis, d.Grpc),
		order.New(d.Conf.Order, d.Log, d.DB, d.Redis, d.Grpc),
	}
}
