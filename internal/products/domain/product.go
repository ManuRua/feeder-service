package products

import "fmt"

type Product struct {
	sku ProductSKU
}

func NewProduct(sku string) (Product, error) {
	skuVO, err := NewProductSKU(sku)
	if err != nil {
		return Product{}, err
	}

	return Product{
		sku: skuVO,
	}, nil
}

func (p Product) SKU() ProductSKU {
	return p.sku
}

func (p Product) String() string {
	return fmt.Sprintf("Product: [SKU] %s", p.sku)
}
