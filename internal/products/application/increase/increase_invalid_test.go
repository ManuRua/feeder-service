package increase

import (
	"feeder-service/internal/shared/infra/counter"
	"testing"
)

func TestIncreaseInvalidProduct(t *testing.T) {
	invalidCounter := &counter.Counter{}
	increaseInvalidUC := NewIncreaseInvalidProductUseCase(invalidCounter)

	if v := invalidCounter.Value(); v != 0 {
		t.Errorf("Expected: %v, got: %v", 0, invalidCounter.Value())
	}

	increaseInvalidUC.IncreaseInvalidProduct()

	if v := invalidCounter.Value(); v != 1 {
		t.Errorf("Expected: %v, got: %v", 1, invalidCounter.Value())
	}
}
