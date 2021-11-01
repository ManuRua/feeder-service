package handler

import (
	"net"
)

// Handler defines a net connection handler
type Handler interface {
	Handle(c net.Conn)
}
