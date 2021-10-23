package server

import (
	"feeder-service/internal/shared/domain/config"
	"fmt"
	"log"
	"net"
	"strconv"
)

type server struct {
	config config.ServerConfig
	handler Handler
}

type Server interface{
	Run() error
}

func New(serverConfig config.ServerConfig, handler Handler) Server {
	srv := &server{
		config: serverConfig,
		handler: handler,
	}

	return srv
}

func (s *server) Run() error {
	portStr := strconv.Itoa(s.config.Port)

	l, err := net.Listen(s.config.Network, ":"+portStr)
	if err != nil {
		fmt.Println(err)
		return err
	}
	defer l.Close()

	log.Println("Server running on port", portStr)

	for {
		c, err := l.Accept()
		if err != nil {
			fmt.Println("Error connecting:", err.Error())
			return err
		}
		fmt.Println("Client connected")

		s.handler.Handle(c)
    }
}
