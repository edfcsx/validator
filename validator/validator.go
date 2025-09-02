package validator

type ValidateFunction interface {
	Handler() error
}

func Validate(validators []ValidateFunction) error {
	for _, fn := range validators {
		err := fn.Handler()

		if err != nil {
			return err
		}
	}

	return nil
}
