package main

import (
	"context"
	"flag"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/achwanyusuf/carrent-lib/pkg/grpcclientpool"
	"github.com/achwanyusuf/carrent-lib/pkg/httpserver"
	"github.com/achwanyusuf/carrent-lib/pkg/logger"
	"github.com/achwanyusuf/carrent-lib/pkg/migration"
	"github.com/achwanyusuf/carrent-lib/pkg/psql"
	"github.com/achwanyusuf/carrent-lib/pkg/redis"
	"github.com/achwanyusuf/carrent-ordersvc/conf"
	"github.com/achwanyusuf/carrent-ordersvc/docs"
	"github.com/achwanyusuf/carrent-ordersvc/src/domain"
	grpcHandler "github.com/achwanyusuf/carrent-ordersvc/src/handler/grpc"
	"github.com/achwanyusuf/carrent-ordersvc/src/handler/rest"
	"github.com/achwanyusuf/carrent-ordersvc/src/model/grpcmodel"
	"github.com/achwanyusuf/carrent-ordersvc/src/usecase"
)

// @contact.name   CarRent Support
// @contact.url		https://www.carrent.com/support
// @contact.email 	support@carrent.com

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html
// @securitydefinitions.oauth2.password OAuth2Password
// @tokenUrl %s
var (
	staticConfPath, Namespace, BuildTime, Version string
	migrateup, migratedown, runHTTP, runGRPC      bool
	OAuth2PasswordTokenUrl                        string
)

func main() {
	flag.StringVar(&staticConfPath, "staticConfPath", "./conf/conf.yaml", "config path")
	flag.BoolVar(&migrateup, "migrateup", false, "run migration up")
	flag.BoolVar(&migratedown, "migratedown", false, "run migration up")
	flag.BoolVar(&runHTTP, "http", true, "run http")
	flag.BoolVar(&runGRPC, "grpc", false, "run grpc")
	flag.Parse()
	cfg, err := conf.New(staticConfPath)
	if err != nil {
		panic(err)
	}

	if Version == "" {
		Version = "v1.0.0"
	}
	// setup logger
	log := logger.New(&logger.Config{
		IsFile: false,
		Level:  logger.LevelDebug,
		CustomFields: map[string]interface{}{
			"namespace":  Namespace,
			"version":    Version,
			"build_time": BuildTime,
			"pid":        os.Getpid(),
		},
	})

	// setup db connection
	psql := psql.PsqlConnect(cfg.App.PSQL)

	if migrateup {
		err = migration.Migrate(psql, cfg.App.PSQL.MigrationPath, true)
		if err != nil {
			log.Error(context.Background(), err)
		}
		return
	}

	if migratedown {
		err = migration.Migrate(psql, cfg.App.PSQL.MigrationPath, false)
		if err != nil {
			log.Error(context.Background(), err)
		}
		return
	}

	// setup redis connection
	redis := redis.RedisConnect(cfg.App.Redis)

	tlsCredentials, err := loadClientTLSCredentials(cfg.App.GRPC.ClientCert, cfg.App.GRPC.ClientHost)
	if err != nil {
		log.Fatal(context.Background(), "cannot load TLS credentials: %v", err)
	}
	grpcClient := grpcclientpool.New(&grpcclientpool.ClientPoolGRPC{
		MaxOpenConnection: 10,
		MaxIdleConnection: 10,
		QueueTotal:        10000,
		Address:           ":9091",
		Credential:        tlsCredentials,
	})

	// init domain
	dom := domain.New(&domain.DomainDep{
		Conf:  cfg.Domain,
		Log:   &log,
		DB:    psql,
		Redis: redis,
		Grpc:  &grpcClient,
	})

	// init usecase
	uc := usecase.New(&usecase.UsecaseDep{
		Conf:   cfg.Usecase,
		Log:    &log,
		Domain: dom,
	})

	readSignal := make(chan os.Signal, 1)

	signal.Notify(
		readSignal,
		syscall.SIGTERM,
		syscall.SIGINT,
	)

	if runGRPC {
		go func() {
			// setup grpc connection
			grpcSetting := GRPC{
				Host: cfg.App.GRPC.Host,
				Port: cfg.App.GRPC.Port,
				Log:  log,
			}
			grpc := grpcSetting.newGRPC(cfg.App.GRPC.ServerCert, cfg.App.GRPC.ServerKey)
			grpcmodel.RegisterOrderServer(grpc, grpcHandler.New(grpcHandler.Config{}, &log, uc))
			log.Info(context.Background(), "server listening at %v", listener.Addr())
			if err := grpc.Serve(listener); err != nil {
				log.Error(context.Background(), "failed to serve: %v", err)
				panic(err)
			}
		}()
	}

	if runHTTP {
		go func() {
			cfg.App.Swagger.Title = Namespace
			cfg.App.Swagger.Version = Version
			// setup http server
			http := httpserver.HTTPSetting{
				Env:     cfg.App.Env,
				Conf:    cfg.App.HTTPServer,
				Swagger: cfg.App.Swagger,
				Log:     log,
			}

			gin := http.NewHTTPServer()
			http.SetSwaggo(docs.SwaggerInfo)

			// init http router
			restCfg := rest.RestDep{
				Conf:    cfg.Rest,
				Log:     &log,
				Usecase: uc,
				Gin:     gin,
			}
			handler := rest.New(&restCfg)

			restCfg.Serve(handler)
			http.Run()
		}()
	}

	<-readSignal

	log.Warn(context.Background(), "closing gracefully . . . ")
	st := time.Now()

	// close all connection here before shutdown
	psql.Close()
	redis.Close()
	grpcClient.Release()
	log.Error(context.Background(), "service shutdown!", time.Since(st).Seconds(), "sec")
}
