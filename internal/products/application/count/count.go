package count

import (
	products "feeder-service/internal/products/domain"
	"feeder-service/internal/shared/infra/counter"
)

type CountProducts struct {
	Uniques    int
	Invalids   int
	Duplicated int
}

type CountProductsUseCase struct {
	tempProductRepository products.ProductRepository
	invalidCounter        *counter.Counter
	duplicatedCounter     *counter.Counter
}

// NewCreateHandler creates a new use case to count all processed products
func NewCountProductsUseCase(
	tempProductRepository products.ProductRepository,
	invalidCounter *counter.Counter,
	duplicatedCounter *counter.Counter,
) CountProductsUseCase {
	return CountProductsUseCase{
		tempProductRepository,
		invalidCounter,
		duplicatedCounter,
	}
}

// CountProducts returns total count of valid, invalid and duplicated products
func (s CountProductsUseCase) CountProducts() CountProducts {
	uniques := s.tempProductRepository.Count()
	invalids := int(s.invalidCounter.Value())
	duplicated := int(s.duplicatedCounter.Value())

	return CountProducts{
		uniques,
		invalids,
		duplicated,
	}
}
