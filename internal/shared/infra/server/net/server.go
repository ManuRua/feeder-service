package server

import (
	"feeder-service/internal/products/infra/server/handler"
	"feeder-service/internal/shared/domain/config"
	"fmt"
	"log"
	"net"
	"strconv"
)

type server struct {
	config  config.ServerConfig
	handler handler.Handler
}

type Server interface {
	Run() error
}

func New(serverConfig config.ServerConfig, handler handler.Handler) Server {
	srv := &server{
		config:  serverConfig,
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

	sem := make(chan int, s.config.ConnLimit)
	id := 1

	for {
		c, err := l.Accept()
		if err != nil {
			fmt.Println("Error connecting:", err.Error())
			return err
		}
		if id > s.config.ConnLimit {
			_, err = c.Write([]byte("There are already 5 clients connected."))
			if err != nil {
				fmt.Println("Error writting:", err.Error())
			}
			c.Close()
		} else {
			sem <- id

			fmt.Println("Client " + c.RemoteAddr().String() + " connected")
			go s.handler.Handle(c, sem, &id)

			id++
		}
	}
}
