package main

import (
	"go-micro.dev/v4/logger"
	"products/app"
)

func main() {
	if err := app.GRPC().Run(); err != nil {
		logger.Fatal(err)
	}
}
