package helper

import "github.com/go-playground/validator/v10"

func Validate(data interface{}) error {
	err := validator.New().Struct(data)
	if err != nil {
		var fieldErrors []validator.FieldError

		for _, errs := range err.(validator.ValidationErrors) {
			fieldErrors = append(fieldErrors, errs)
		}
		return ErrorArrayToError(fieldErrors)
	}
	return nil
}
