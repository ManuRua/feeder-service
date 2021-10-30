package create_handler

import (
	"bufio"
	"feeder-service/internal/products/application/create"
	"feeder-service/internal/products/application/increase"
	products "feeder-service/internal/products/domain"
	"feeder-service/internal/shared/infra/server/net/handler"
	"fmt"
	"net"
)

type createHandler struct {
	createProductUseCase             create.CreateProductUseCase
	increaseInvalidProductUseCase    increase.IncreaseInvalidProductUseCase
	increaseDuplicatedProductUseCase increase.IncreaseDuplicatedProductUseCase
}

func NewCreateHandler(
	createProductUseCase create.CreateProductUseCase,
	increaseInvalidProductUseCase increase.IncreaseInvalidProductUseCase,
	increaseDuplicatedProductUseCase increase.IncreaseDuplicatedProductUseCase,
) handler.Handler {
	handler := &createHandler{
		createProductUseCase:             createProductUseCase,
		increaseInvalidProductUseCase:    increaseInvalidProductUseCase,
		increaseDuplicatedProductUseCase: increaseDuplicatedProductUseCase,
	}

	return handler
}

func (h *createHandler) Handle(c net.Conn) {
	str, err := bufio.NewReader(c).ReadString('\n')
	if err != nil {
		fmt.Println(err)
		return
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