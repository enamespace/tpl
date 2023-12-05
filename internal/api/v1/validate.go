package v1

import (
	"github.com/go-playground/validator/v10"
)

func (u *User) Validate() []error {

	var errs []error
	validate := validator.New(validator.WithRequiredStructEnabled())

	err := validate.Struct(u)
	if err != nil {
		if _, ok := err.(*validator.InvalidValidationError); ok {
			return []error{err}
		}
		for _, err := range err.(validator.ValidationErrors) {
			errs = append(errs, err)
		}

		return errs
	}

	return nil
}
