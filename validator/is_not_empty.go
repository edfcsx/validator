package validator

import (
	"errors"
	"fmt"
)

type ValidateIsNotEmpty struct {
	value     string
	fieldName string
}

func IsNotEmpty(value string, fieldName string) *ValidateIsNotEmpty {
	return &ValidateIsNotEmpty{
		value:     value,
		fieldName: fieldName,
	}
}

func (h *ValidateIsNotEmpty) Handler() error {
	if len(h.value) == 0 {
		return errors.New(fmt.Sprintf("Field '%s' is required and cannot be empty", h.fieldName))
	}

	return nil
}
