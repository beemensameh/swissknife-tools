package swissuuid

import (
	"errors"
	"fmt"
	"strings"

	"github.com/go-playground/validator/v10"
)

type UUIDTypeEnum string

const (
	DCEGroup    UUIDTypeEnum = "group"
	DCEPerson   UUIDTypeEnum = "person"
	DCESecurity UUIDTypeEnum = "security"
)

type UUIDDCESecurityEnum string

const (
	Person UUIDDCESecurityEnum = "person"
	Group  UUIDDCESecurityEnum = "group"
	Org    UUIDDCESecurityEnum = "org"
)

type UUIDCLI struct {
	Version          int                 `validate:"required,gte=1,lte=5"`
	UUIDType         UUIDTypeEnum        `validate:"required_if=version 2,oneof=group person security"`
	UUIDSecurityType UUIDDCESecurityEnum `validate:"required_if=uuidType security,oneof=group person org"`
	Name             []byte              `validate:"omitempty"`
	Number           int64               `validate:"required,number"`
	Separate         string              `validate:"omitempty"`
}

func (uuidCLI *UUIDCLI) validated() error {
	validate := validator.New()
	err := validate.Struct(uuidCLI)
	if err != nil {
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
