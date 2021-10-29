package increase

import (
	"feeder-service/internal/shared/infra/counter"
)

type IncreaseDuplicatedProductUseCase struct {
	counter *counter.Counter
}

func NewIncreaseDuplicatedProductUseCase(
	counter *counter.Counter,
) IncreaseDuplicatedProductUseCase {
	return IncreaseDuplicatedProductUseCase{
		counter: counter,
	}
}

func (s IncreaseDuplicatedProductUseCase) IncreaseDuplicatedProduct() {
	s.counter.Inc()
}
