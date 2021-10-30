package handler

import (
	"net"
)

type Handler interface {
	Handle(c net.Conn)
}
