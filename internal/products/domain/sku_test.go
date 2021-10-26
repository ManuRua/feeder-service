package products

import (
	"testing"
)

var newSKUTests = []struct {
	sku  string
	pass bool
}{
	{"KASL-3423", true},
	{"LPOS-32411", false},
	{"ABCD1234-", false},
	{"-ABCD1234", false},
	{"1234-ABCD", false},
}

func TestNewSKU(t *testing.T) {
	for _, tt := range newSKUTests {
		t.Run(tt.sku, func(t *testing.T) {
			expected := SKU(tt.sku)

			sku, err := NewSKU(tt.sku)
			if sku != nil && *sku != expected {
				t.Errorf("Expected: %v, got: %v", expected, *sku)
			}

			if tt.pass && err != nil {
				t.Errorf("Pass: %t, got: %v", tt.pass, err)
			}
			if !tt.pass && !IsNotValidSKUError(err) {
				t.Errorf("Pass: %t, got: %v", tt.pass, err)
			}
		})
	}
}
