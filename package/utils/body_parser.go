package utils

import (
	"encoding/json"
	"errors"
	"io"
)

func BindJSON[T any](body io.Reader, requestDto *T) error {
	if body == nil {
		return errors.New("request body is required")
	}

	dataByte, err := io.ReadAll(body)
	if err != nil {
		return err
	}

	isValidJson := json.Valid(dataByte)
	if !isValidJson {
		return err
	}

	err = json.Unmarshal(dataByte, requestDto)
	if err != nil {
		return err
	}

	return nil
}
