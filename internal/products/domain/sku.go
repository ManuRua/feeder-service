package products

import (
	"regexp"
)

// ProductSKU is the value object that represents unique identifier of a product
type ProductSKU struct {
	value string
}

var rxPat = regexp.MustCompile(`^[A-Z]{4}-[0-9]{4}$`)

// NewProductSKU creates a valid product SKU
func NewProductSKU(value string) (ProductSKU, error) {
	err := ensureIsValidProductSKU(value)
	if err != nil {
		return ProductSKU{}, err
	}

	return ProductSKU{
		value: value,
	}, nil
}

// String gets the value of product SKU
func (sku ProductSKU) String() string {
	return sku.value
}

// ensureIsValidProductSKU returns an error if value has not valid format
func ensureIsValidProductSKU(value string) error {
	if !rxPat.MatchString(value) {
		return NewErrInvalidProductSKU(value)
	} else {
		return nil
	}
}
