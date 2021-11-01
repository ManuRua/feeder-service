package create

import (
	products "feeder-service/internal/products/domain"
)

type CreateProductUseCase struct {
	persistProductRepository products.ProductRepository
	tempProductRepository    products.ProductRepository
}

// NewCreateHandler creates a new use case to create a product
func NewCreateProductUseCase(
	persistProductRepository products.ProductRepository,
	tempProductRepository products.ProductRepository,
) CreateProductUseCase {
	return CreateProductUseCase{
		persistProductRepository,
		tempProductRepository,
	}
}

// CreateProduct creates a valid product and save it properly
func (s CreateProductUseCase) CreateProduct(sku string) error {
	product, err := products.NewProduct(sku)
	if err != nil {
		return err
	}

	err = s.tempProductRepository.Save(&product)
	if err != nil {
		return err
	}

	return s.persistProductRepository.Save(&product)
}
