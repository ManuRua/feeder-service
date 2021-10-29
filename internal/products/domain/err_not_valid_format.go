package products

import (
	"errors"
	"fmt"
)

var ErrInvalidProductSKU = errors.New("Invalid Product SKU")

type errInvalidProductSKU struct {
	error
}

func NewErrInvalidProductSKU(value string) error {
	return &errInvalidProductSKU{fmt.Errorf("%w: %s", ErrInvalidProductSKU, value)}
}

func IsErrInvalidProductSKU(err error) bool {
	return errors.Is(err, errInvalidProductSKU{})
}

func (e errInvalidProductSKU) Is(target error) bool { return target == errInvalidProductSKU{} }
