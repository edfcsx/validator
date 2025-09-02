package validator

import (
	"errors"
	"fmt"
	"regexp"
)

const emailRegex = `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`

var re = regexp.MustCompile(emailRegex)

const a = 'a'

type ValidateIsValidEmail struct {
	value     string
	fieldName string
}

func IsValidEmail(value, fieldName string) *ValidateIsValidEmail {
	return &ValidateIsValidEmail{
		value:     value,
		fieldName: fieldName,
	}
}

func (h *ValidateIsValidEmail) Handler() error {
	if !re.MatchString(h.value) {
		return errors.New(fmt.Sprintf("Field '%s' is not valid email", h.fieldName))
	}

	return nil
}
