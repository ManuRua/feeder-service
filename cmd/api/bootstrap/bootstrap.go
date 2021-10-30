package bootstrap

import (
	"context"
	"feeder-service/internal/products/application/create"
	"feeder-service/internal/products/application/increase"
	create_handler "feeder-service/internal/products/infra/server/net/handler"
	"feeder-service/internal/products/infra/storage/fs"
	"feeder-service/internal/products/infra/storage/in_memory"
	"feeder-service/internal/shared/domain/config"
	server "feeder-service/internal/shared/infra/server/net"
)

const (
	network   = "tcp"
	port      = 4000
	connLimit = 5
	timeout   = 30
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

	createUC := create.NewCreateProductUseCase(persistRepository, tempRepository)
	increaseInvalidUC := increase.NewIncreaseInvalidProductUseCase()
	increaseDuplicatedUC := increase.NewIncreaseDuplicatedProductUseCase()

	handler := create_handler.NewCreateHandler(createUC, increaseInvalidUC, increaseDuplicatedUC)

	ctx, srv := server.New(context.Background(), serverConfig, handler)

	defer srv.Shutdown()
	return srv.Run(ctx)
}
