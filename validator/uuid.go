package validator

import (
	"errors"
	"fmt"

	"github.com/google/uuid"
)

type ValidateIsValidUUID struct {
	value     string
	fieldName string
}

func IsValidUUID(value, fieldName string) *ValidateIsValidUUID {
	return &ValidateIsValidUUID{
		value:     value,
		fieldName: fieldName,
	}
}

func (h *ValidateIsValidUUID) Handler() error {
	if _, err := uuid.Parse(h.value); err != nil {
		return errors.New(fmt.Sprintf("Field '%s' is not valid UUID", h.fieldName))
	}

	return nil
}
