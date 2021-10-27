package products

import (
	"testing"
)

func TestNewProduct(t *testing.T) {
	sku := ProductSKU{
		value: "KASL-3423",
	}
	expected := Product{
		sku: sku,
	}

	product, err := NewProduct(sku.value)
	if err != nil || product != expected {
		t.Errorf("Expected: %v, got: %v", expected, product)
	}

	if product.SKU() != sku {
		t.Errorf("Expected: %s, got: %s", product.SKU(), sku)
	}
}
