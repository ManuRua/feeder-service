package products

import (
	"regexp"
)

type ProductSKU struct {
	value string
}

var rxPat = regexp.MustCompile(`^[A-Z]{4}-[0-9]{4}$`)

func NewProductSKU(value string) (ProductSKU, error) {
	err := ensureIsValidProductSKU(value)
	if err != nil {
		return ProductSKU{}, err
	}

	return ProductSKU{
		value: value,
	}, nil
}

func (sku ProductSKU) String() string {
	return sku.value
}

func ensureIsValidProductSKU(value string) error {
	if !rxPat.MatchString(value) {
		return NewErrInvalidProductSKU(value)
	} else {
		return nil
	}
}
