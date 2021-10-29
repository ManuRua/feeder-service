package in_memory

import (
	products "feeder-service/internal/products/domain"
	"sync"
)

type productRepository struct {
	mu  sync.Mutex
	set map[products.Product]bool
}

// NewProductRepository init a new in memory product repository.
func NewProductRepository() products.ProductRepository {
	return &productRepository{
		set: make(map[products.Product]bool),
	}
}

// Save add entry to map of products if not exists yet.
func (r *productRepository) Save(product *products.Product) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	if r.set[*product] {
		return products.NewErrProductExists(product)
	}
	r.set[*product] = true

	return nil
}
