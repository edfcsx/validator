package validator

import (
	"errors"
	"fmt"
)

type ValidateMaxLength struct {
	value     string
	fieldName string
	maxLength int
}

func MaxLength(value, fieldName string, maxLength int) *ValidateMaxLength {
	return &ValidateMaxLength{
		value:     value,
		fieldName: fieldName,
		maxLength: maxLength,
	}
}

func (h *ValidateMaxLength) Handler() error {
	if len(h.value) > h.maxLength {
		return errors.New(fmt.Sprintf("Field '%s' is greater than %d", h.fieldName, h.maxLength))
	}

	return nil
}
