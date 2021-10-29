package handler

import (
	"bufio"
	"feeder-service/internal/products/application/create"
	"feeder-service/internal/products/application/increase"
	products "feeder-service/internal/products/domain"
	"fmt"
	"net"
)

type createHandler struct {
	createProductUseCase             create.CreateProductUseCase
	increaseInvalidProductUseCase    increase.IncreaseInvalidProductUseCase
	increaseDuplicatedProductUseCase increase.IncreaseDuplicatedProductUseCase
}

type Handler interface {
	Handle(c net.Conn, sem chan int, id *int)
}

func NewCreateHandler(
	createProductUseCase create.CreateProductUseCase,
	increaseInvalidProductUseCase increase.IncreaseInvalidProductUseCase,
	increaseDuplicatedProductUseCase increase.IncreaseDuplicatedProductUseCase,
) Handler {
	handler := &createHandler{
		createProductUseCase:             createProductUseCase,
		increaseInvalidProductUseCase:    increaseInvalidProductUseCase,
		increaseDuplicatedProductUseCase: increaseDuplicatedProductUseCase,
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
		case products.IsErrInvalidProductSKU(err):
			h.increaseInvalidProductUseCase.IncreaseInvalidProduct()
		case products.IsErrProductExists(err):
			h.increaseDuplicatedProductUseCase.IncreaseDuplicatedProduct()
		default:
			fmt.Println(err)
		}
	}

	h.Handle(c, sem, id)
}
