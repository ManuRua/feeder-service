package increase

import (
	"feeder-service/internal/shared/infra/counter"
)

type IncreaseInvalidProductUseCase struct {
	counter *counter.Counter
}

// NewIncreaseDuplicatedProductUseCase creates a new use case to increment invalid products' counter
func NewIncreaseInvalidProductUseCase(counter *counter.Counter) IncreaseInvalidProductUseCase {
	return IncreaseInvalidProductUseCase{
		counter,
	}
}

// IncreaseDuplicatedProduct increments in one the count of invalid products
func (s IncreaseInvalidProductUseCase) IncreaseInvalidProduct() {
	s.counter.Inc()
}
