package products

import "fmt"

type Product struct {
	sku string
}

func NewProduct(sku string) Product {
	return Product{
		sku: sku,
	}
}

func (p Product) SKU() string {
	return p.sku
}

func (p Product) String() string {
	return fmt.Sprintf("Product: [SKU] %s", p.sku)
}
