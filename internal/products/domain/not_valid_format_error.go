package products

import "errors"

type notValidSKUError struct {
	error
}

func NewNotValidSKUError() error {
	return &notValidSKUError{errors.New("SKU has invalid format.")}
}

func IsNotValidSKUError(err error) bool {
	return errors.Is(err, notValidSKUError{})
}

func (e notValidSKUError) Is(target error) bool { return target == notValidSKUError{} }
