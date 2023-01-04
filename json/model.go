package swissjson

import (
	"errors"
	"fmt"
	"strings"

	"github.com/go-playground/validator/v10"
)

type JSONValidation struct {
	InputPath  string `validate:"required,file"`
	OutputPath string `validate:"omitempty"`
	JSON       string `validate:"required,json"`
}

func (jsonValidation *JSONValidation) validated() error {
	validate := validator.New()
	if err := validate.Struct(jsonValidation); err != nil {
		var errorMessages []string

		if _, ok := err.(*validator.InvalidValidationError); ok {
			return errors.New(err.Error())
		}

		for _, err := range err.(validator.ValidationErrors) {
			errorMessages = append(errorMessages, fmt.Sprintf("%s should be %s (%s)", err.Field(), err.Tag(), err.Param()))
		}

		return errors.New(strings.Join(errorMessages, "\n"))
	}

	return nil
}
