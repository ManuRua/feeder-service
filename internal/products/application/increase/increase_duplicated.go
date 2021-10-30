package increase

import (
	"feeder-service/internal/shared/infra/counter"
)

type IncreaseDuplicatedProductUseCase struct {
	counter *counter.Counter
}

func NewIncreaseDuplicatedProductUseCase() IncreaseDuplicatedProductUseCase {
	return IncreaseDuplicatedProductUseCase{
		counter: &counter.Counter{},
	}
}

func (s IncreaseDuplicatedProductUseCase) IncreaseDuplicatedProduct() {
	s.counter.Inc()
}
