package create

import products "feeder-service/internal/products/domain"

type CreateProductUseCase struct {
	productRepository products.ProductRepository
}

func NewCreateProductUseCase(productRepository products.ProductRepository) CreateProductUseCase {
	return CreateProductUseCase{
		productRepository: productRepository,
	}
}

func (s CreateProductUseCase) CreateProduct(sku string) error {
	product, err := products.NewProduct(sku)
	if err != nil {
		return err
	}

	return s.productRepository.Save(&product)
}
