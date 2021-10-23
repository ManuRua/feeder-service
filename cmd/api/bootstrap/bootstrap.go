package bootstrap

import (
	"feeder-service/internal/shared/domain/config"
	server "feeder-service/internal/shared/infra/server/net"
)

const (
	network = "tcp"
	port = 4000

	dbUser = "deporvillage"
	dbPass = "deporvillage"
	dbHost = "localhost"
	dbPort = "5432"
	dbName = "deporvillage"
)

func Run() error {
	serverConfig := config.ServerConfig{
		Network: network,
		Port: port,
	}

	srv := server.New(serverConfig)
	return srv.Run()
}