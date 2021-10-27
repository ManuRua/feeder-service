package bootstrap

import (
	"feeder-service/internal/products/application/create"
	"feeder-service/internal/products/infra/server/handler"
	"feeder-service/internal/products/infra/storage/fs"
	"feeder-service/internal/shared/domain/config"
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

	repository := fs.NewProductRepository()
	createUC := create.NewCreateProductUseCase(repository)
	handler := handler.NewCreateHandler(createUC)

	srv := server.New(serverConfig, handler)
	return srv.Run()
}
