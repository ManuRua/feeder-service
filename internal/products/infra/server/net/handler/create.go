package handler

import (
	"bufio"
	"feeder-service/internal/products/application/create"
	"feeder-service/internal/products/application/increase"
	products "feeder-service/internal/products/domain"
	"feeder-service/internal/shared/infra/server/net/handler"
	"fmt"
	"net"
	"syscall"
)

type createHandler struct {
	createProductUseCase             create.CreateProductUseCase
	increaseInvalidProductUseCase    increase.IncreaseInvalidProductUseCase
	increaseDuplicatedProductUseCase increase.IncreaseDuplicatedProductUseCase
}

// NewCreateHandler creates a new handler to create a product
func NewCreateHandler(
	createProductUseCase create.CreateProductUseCase,
	increaseInvalidProductUseCase increase.IncreaseInvalidProductUseCase,
	increaseDuplicatedProductUseCase increase.IncreaseDuplicatedProductUseCase,
) handler.Handler {
	handler := &createHandler{
		createProductUseCase,
		increaseInvalidProductUseCase,
		increaseDuplicatedProductUseCase,
	}

	return handler
}

// Handle receives string input and manages validity of this
func (h *createHandler) Handle(c net.Conn) {
	str, err := bufio.NewReader(c).ReadString('\n')
	if err != nil {
		fmt.Println(err)
		return
	}

	input := str[:len(str)-1]
	if input == "terminate" {
		err = syscall.Kill(syscall.Getpid(), syscall.SIGINT)
		if err != nil {
			fmt.Println(err)
			return
		}
	}

	err = h.createProductUseCase.CreateProduct(str[:len(str)-1])

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
}
