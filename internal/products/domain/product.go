package products

import "fmt"

// Product is the aggregate that represents a product in the store
type Product struct {
	sku ProductSKU
}

// NewProductSKU creates a valid Product
func NewProduct(sku string) (Product, error) {
	skuVO, err := NewProductSKU(sku)
	if err != nil {
		return Product{}, err
	}

	return Product{
		sku: skuVO,
	}, nil
}

// String gets the SKU of product
func (p Product) SKU() ProductSKU {
	return p.sku
}

func (p Product) String() string {
	return fmt.Sprintf("Product: [SKU] %s", p.sku)
}
