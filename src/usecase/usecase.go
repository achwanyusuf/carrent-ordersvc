package usecase

import (
	"github.com/achwanyusuf/carrent-lib/pkg/logger"
	"github.com/achwanyusuf/carrent-ordersvc/src/domain"
	"github.com/achwanyusuf/carrent-ordersvc/src/usecase/car"
	"github.com/achwanyusuf/carrent-ordersvc/src/usecase/order"
)

type UsecaseDep struct {
	Conf   Config
	Log    *logger.Logger
	Domain *domain.DomainInterface
}

type Config struct {
	Car   car.Conf
	Order order.Conf
}

type UsecaseInterface struct {
	Car   car.CarInterface
	Order order.OrderInterface
}

func New(u *UsecaseDep) *UsecaseInterface {
	return &UsecaseInterface{
		car.New(u.Conf.Car, u.Log, u.Domain.Car),
		order.New(u.Conf.Order, u.Log, u.Domain.Order, u.Domain.Car),
	}
}
