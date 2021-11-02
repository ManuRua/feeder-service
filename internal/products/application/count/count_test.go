package count

import (
	products "feeder-service/internal/products/domain"
	"feeder-service/internal/products/infra/storage/in_memory"
	"feeder-service/internal/shared/infra/counter"
	"testing"
)

func setupTest() (CountProductsUseCase, products.ProductRepository, *counter.Counter, *counter.Counter) {
	tempRepository := in_memory.NewProductRepository()
	invalidCounter := &counter.Counter{}
	duplicatedCounter := &counter.Counter{}

	countUC := NewCountProductsUseCase(tempRepository, invalidCounter, duplicatedCounter)

	return countUC, tempRepository, invalidCounter, duplicatedCounter
}

func TestCountProductsEmpty(t *testing.T) {
	countUC, _, _, _ := setupTest()

	expected := CountProducts{
		Uniques:    0,
		Invalids:   0,
		Duplicated: 0,
	}

	counts := countUC.CountProducts()
	if counts != expected {
		t.Errorf("Expected: %v, got: %v", expected, counts)
	}
}

func TestCountProductsWithValue(t *testing.T) {
	countUC, tempRepository, invalidCounter, duplicatedCounter := setupTest()

	product, _ := products.NewProduct("KASL-3423")

	tempRepository.Save(&product)
	invalidCounter.Inc()
	duplicatedCounter.Inc()

	expected := CountProducts{
		Uniques:    1,
		Invalids:   1,
		Duplicated: 1,
	}

	counts := countUC.CountProducts()
	if counts != expected {
		t.Errorf("Expected: %v, got: %v", expected, counts)
	}
}
