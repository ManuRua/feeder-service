package increase

import (
	"feeder-service/internal/shared/infra/counter"
	"testing"
)

func TestIncreaseDuplicatedProduct(t *testing.T) {
	duplicatedCounter := &counter.Counter{}
	increaseDuplicatedUC := NewIncreaseDuplicatedProductUseCase(duplicatedCounter)

	if v := duplicatedCounter.Value(); v != 0 {
		t.Errorf("Expected: %v, got: %v", 0, duplicatedCounter.Value())
	}

	increaseDuplicatedUC.IncreaseDuplicatedProduct()

	if v := duplicatedCounter.Value(); v != 1 {
		t.Errorf("Expected: %v, got: %v", 1, duplicatedCounter.Value())
	}
}
