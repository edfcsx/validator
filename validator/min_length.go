package validator

import (
	"errors"
	"fmt"
)

type ValidateMinLength struct {
	value     string
	fieldName string
	minLength int
}

func MinLength(value, fieldName string, minLength int) *ValidateMinLength {
	return &ValidateMinLength{
		value:     value,
		fieldName: fieldName,
		minLength: minLength,
	}
}

func (h *ValidateMinLength) Handler() error {
	if len(h.value) < h.minLength {
		return errors.New(fmt.Sprintf("Field '%s' is less than %d", h.fieldName, h.minLength))
	}

	return nil
}
