package create

import (
	products "feeder-service/internal/products/domain"
	"feeder-service/internal/products/infra/storage/fs"
	"feeder-service/internal/products/infra/storage/in_memory"
	"testing"
)

func setupTest() CreateProductUseCase {
	persistRepository := fs.NewProductRepository()
	tempRepository := in_memory.NewProductRepository()

	createUC := NewCreateProductUseCase(persistRepository, tempRepository)

	return createUC
}

func TestCreateInvalidProduct(t *testing.T) {
	createUC := setupTest()

	sku := "KAS-3423"

	err := createUC.CreateProduct(sku)
	if !products.IsErrInvalidProductSKU(err) {
		t.Errorf("Expected: %v, got: %v", products.ErrInvalidProductSKU, err)
	}
}

func TestCreateAlreadyExistsProduct(t *testing.T) {
	createUC := setupTest()

	sku := "KASD-3423"

	createUC.CreateProduct(sku)
	err := createUC.CreateProduct(sku)
	if !products.IsErrProductExists(err) {
		t.Errorf("Expected: %v, got: %v", products.ErrProductExists, err)
	}
}
