package server

import (
	"bufio"
	"fmt"
	"log"
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
	buffer, err := bufio.NewReader(c).ReadBytes('\n')
    if err != nil {
        fmt.Println("Client left.")
        c.Close()
        return
    }

    log.Println("Client message:", string(buffer[:len(buffer)-1]))

    h.Handle(c)
}