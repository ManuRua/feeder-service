package bootstrap

import (
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

	handler := server.NewHandler()

	srv := server.New(serverConfig, handler)
	return srv.Run()
}
