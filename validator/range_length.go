package validator

type ValidateRangeLength struct {
	value     string
	fieldName string
	minLength int
	maxLength int
}

func RangeLength(value, fieldName string, minLength, maxLength int) *ValidateRangeLength {
	return &ValidateRangeLength{
		value:     value,
		fieldName: fieldName,
		minLength: minLength,
		maxLength: maxLength,
	}
}

func (u *ValidateRangeLength) Handler() error {
	err := MinLength(u.value, u.fieldName, u.minLength).Handler()

	if err != nil {
		return err
	}

	err = MaxLength(u.value, u.fieldName, u.maxLength).Handler()

	if err != nil {
		return err
	}

	return nil
}
