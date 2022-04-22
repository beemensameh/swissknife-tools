package swissuuid

import (
	"errors"
	"fmt"
	"strings"

	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
)

type UUIDCLI struct {
	Version  int         `validate:"required,gte=0,lte=5"`
	Domain   uuid.Domain `validate:"required_if=version 2,oneof=0 1 2"`
	Name     []byte      `validate:"omitempty"`
	Number   int64       `validate:"required,number"`
	Separate string      `validate:"omitempty"`
}

func (uuidCLI *UUIDCLI) validated() error {
	validate := validator.New()
	err := validate.Struct(uuidCLI)
	if err != nil {
		var errorMessages []string

		for _, err := range err.(validator.ValidationErrors) {
			errorMessages = append(errorMessages, fmt.Sprintf("%s should be %s (%s)", err.Field(), err.Tag(), err.Param()))
		}

		return errors.New(strings.Join(errorMessages, "\n"))
	}

	return nil
}
