package validators

import (
	"github.com/go-playground/validator/v10"
)

var Validate *validator.Validate

func InitializeValidator() {
	validate := validator.New(validator.WithRequiredStructEnabled())
	validate.RegisterValidation("unique", UniqueEntityValidator)
	Validate = validate
}

func ValidateStruct(value any) (map[string]string, error) {
	errorMap := make(map[string]string)

	err := Validate.Struct(value)

	if err != nil {
		errorMap["message"] = err.Error()
	}

	return errorMap, err
}
