package utils

import (
	"fmt"

	"github.com/go-playground/validator/v10"
)

func FormatValidationErrors(validationErrors validator.ValidationErrors) []string {
	var errs []string

	for _, err := range validationErrors {
		errs = append(errs, fmt.Sprintf("%s %s", err.Field(), getErrorMessage(err)))
	}

	return errs
}

func getErrorMessage(e validator.FieldError) string {
	switch e.Tag() {
	case "required":
		return "field is required"
	case "min":
		return fmt.Sprintf("field must be grater than %s characters", e.Param())
	case "alpha":
		return "field must contain alphabets only"
	default:
		return "field is not valid"
	}
}
