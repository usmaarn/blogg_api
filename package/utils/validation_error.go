package utils

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/go-playground/validator/v10"
	"github.com/usmaarn/blogg_api/package/utils/response"
)

func ValidateJson(e error, v any) any {
	if validationErrors, ok := e.(validator.ValidationErrors); ok {
		return response.ErrorResponse("", FormatValidationError(validationErrors, v))
	}
	return response.ErrorResponse(e.Error(), nil)
}

func FormatValidationError(validationErrors validator.ValidationErrors, s any) map[string]string {
	errorMap := make(map[string]string)

	for _, e := range validationErrors {
		field := GetJSONFieldName(s, e.Field())
		tag := e.Tag()
		param := e.Param()

		errorMap[field] = GetValidationErrorMessage(field, tag, param)
	}
	return errorMap
}

func GetValidationErrorMessage(field string, tag string, param string) string {
	switch tag {
	case "required":
		return fmt.Sprintf("%s is required", field)
	case "unique":
		return fmt.Sprintf("%s already exists", field)
	case "min":
		return fmt.Sprintf("%s must be %s and above", field, param)
	default:
		return fmt.Sprintf("%s is not valid", field)
	}
}

func GetJSONFieldName(structType interface{}, structField string) string {
	t := reflect.TypeOf(structType)
	if t.Kind() == reflect.Ptr {
		t = t.Elem()
	}

	field, ok := t.FieldByName(structField)
	if !ok {
		return structField // fallback
	}

	jsonTag := field.Tag.Get("json")
	if jsonTag == "" {
		return structField
	}

	return strings.Split(jsonTag, ",")[0]
}
