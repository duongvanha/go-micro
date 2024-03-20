package app

import (
	"github.com/go-micro/plugins/v4/server/grpc"
	"github.com/samber/do"
	"go-micro.dev/v4"
	"go-micro.dev/v4/logger"
	"go-micro.dev/v4/server"
	"products/app/factory"
	"products/handler"
	"proto"
	"proto/services"
)

var (
	version = "latest"
)

func GRPC(options ...micro.Option) micro.Service {
	container, err := factory.Init()
	if err != nil {
		logger.Fatal("Error initializing registry: ", err)
	}

	srv := micro.NewService(micro.Server(grpc.NewServer(
		server.Name(proto.ProductServiceName),
	)))

	productHandler, err := do.Invoke[*handler.Products](container)
	if err != nil {
		logger.Fatal("Error invoking product handler: ", err)
	}

	err = services.RegisterProductsHandler(srv.Server(), productHandler)

	if err != nil {
		logger.Fatal("Error registering grpc handler: ", err)
	}

	return srv
}
