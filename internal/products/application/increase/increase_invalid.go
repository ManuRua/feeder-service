package increase

import (
	"feeder-service/internal/shared/infra/counter"
)

type IncreaseInvalidProductUseCase struct {
	counter *counter.Counter
}

func NewIncreaseInvalidProductUseCase(
	counter *counter.Counter,
) IncreaseInvalidProductUseCase {
	return IncreaseInvalidProductUseCase{
		counter: counter,
	}
}

func (s IncreaseInvalidProductUseCase) IncreaseInvalidProduct() {
	s.counter.Inc()
}
