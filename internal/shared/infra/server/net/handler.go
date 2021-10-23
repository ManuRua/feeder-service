package server

import (
	"fmt"
	"net"
)

type handler struct {}

type Handler interface {
	Handle(c net.Conn)
}

func NewHandler() Handler {
	handler := &handler{}

	return handler
}

func (h *handler) Handle(c net.Conn) {
	fmt.Println("Client " + c.RemoteAddr().String() + " connected")
}