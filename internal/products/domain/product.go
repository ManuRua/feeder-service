package products

import "fmt"

type Product struct {
	sku SKU
}

func NewProduct(sku SKU) Product {
	return Product{
		sku: sku,
	}
}

func (p Product) SKU() SKU {
	return p.sku
}

func (p Product) String() string {
	return fmt.Sprintf("Product: [SKU] %s", p.sku)
}
