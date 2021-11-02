package fs

import (
	products "feeder-service/internal/products/domain"
	"testing"
)

func TestCount(t *testing.T) {
	product, _ := products.NewProduct("KASL-3423")

	repo := NewProductRepository()

	empty := repo.Count()
	if empty != 0 {
		t.Errorf("Expected: %v, got: %v", 0, empty)
	}

	repo.Save(&product)

	empty = repo.Count()
	if empty != 0 {
		t.Errorf("Expected: %v, got: %v", 0, empty)
	}
}
