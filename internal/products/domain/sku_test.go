package products

import (
	"testing"
)

var newProductSKUTests = []struct {
	sku  string
	pass bool
}{
	{"KASL-3423", true},
	{"LPOS-32411", false},
	{"ABCD1234-", false},
	{"-ABCD1234", false},
	{"1234-ABCD", false},
}

func TestNewProductSKU(t *testing.T) {
	for _, tt := range newProductSKUTests {
		t.Run(tt.sku, func(t *testing.T) {
			expected := ProductSKU{
				value: tt.sku,
			}

			sku, err := NewProductSKU(tt.sku)
			if err == nil && sku != expected {
				t.Errorf("Expected: %v, got: %v", expected, sku)
			}

			if tt.pass && err != nil {
				t.Errorf("Pass: %t, got: %v", tt.pass, err)
			}
			if !tt.pass && !IsErrInvalidProductSKU(err) {
				t.Errorf("Pass: %t, got: %v", tt.pass, err)
			}
		})
	}
}
