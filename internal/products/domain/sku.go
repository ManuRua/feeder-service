package products

import (
	"regexp"
)

type SKU string

var rxPat = regexp.MustCompile(`^[A-Z]{4}-[0-9]{4}$`)

func NewSKU(sku string) (*SKU, error) {
	err := ensureIsValidSKU(sku)
	if err != nil {
		return nil, err
	}

	skuVO := SKU(sku)

	return &skuVO, nil
}

func ensureIsValidSKU(sku string) error {
	if !rxPat.MatchString(sku) {
		return NewNotValidSKUError()
	} else {
		return nil
	}
}
