package conf

import (
	"github.com/achwanyusuf/carrent-lib/pkg/httpserver"
	"github.com/achwanyusuf/carrent-lib/pkg/psql"
	"github.com/achwanyusuf/carrent-lib/pkg/redis"
	"github.com/achwanyusuf/carrent-ordersvc/src/domain"
	"github.com/achwanyusuf/carrent-ordersvc/src/handler/rest"
	"github.com/achwanyusuf/carrent-ordersvc/src/usecase"
)

type Config struct {
	App     App            `mapstructure:"app"`
	Rest    rest.Config    `mapstructure:"rest"`
	Usecase usecase.Config `mapstructure:"usecase"`
	Domain  domain.Config  `mapstructure:"domain"`
}

type GRPC struct {
	Host       string `mapstructure:"host"`
	Port       int    `mapstructure:"port"`
	ServerCert string `mapstructure:"server_cert"`
	ServerKey  string `mapstructure:"server_key"`
	ClientCert string `mapstructure:"client_cert"`
	ClientHost string `mapstructure:"client_host"`
}

type App struct {
	Env        string                `mapstructure:"env"`
	HTTPServer httpserver.HTTPServer `mapstructure:"http_server"`
	Swagger    httpserver.Swagger    `mapstructure:"swagger"`
	PSQL       psql.PSQL             `mapstructure:"psql"`
	Redis      redis.Redis           `mapstructure:"redis"`
	GRPC       GRPC                  `mapstructure:"grpc"`
}
