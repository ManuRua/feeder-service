package server

import (
	"feeder-service/internal/shared/domain/config"
	"fmt"
	"log"
	"net"
	"strconv"
)

type Server struct {
	config config.ServerConfig
}

func New(serverConfig config.ServerConfig) Server {
	srv := Server{
		config: serverConfig,
	}

	return srv
}

func (s *Server) Run() error {
	portStr := strconv.Itoa(s.config.Port)
	l, err := net.Listen(s.config.Network, ":"+portStr)
	if err != nil {
		fmt.Println(err)
		return err
	}
	defer l.Close()

	log.Println("Server running on port", portStr)

	for {
    }
}
