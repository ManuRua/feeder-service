package increase

import (
	"feeder-service/internal/shared/infra/counter"
)

type IncreaseDuplicatedProductUseCase struct {
	counter *counter.Counter
}

// NewIncreaseDuplicatedProductUseCase creates a new use case to increment duplicated products' counter
func NewIncreaseDuplicatedProductUseCase(counter *counter.Counter) IncreaseDuplicatedProductUseCase {
	return IncreaseDuplicatedProductUseCase{
		counter,
	}
}

// IncreaseDuplicatedProduct increments in one the count of duplicated products
func (s IncreaseDuplicatedProductUseCase) IncreaseDuplicatedProduct() {
	s.counter.Inc()
}
