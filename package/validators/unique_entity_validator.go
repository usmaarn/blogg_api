package validators

import (
	"fmt"
	"strings"

	"github.com/go-playground/validator/v10"
	"github.com/usmaarn/blogg_api/cmd/config"
)

func UniqueEntityValidator(fl validator.FieldLevel) bool {
	field := fl.Field().String()
	params := strings.Split(fl.Param(), ".")

	queryString := fmt.Sprintf("SELECT count(id) FROM %s WHERE %s = $1", params[0], params[1])
	row := config.DB.QueryRow(queryString, field)

	var count int
	err := row.Scan(&count)
	if err != nil {
		return false
	}

	return count == 0
}
