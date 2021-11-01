package bootstrap

import (
	"context"
	"feeder-service/internal/products/application/count"
	"feeder-service/internal/products/application/create"
	"feeder-service/internal/products/application/increase"
	"feeder-service/internal/products/infra/server/net/handler"
	"feeder-service/internal/products/infra/storage/fs"
	"feeder-service/internal/products/infra/storage/in_memory"
	"feeder-service/internal/shared/domain/config"
	"feeder-service/internal/shared/infra/counter"
	server "feeder-service/internal/shared/infra/server/net"
)

const (
	network   = "tcp"
	port      = 4000
	connLimit = 5
	timeout   = 10
)

// Run manages dependency injection and launch server
func Run() error {
	serverConfig := config.ServerConfig{
		Network:   network,
		Port:      port,
		ConnLimit: connLimit,
		Timeout:   timeout,
	}

	persistRepository := fs.NewProductRepository()
	tempRepository := in_memory.NewProductRepository()

	invalidCounter := &counter.Counter{}
	duplicatedCounter := &counter.Counter{}

	createUC := create.NewCreateProductUseCase(persistRepository, tempRepository)
	increaseInvalidUC := increase.NewIncreaseInvalidProductUseCase(invalidCounter)
	increaseDuplicatedUC := increase.NewIncreaseDuplicatedProductUseCase(duplicatedCounter)
	countUC := count.NewCountProductsUseCase(tempRepository, invalidCounter, duplicatedCounter)

	createHandler := handler.NewCreateHandler(createUC, increaseInvalidUC, increaseDuplicatedUC)
	reportHandler := handler.NewReportHandler(countUC)

	ctx, srv := server.New(context.Background(), serverConfig, createHandler, reportHandler)

	defer srv.Shutdown()
	return srv.Run(ctx)
}
