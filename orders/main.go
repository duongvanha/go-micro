package main

import (
	"context"
	"github.com/go-micro/plugins/v4/client/grpc"
	"go-micro.dev/v4/logger"
	"proto"
	"proto/services"
)

func main() {

	//service := micro.NewService(micro.Name("product-service"), micro.Server())
	//service.Init()

	// Request message
	//req := service.Client().NewRequest("serviceA", "Greeter.Hello", &map[string]interface{}{}, client.WithContentType("application/json"))

	productService := services.NewProductsService(proto.ProductServiceName, grpc.NewClient())

	res, err := productService.Call(context.Background(), &services.CallRequest{
		Name: "Apple",
	})

	if err != nil {
		logger.Fatal(err)
	}

	logger.Info(res.Msg)
}
