package factory

import (
	"github.com/samber/do"
	"go-micro.dev/v4/cache"
	"products/handler"
)

func Init() (*do.Injector, error) {
	container := do.New()

	do.Provide(container, func(injector *do.Injector) (cache.Cache, error) {
		// provide default is memory adapter
		return cache.NewCache(), nil
	})

	do.Provide(container, handler.NewProducts)

	return container, nil
}
