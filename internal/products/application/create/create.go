package create

import (
	products "feeder-service/internal/products/domain"
)

type CreateProductUseCase struct {
	persistProductRepository products.ProductRepository
	tempProductRepository    products.ProductRepository
}

func NewCreateProductUseCase(
	persistProductRepository products.ProductRepository,
	tempProductRepository products.ProductRepository,
) CreateProductUseCase {
	return CreateProductUseCase{
		persistProductRepository: persistProductRepository,
		tempProductRepository:    tempProductRepository,
	}
}

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
