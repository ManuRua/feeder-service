package handler

import (
	"bufio"
	"errors"
	"feeder-service/internal/products/application/create"
	products "feeder-service/internal/products/domain"
	"fmt"
	"net"
)

type createHandler struct {
	createProductUseCase create.CreateProductUseCase
}

type Handler interface {
	Handle(c net.Conn, sem chan int, id *int)
}

func NewCreateHandler(createProductUseCase create.CreateProductUseCase) Handler {
	handler := &createHandler{
		createProductUseCase: createProductUseCase,
	}

	return handler
}

func (h *createHandler) Handle(c net.Conn, sem chan int, id *int) {
	buffer, err := bufio.NewReader(c).ReadBytes('\n')
	if err != nil {
		fmt.Println("Client left.")
		c.Close()
		<-sem
		(*id)--
		return
	}

	err = h.createProductUseCase.CreateProduct(string(buffer[:len(buffer)-1]))
	if err != nil {
		switch {
		case errors.Is(err, products.ErrInvalidProductSKU):
			// TODO: Add to invalid SKU counter
		default:
		}
	}

	h.Handle(c, sem, id)
}
