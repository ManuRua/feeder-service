package in_memory

import (
	"errors"
	products "feeder-service/internal/products/domain"
	"fmt"
	"sync"
)

var ErrProductExists = errors.New("Product already exists")

type productRepository struct {
	mu  sync.Mutex
	set map[products.Product]bool
}

// NewProductRepository init a new in memory product repository.
func NewProductRepository(set map[products.Product]bool) products.ProductRepository {
	return &productRepository{
		set: set,
	}
}

// Save add entry to map of products if not exists yet.
func (r *productRepository) Save(product *products.Product) error {
	r.mu.Lock()

	if r.set[*product] {
		return fmt.Errorf("%w: %v", ErrProductExists, *product)
	}
	r.set[*product] = true

	r.mu.Unlock()

	return nil
}
