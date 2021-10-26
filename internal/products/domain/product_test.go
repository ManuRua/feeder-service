package products

import "testing"

func TestNewProduct(t *testing.T) {
	sku := SKU("KASL-3423")
	expected := Product{
		sku: sku,
	}

	product := NewProduct(sku)
	if product != expected {
		t.Errorf("Expected: %v, got: %v", expected, product)
	}

	if product.SKU() != sku {
		t.Errorf("Expected: %s, got: %s", product.SKU(), sku)
	}
}
