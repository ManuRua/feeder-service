package increase

import (
	"feeder-service/internal/shared/infra/counter"
)

type IncreaseInvalidProductUseCase struct {
	counter *counter.Counter
}

func NewIncreaseInvalidProductUseCase() IncreaseInvalidProductUseCase {
	return IncreaseInvalidProductUseCase{
		counter: &counter.Counter{},
	}
}

func (s IncreaseInvalidProductUseCase) IncreaseInvalidProduct() {
	s.counter.Inc()
}
