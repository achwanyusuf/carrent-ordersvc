package rest

import (
	"github.com/achwanyusuf/carrent-lib/pkg/httpserver"
	"github.com/achwanyusuf/carrent-lib/pkg/jwt"
	"github.com/achwanyusuf/carrent-lib/pkg/logger"
	"github.com/achwanyusuf/carrent-ordersvc/src/handler/rest/car"
	"github.com/achwanyusuf/carrent-ordersvc/src/handler/rest/order"
	"github.com/achwanyusuf/carrent-ordersvc/src/model"
	"github.com/achwanyusuf/carrent-ordersvc/src/usecase"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type RestDep struct {
	Conf     Config
	Log      *logger.Logger
	Usecase  *usecase.UsecaseInterface
	Gin      *gin.Engine
	Validate *validator.Validate
}

type Config struct {
	TokenSecret string     `mapstructure:"token_secret"`
	Car         car.Conf   `mapstructure:"car"`
	Order       order.Conf `mapstructure:"order"`
}

type RestInterface struct {
	car   car.CarInterface
	order order.OrderInterface
}

func New(r *RestDep) *RestInterface {
	return &RestInterface{
		car.New(r.Conf.Car, r.Log, r.Usecase.Car, r.Validate),
		order.New(r.Conf.Order, r.Log, r.Usecase.Order, r.Validate),
	}
}

func (r *RestDep) Serve(handler *RestInterface) {
	api := r.Gin.Group("/api")
	api.Use(jwt.JWT(*r.Log, []byte(r.Conf.TokenSecret)))
	{
		api.POST("/car", httpserver.ValidateScope(*r.Log, []string{model.SuperAdminScope, model.StoreScope, model.CustomerScope}), handler.car.Create)
		api.PUT("/car/:id", httpserver.ValidateScope(*r.Log, []string{model.SuperAdminScope, model.StoreScope, model.CustomerScope}), handler.car.UpdateByID)
		api.GET("/car", httpserver.ValidateScope(*r.Log, []string{model.SuperAdminScope, model.StoreScope, model.CustomerScope}), handler.car.Read)
		api.GET("/car/:id", httpserver.ValidateScope(*r.Log, []string{model.SuperAdminScope, model.StoreScope, model.CustomerScope}), handler.car.GetByID)
		api.DELETE("/car/:id", httpserver.ValidateScope(*r.Log, []string{model.SuperAdminScope, model.StoreScope, model.CustomerScope}), handler.car.DeleteByID)

		api.POST("/order", httpserver.ValidateScope(*r.Log, []string{model.SuperAdminScope, model.StoreScope, model.CustomerScope}), handler.order.Create)
		api.PUT("/order/:id", httpserver.ValidateScope(*r.Log, []string{model.SuperAdminScope, model.StoreScope, model.CustomerScope}), handler.order.UpdateByID)
		api.GET("/order", httpserver.ValidateScope(*r.Log, []string{model.SuperAdminScope, model.StoreScope, model.CustomerScope}), handler.order.Read)
		api.GET("/order/:id", httpserver.ValidateScope(*r.Log, []string{model.SuperAdminScope, model.StoreScope, model.CustomerScope}), handler.order.GetByID)
		api.DELETE("/order/:id", httpserver.ValidateScope(*r.Log, []string{model.SuperAdminScope, model.StoreScope, model.CustomerScope}), handler.order.DeleteByID)
	}
}
