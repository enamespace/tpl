package v1

import "github.com/go-playground/validator/v10"

type User struct {
	ID       uint64 `json:"id,omitempty"       gorm:"primary_key;AUTO_INCREMENT;column:id"`
	Name     string `json:"name,omitempty"     gorm:"column:name;not null"                 validate:"required"`
	Password string `json:"password,omitempty" gorm:"column:password"                      validate:"required"`
	Email    string `json:"email"              gorm:"column:email"                         validate:"required,email"`
	Gender   string `json:"gender"             gorm:"column:gender"                        validate:"oneof=male female prefer_not_to"`
	Status   int    `json:"status"             gorm:"column:status"                        validate:"omitempty"`
}

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
