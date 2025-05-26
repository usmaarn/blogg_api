package config

import (
	"fmt"
	"strings"

	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
)

func InitializeValidator() {
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterValidation("unique", uniqueEntityValidator)
	}
}

func uniqueEntityValidator(fl validator.FieldLevel) bool {
	field := fl.Field().String()
	params := strings.Split(fl.Param(), ".")

	db, err := DB.DB()
	if err != nil {
		return false
	}

	queryString := fmt.Sprintf("SELECT count(id) FROM %s WHERE %s = $1", params[0], params[1])
	row := db.QueryRow(queryString, field)

	var count int
	err = row.Scan(&count)
	if err != nil {
		return false
	}

	return count == 0
}
