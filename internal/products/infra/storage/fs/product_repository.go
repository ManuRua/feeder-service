package fs

import (
	"errors"
	products "feeder-service/internal/products/domain"
	"fmt"
	"os"
)

const (
	fsProductFilename = "tmp/products.log"
)

var ErrCreateProduct = errors.New("Error trying to persist product on log file")

type productRepository struct {
}

// NewProductRepository init a new filesystem product repository.
func NewProductRepository() products.ProductRepository {
	return &productRepository{}
}

// Save add entry to log file of products.
func (r *productRepository) Save(product *products.Product) error {
	f, err := os.OpenFile(fsProductFilename, os.O_CREATE|os.O_APPEND|os.O_RDWR, 0644)
	if err != nil {
		return fmt.Errorf("%w: %v", ErrCreateProduct, err)
	}
	defer f.Close()

	_, err = f.WriteString(product.SKU().String() + "\n")
	if err != nil {
		return fmt.Errorf("%w: %v", ErrCreateProduct, err)
	}

	err = f.Sync()
	if err != nil {
		return fmt.Errorf("%w: %v", ErrCreateProduct, err)
	}

	return nil
}

// Count returns the length of stored products
func (r *productRepository) Count() int {
	return 0
}
