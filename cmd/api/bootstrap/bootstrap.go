package bootstrap

import (
	"feeder-service/internal/products/application/create"
	"feeder-service/internal/products/application/increase"
	"feeder-service/internal/products/infra/server/handler"
	"feeder-service/internal/products/infra/storage/fs"
	"feeder-service/internal/products/infra/storage/in_memory"
	"feeder-service/internal/shared/domain/config"
	"feeder-service/internal/shared/infra/counter"
	server "feeder-service/internal/shared/infra/server/net"
)

const (
	network = "tcp"
	port    = 4000
	limit   = 5

	dbUser = "deporvillage"
	dbPass = "deporvillage"
	dbHost = "localhost"
	dbPort = "5432"
	dbName = "deporvillage"
)

func Run() error {
	serverConfig := config.ServerConfig{
		Network:   network,
		Port:      port,
		ConnLimit: limit,
	}

	persistRepository := fs.NewProductRepository()
	tempRepository := in_memory.NewProductRepository()

	invalidCounter := counter.Counter{}
	duplicatedCounter := counter.Counter{}

	createUC := create.NewCreateProductUseCase(persistRepository, tempRepository)
	increaseInvalidUC := increase.NewIncreaseInvalidProductUseCase(&invalidCounter)
	increaseDuplicatedUC := increase.NewIncreaseDuplicatedProductUseCase(&duplicatedCounter)

	handler := handler.NewCreateHandler(createUC, increaseInvalidUC, increaseDuplicatedUC)

	srv := server.New(serverConfig, handler)
	return srv.Run()
}
