package in_memory

import (
	products "feeder-service/internal/products/domain"
	"testing"
)

func TestSave(t *testing.T) {
	product, _ := products.NewProduct("KASL-3423")

	repo := NewProductRepository()

	err := repo.Save(&product)
	if err != nil {
		t.Errorf("Expected: %v, got: %v", nil, err)
	}
}

func TestSaveErrProductExists(t *testing.T) {
	product, _ := products.NewProduct("KASL-3423")

	repo := NewProductRepository()

	repo.Save(&product)

	err := repo.Save(&product)
	if !products.IsErrProductExists(err) {
		t.Errorf("Expected: %v, got: %v", products.ErrProductExists, err)
	}
}

func TestCount(t *testing.T) {
	product, _ := products.NewProduct("KASL-3423")

	repo := NewProductRepository()

	empty := repo.Count()
	if empty != 0 {
		t.Errorf("Expected: %v, got: %v", 0, empty)
	}

	repo.Save(&product)

	one := repo.Count()
	if one != 1 {
		t.Errorf("Expected: %v, got: %v", 1, one)
	}
}
