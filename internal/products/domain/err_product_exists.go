package products

import (
	"errors"
	"fmt"
)

var ErrProductExists = errors.New("Product already exists")

type errProductExists struct {
	error
}

func NewErrProductExists(product *Product) error {
	return &errProductExists{fmt.Errorf("%w: %v", ErrProductExists, *product)}
}

func IsErrProductExists(err error) bool {
	return errors.Is(err, errProductExists{})
}

func (e errProductExists) Is(target error) bool { return target == errProductExists{} }
